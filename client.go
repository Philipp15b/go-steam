package steam

import (
	"bytes"
	"compress/gzip"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"github.com/Philipp15b/go-steam/v3/appticket"
	"google.golang.org/protobuf/proto"
	"hash/crc32"
	"io/ioutil"
	"log"
	"net"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Philipp15b/go-steam/v3/cryptoutil"
	"github.com/Philipp15b/go-steam/v3/netutil"
	"github.com/Philipp15b/go-steam/v3/protocol"
	"github.com/Philipp15b/go-steam/v3/protocol/protobuf"
	"github.com/Philipp15b/go-steam/v3/protocol/steamlang"
	"github.com/Philipp15b/go-steam/v3/steamid"
)

// Represents a client to the Steam network.
// Always poll events from the channel returned by Events() or receiving messages will stop.
// All access, unless otherwise noted, should be threadsafe.
//
// When a FatalErrorEvent is emitted, the connection is automatically closed. The same client can be used to reconnect.
// Other errors don't have any effect.
type Client struct {
	// these need to be 64 bit aligned for sync/atomic on 32bit
	sessionId    int32
	_            uint32
	steamId      uint64
	currentJobId uint64

	Auth          *Auth
	Social        *Social
	Web           *Web
	Notifications *Notifications
	Trading       *Trading
	GC            *GameCoordinator

	events        chan interface{}
	handlers      []PacketHandler
	handlersMutex sync.RWMutex

	tempSessionKey []byte

	ConnectionTimeout time.Duration

	connectTime     time.Time
	connectionCount int
	gcTokens        [][]byte
	gcTokensMutex   sync.Mutex
	publicIp        uint32

	mutex     sync.RWMutex // guarding conn and writeChan
	conn      connection
	writeChan chan protocol.IMsg
	writeBuf  *bytes.Buffer
	heartbeat *time.Ticker
}

type PacketHandler interface {
	HandlePacket(*protocol.Packet)
}

func NewClient() *Client {
	client := &Client{
		events:   make(chan interface{}, 3),
		writeBuf: new(bytes.Buffer),
		gcTokens: make([][]byte, 0),
	}

	client.Auth = &Auth{client: client}
	client.RegisterPacketHandler(client.Auth)

	client.Social = newSocial(client)
	client.RegisterPacketHandler(client.Social)

	client.Web = &Web{client: client}
	client.RegisterPacketHandler(client.Web)

	client.Notifications = newNotifications(client)
	client.RegisterPacketHandler(client.Notifications)

	client.Trading = &Trading{client: client}
	client.RegisterPacketHandler(client.Trading)

	client.GC = newGC(client)
	client.RegisterPacketHandler(client.GC)

	return client
}

// Get the event channel. By convention all events are pointers, except for errors.
// It is never closed.
func (c *Client) Events() <-chan interface{} {
	return c.events
}

func (c *Client) Emit(event interface{}) {
	c.events <- event
}

// Emits a FatalErrorEvent formatted with fmt.Errorf and disconnects.
func (c *Client) Fatalf(format string, a ...interface{}) {
	c.Emit(FatalErrorEvent(fmt.Errorf(format, a...)))
	c.Disconnect()
}

// Emits an error formatted with fmt.Errorf.
func (c *Client) Errorf(format string, a ...interface{}) {
	c.Emit(fmt.Errorf(format, a...))
}

// Registers a PacketHandler that receives all incoming packets.
func (c *Client) RegisterPacketHandler(handler PacketHandler) {
	c.handlersMutex.Lock()
	defer c.handlersMutex.Unlock()
	c.handlers = append(c.handlers, handler)
}

func (c *Client) GetNextJobId() protocol.JobId {
	return protocol.JobId(atomic.AddUint64(&c.currentJobId, 1))
}

func (c *Client) SteamId() steamid.SteamId {
	return steamid.SteamId(atomic.LoadUint64(&c.steamId))
}

func (c *Client) SessionId() int32 {
	return atomic.LoadInt32(&c.sessionId)
}

func (c *Client) Connected() bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.conn != nil
}

// Connects to a random Steam server and returns its address.
// If this client is already connected, it is disconnected first.
// This method tries to use an address from the Steam Directory and falls
// back to the built-in server list if the Steam Directory can't be reached.
// If you want to connect to a specific server, use `ConnectTo`.
func (c *Client) Connect() (*netutil.PortAddr, error) {
	var server *netutil.PortAddr

	// try to initialize the directory cache
	if !steamDirectoryCache.IsInitialized() {
		_ = steamDirectoryCache.Initialize()
	}
	if steamDirectoryCache.IsInitialized() {
		server = steamDirectoryCache.GetRandomCM()
	} else {
		server = GetRandomCM()
	}

	err := c.ConnectTo(server)
	return server, err
}

// Connects to a specific server.
// You may want to use one of the `GetRandom*CM()` functions in this package.
// If this client is already connected, it is disconnected first.
func (c *Client) ConnectTo(addr *netutil.PortAddr) error {
	return c.ConnectToBind(addr, nil)
}

// Connects to a specific server, and binds to a specified local IP
// If this client is already connected, it is disconnected first.
func (c *Client) ConnectToBind(addr *netutil.PortAddr, local *net.TCPAddr) error {
	c.Disconnect()

	conn, err := dialTCP(local, addr.ToTCPAddr())
	if err != nil {
		c.Fatalf("Connect failed: %v", err)
		return err
	}
	c.conn = conn
	c.writeChan = make(chan protocol.IMsg, 5)

	go c.readLoop()
	go c.writeLoop()

	return nil
}

func (c *Client) Disconnect() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.conn == nil {
		return
	}

	c.conn.Close()
	c.conn = nil
	if c.heartbeat != nil {
		c.heartbeat.Stop()
	}
	close(c.writeChan)
	c.Emit(&DisconnectedEvent{})

}

// Adds a message to the send queue. Modifications to the given message after
// writing are not allowed (possible race conditions).
//
// Writes to this client when not connected are ignored.
func (c *Client) Write(msg protocol.IMsg) {
	if cm, ok := msg.(protocol.IClientMsg); ok {
		cm.SetSessionId(c.SessionId())
		cm.SetSteamId(c.SteamId())
	}
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if c.conn == nil {
		return
	}
	c.writeChan <- msg
}

func (c *Client) readLoop() {
	for {
		// This *should* be atomic on most platforms, but the Go spec doesn't guarantee it
		c.mutex.RLock()
		conn := c.conn
		c.mutex.RUnlock()
		if conn == nil {
			return
		}
		packet, err := conn.Read()

		if err != nil {
			c.Fatalf("Error reading from the connection: %v", err)
			return
		}
		c.handlePacket(packet)
	}
}

func (c *Client) writeLoop() {
	for {
		c.mutex.RLock()
		conn := c.conn
		c.mutex.RUnlock()
		if conn == nil {
			return
		}

		msg, ok := <-c.writeChan
		if !ok {
			return
		}

		err := msg.Serialize(c.writeBuf)
		if err != nil {
			c.writeBuf.Reset()
			c.Fatalf("Error serializing message %v: %v", msg, err)
			return
		}

		err = conn.Write(c.writeBuf.Bytes())

		c.writeBuf.Reset()

		if err != nil {
			c.Fatalf("Error writing message %v: %v", msg, err)
			return
		}
	}
}

func (c *Client) heartbeatLoop(seconds time.Duration) {
	if c.heartbeat != nil {
		c.heartbeat.Stop()
	}
	c.heartbeat = time.NewTicker(seconds * time.Second)
	for {
		_, ok := <-c.heartbeat.C
		if !ok {
			break
		}
		c.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientHeartBeat, new(protobuf.CMsgClientHeartBeat)))
	}
	c.heartbeat = nil
}

func (c *Client) handlePacket(packet *protocol.Packet) {
	switch packet.EMsg {
	case steamlang.EMsg_ChannelEncryptRequest:
		c.handleChannelEncryptRequest(packet)
	case steamlang.EMsg_ChannelEncryptResult:
		c.handleChannelEncryptResult(packet)
	case steamlang.EMsg_Multi:
		c.handleMulti(packet)
	case steamlang.EMsg_ClientCMList:
		c.handleClientCMList(packet)
	case steamlang.EMsg_ClientGameConnectTokens:
		c.handleGameClientToken(packet)
	case steamlang.EMsg_ClientGetAppOwnershipTicketResponse:
		c.handleGetAppOwnershipTicketResponse(packet)
	case steamlang.EMsg_ClientAuthListAck:
		c.handleClientAuthListAck(packet)
	case steamlang.EMsg_ClientTicketAuthComplete:
		c.handleClientTicketAuthComplete(packet)
	}

	c.handlersMutex.RLock()
	defer c.handlersMutex.RUnlock()
	for _, handler := range c.handlers {
		handler.HandlePacket(packet)
	}
}

func (c *Client) handleChannelEncryptRequest(packet *protocol.Packet) {
	body := steamlang.NewMsgChannelEncryptRequest()
	packet.ReadMsg(body)

	if body.Universe != steamlang.EUniverse_Public {
		c.Fatalf("Invalid univserse %v!", body.Universe)
	}

	c.tempSessionKey = make([]byte, 32)
	rand.Read(c.tempSessionKey)
	encryptedKey := cryptoutil.RSAEncrypt(GetPublicKey(steamlang.EUniverse_Public), c.tempSessionKey)

	payload := new(bytes.Buffer)
	payload.Write(encryptedKey)
	binary.Write(payload, binary.LittleEndian, crc32.ChecksumIEEE(encryptedKey))
	payload.WriteByte(0)
	payload.WriteByte(0)
	payload.WriteByte(0)
	payload.WriteByte(0)

	c.Write(protocol.NewMsg(steamlang.NewMsgChannelEncryptResponse(), payload.Bytes()))
}

func (c *Client) handleChannelEncryptResult(packet *protocol.Packet) {
	body := steamlang.NewMsgChannelEncryptResult()
	packet.ReadMsg(body)

	if body.Result != steamlang.EResult_OK {
		c.Fatalf("Encryption failed: %v", body.Result)
		return
	}
	c.conn.SetEncryptionKey(c.tempSessionKey)
	c.tempSessionKey = nil

	c.Emit(&ConnectedEvent{})
}

func (c *Client) handleMulti(packet *protocol.Packet) {
	body := new(protobuf.CMsgMulti)
	packet.ReadProtoMsg(body)

	payload := body.GetMessageBody()

	if body.GetSizeUnzipped() > 0 {
		r, err := gzip.NewReader(bytes.NewReader(payload))
		if err != nil {
			c.Errorf("handleMulti: Error while decompressing: %v", err)
			return
		}

		payload, err = ioutil.ReadAll(r)
		if err != nil {
			c.Errorf("handleMulti: Error while decompressing: %v", err)
			return
		}
	}

	pr := bytes.NewReader(payload)
	for pr.Len() > 0 {
		var length uint32
		binary.Read(pr, binary.LittleEndian, &length)
		packetData := make([]byte, length)
		pr.Read(packetData)
		p, err := protocol.NewPacket(packetData)
		if err != nil {
			c.Errorf("Error reading packet in Multi msg %v: %v", packet, err)
			continue
		}
		c.handlePacket(p)
	}
}

func (c *Client) handleClientCMList(packet *protocol.Packet) {
	body := new(protobuf.CMsgClientCMList)
	packet.ReadProtoMsg(body)

	l := make([]*netutil.PortAddr, 0)
	for i, ip := range body.GetCmAddresses() {
		l = append(l, &netutil.PortAddr{
			IP:   readIp(ip),
			Port: uint16(body.GetCmPorts()[i]),
		})
	}

	c.Emit(&ClientCMListEvent{l})
}

func (c *Client) handleGameClientToken(packet *protocol.Packet) {
	body := new(protobuf.CMsgClientGameConnectTokens)
	packet.ReadProtoMsg(body)

	c.gcTokensMutex.Lock()
	defer c.gcTokensMutex.Unlock()

	c.gcTokens = append(c.gcTokens, body.Tokens...)
}

func (c *Client) pullGcToken() []byte {
	c.gcTokensMutex.Lock()
	defer c.gcTokensMutex.Unlock()

	if len(c.gcTokens) == 0 {
		return nil
	}
	token := c.gcTokens[0]
	c.gcTokens = c.gcTokens[1:len(c.gcTokens)]

	return token
}

func (c *Client) handleGetAppOwnershipTicketResponse(packet *protocol.Packet) {
	body := new(protobuf.CMsgClientGetAppOwnershipTicketResponse)
	packet.ReadProtoMsg(body)
	if body.GetEresult() != uint32(steamlang.EResult_OK) {
		log.Fatalf("CMsgClientGetAppOwnershipTicketResponse error: %+#v", body)
		// TODO: Add event
		return
	}

	ioutil.WriteFile(c.generateAppTicketFileCacheName(body.GetAppId()), body.GetTicket(), 0744)
	ticket, err := appticket.NewAppTicket(body.GetTicket())
	if err != nil {
		log.Fatalf("CMsgClientGetAppOwnershipTicketResponse invalid ticket: %v\n %+#v", err, body)
		// TODO: Add event?
		return
	}

	c.onVaildAppOwnershipTicket(ticket)
}

func (c *Client) generateAppTicketFileCacheName(appId uint32) string {
	return fmt.Sprintf("appOwnershipTicket_%d_%d.bin", c.steamId, appId)
}

func (c *Client) GetAppOwnershipTicket(appId uint32) {
	cachedTickedFileName := c.generateAppTicketFileCacheName(appId)

	appTicket, err := ioutil.ReadFile(cachedTickedFileName)
	if err == nil {
		parsedTicket, err := appticket.NewAppTicket(appTicket)

		if err == nil && parsedTicket.IsExpired(time.Now().Add(time.Minute)) == false {
			go c.onVaildAppOwnershipTicket(parsedTicket)
			return
		} else {
			os.Remove(cachedTickedFileName)
		}
	}

	c.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientGetAppOwnershipTicket, &protobuf.CMsgClientGetAppOwnershipTicket{
		AppId: proto.Uint32(appId),
	}))
}

func (c *Client) onVaildAppOwnershipTicket(appticket *appticket.AppTicket) {
	c.Emit(&AppOwnershipTicket{AppOwnershipTicket: appticket})
}

func (c *Client) createAuthToken(gameConnectToken []byte) []byte {
	var buf1 [4]byte
	var buf2 [28]byte

	timestamp := uint32(time.Now().Sub(c.connectTime).Milliseconds())

	binary.LittleEndian.PutUint32(buf1[0:], uint32(len(gameConnectToken)))

	binary.LittleEndian.PutUint32(buf2[0:], 6*4)
	binary.LittleEndian.PutUint32(buf2[4:], 1)
	binary.LittleEndian.PutUint32(buf2[8:], 2)
	binary.BigEndian.PutUint32(buf2[12:], c.publicIp)
	binary.LittleEndian.PutUint32(buf2[16:], 0)
	binary.LittleEndian.PutUint32(buf2[20:], timestamp)
	binary.LittleEndian.PutUint32(buf2[24:], 1)

	result := make([]byte, 0)
	result = append(result, buf1[:]...)
	result = append(result, gameConnectToken...)
	result = append(result, buf2[:]...)

	return result
}

func (c *Client) handleClientAuthListAck(packet *protocol.Packet) {
	c.Emit(&TicketAuthAck{})
}

func (c *Client) handleClientTicketAuthComplete(packet *protocol.Packet) {
	c.Emit(&TicketAuthComplete{})
}

func (c *Client) AuthSessionTicket(ticket *appticket.AppTicket) ([]byte, error) {
	var buf1 [4]byte
	var buf2 [32]byte

	gcToken := c.pullGcToken()
	if gcToken == nil {
		return gcToken, fmt.Errorf("empty gc token")
	}

	bufTicket := ticket.OriginalBuffer()

	c.connectionCount++

	timestamp := uint32(time.Now().Sub(c.connectTime).Milliseconds())

	binary.LittleEndian.PutUint32(buf1[:], uint32(len(gcToken)))

	binary.LittleEndian.PutUint32(buf2[0:], 24)
	binary.LittleEndian.PutUint32(buf2[4:], 1)
	binary.LittleEndian.PutUint32(buf2[8:], 2)
	binary.BigEndian.PutUint32(buf2[12:], c.publicIp)
	binary.LittleEndian.PutUint32(buf2[16:], 0)
	binary.LittleEndian.PutUint32(buf2[20:], timestamp)
	binary.LittleEndian.PutUint32(buf2[24:], uint32(c.connectionCount))
	binary.LittleEndian.PutUint32(buf2[28:], uint32(len(bufTicket)))

	result := make([]byte, 0)
	result = append(result, buf1[:]...)
	result = append(result, gcToken...)
	result = append(result, buf2[:]...)
	result = append(result, bufTicket...)

	gameId := uint64(ticket.AppId)

	tokensLeft := uint32(len(c.gcTokens))

	authToken := c.createAuthToken(gcToken)
	ticketCrc := crc32.ChecksumIEEE(authToken)

	authList := new(protobuf.CMsgClientAuthList)
	authList.TokensLeft = &tokensLeft
	authList.AppIds = []uint32{ticket.AppId}
	authList.Tickets = []*protobuf.CMsgAuthTicket{
		&protobuf.CMsgAuthTicket{
			Gameid:    &gameId,
			Ticket:    gcToken,
			TicketCrc: &ticketCrc,
		},
	}

	c.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientAuthList, authList))

	return result, nil
}

func readIp(ip uint32) net.IP {
	r := make(net.IP, 4)
	r[3] = byte(ip)
	r[2] = byte(ip >> 8)
	r[1] = byte(ip >> 16)
	r[0] = byte(ip >> 24)
	return r
}
