package steam

import (
	"archive/zip"
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"github.com/Philipp15b/go-steam/cryptoutil"
	. "github.com/Philipp15b/go-steam/internal"
	. "github.com/Philipp15b/go-steam/steamid"
	"hash/crc32"
	"io/ioutil"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

// Represents a client to the Steam network.
// Always poll events from the channel returned by Events() or receiving messages will stop.
// All access, unless otherwise noted, should be threadsafe.
//
// When a FatalError is emitted, the connection is automatically closed. The same client can be used to reconnect.
// Other errors don't have any effect.
type Client struct {
	Auth    *Auth
	Social  *Social
	Web     *Web
	Trading *Trading
	GC      *GameCoordinator

	sessionId int32
	steamId   uint64

	currentJobId uint64

	events   chan interface{}
	handlers []PacketHandler

	tempSessionKey []byte

	ConnectionTimeout time.Duration

	mutex     sync.RWMutex // guarding connection, heartbeat and writeChan
	conn      connection
	writeChan chan IMsg
	writeBuf  *bytes.Buffer
	heartbeat *time.Ticker
}

type PacketHandler interface {
	HandlePacket(*PacketMsg)
}

func NewClient() *Client {
	client := &Client{
		events:    make(chan interface{}, 3),
		writeChan: make(chan IMsg, 5),
		writeBuf:  new(bytes.Buffer),
	}
	client.Auth = &Auth{client: client}
	client.RegisterPacketHandler(client.Auth)
	client.Social = newSocial(client)
	client.RegisterPacketHandler(client.Social)
	client.Web = &Web{client: client}
	client.RegisterPacketHandler(client.Web)
	client.Trading = &Trading{client: client}
	client.RegisterPacketHandler(client.Trading)
	client.GC = newGC(client)
	client.RegisterPacketHandler(client.GC)
	return client
}

// Get the event channel. By convention all events are pointers, except for errors.
func (c *Client) Events() <-chan interface{} {
	return c.events
}

func (c *Client) Emit(event interface{}) {
	c.events <- event
}

// When this error is emitted by the Client, the connection is automatically closed.
// This may be a network error, for example.
type FatalError error

// Emits a FatalError formatted with fmt.Errorf and disconnects.
func (c *Client) Fatalf(format string, a ...interface{}) {
	c.Emit(FatalError(fmt.Errorf(format, a...)))
	c.Disconnect()
}

// Emits an error formatted with fmt.Errorf.
func (c *Client) Errorf(format string, a ...interface{}) {
	c.Emit(fmt.Errorf(format, a...))
}

// Registers a PacketHandler that receives all incoming packets.
func (c *Client) RegisterPacketHandler(handler PacketHandler) {
	c.handlers = append(c.handlers, handler)
}

func (c *Client) GetNextJobId() JobId {
	return JobId(atomic.AddUint64(&c.currentJobId, 1))
}

func (c *Client) SteamId() SteamId {
	return SteamId(atomic.LoadUint64(&c.steamId))
}

func (c *Client) SessionId() int32 {
	return atomic.LoadInt32(&c.sessionId)
}

func (c *Client) Connected() bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.conn != nil
}

// Connects to a random server of the Steam network and returns the server.
// If this client is already connected, it is disconnected first.
func (c *Client) Connect() string {
	server := getRandomCM()
	c.ConnectTo(server)
	return server
}

// Connects to a specific server.
// If this client is already connected, it is disconnected first.
func (c *Client) ConnectTo(address string) {
	c.Disconnect()

	conn, err := dialTCP(address)
	if err != nil {
		log.Fatal(err)
	}
	c.conn = conn

	go c.readLoop()
	go c.writeLoop()
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
		c.heartbeat = nil
	}
	close(c.writeChan)
}

// Adds a message to the send queue. Modifications to the given message after
// writing are not allowed (possible race conditions).
//
// Writes to this client when not connected are ignored.
func (c *Client) Write(msg IMsg) {
	if cm, ok := msg.(IClientMsg); ok {
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
		c.mutex.RLock()
		packet, err := c.conn.Read()
		c.mutex.RUnlock()
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
		if c.conn == nil {
			c.mutex.RUnlock()
			return
		}
		msg, ok := <-c.writeChan
		c.mutex.RUnlock()
		if !ok {
			return
		}

		err := msg.Serialize(c.writeBuf)
		if err != nil {
			c.writeBuf.Reset()
			c.Errorf("Error serializing message %v: %v", msg, err)
			return
		}

		c.mutex.RLock()
		err = c.conn.Write(c.writeBuf.Bytes())
		c.mutex.RLock()

		c.writeBuf.Reset()

		if err != nil {
			c.writeBuf.Reset()
			c.Errorf("Error writing message %v: %v", msg, err)
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
		c.Write(NewClientMsgProtobuf(EMsg_ClientHeartBeat, new(CMsgClientHeartBeat)))
	}
	c.heartbeat = nil
}

func (c *Client) handlePacket(packet *PacketMsg) {
	switch packet.EMsg {
	case EMsg_ChannelEncryptRequest:
		c.handleChannelEncryptRequest(packet)
	case EMsg_ChannelEncryptResult:
		c.handleChannelEncryptResult(packet)
	case EMsg_Multi:
		c.handleMulti(packet)
	}
	for _, handler := range c.handlers {
		handler.HandlePacket(packet)
	}
}

func (c *Client) handleChannelEncryptRequest(packet *PacketMsg) {
	body := NewMsgChannelEncryptRequest()
	packet.ReadMsg(body)

	if body.Universe != EUniverse_Public {
		c.Fatalf("Invalid univserse %v!", body.Universe)
	}

	c.tempSessionKey = make([]byte, 32)
	rand.Read(c.tempSessionKey)
	encryptedKey := cryptoutil.RSAEncrypt(GetPublicKey(EUniverse_Public), c.tempSessionKey)

	payload := new(bytes.Buffer)
	payload.Write(encryptedKey)
	binary.Write(payload, binary.LittleEndian, crc32.ChecksumIEEE(encryptedKey))
	payload.WriteByte(0)
	payload.WriteByte(0)
	payload.WriteByte(0)
	payload.WriteByte(0)

	c.Write(NewMsg(NewMsgChannelEncryptResponse(), payload.Bytes()))
}

type ConnectedEvent struct{}

func (c *Client) handleChannelEncryptResult(packet *PacketMsg) {
	body := NewMsgChannelEncryptResult()
	packet.ReadMsg(body)

	if body.Result != EResult_OK {
		c.Fatalf("Encryption failed: %v", body.Result)
		return
	}
	c.conn.SetEncryptionKey(c.tempSessionKey)
	c.tempSessionKey = nil

	c.Emit(new(ConnectedEvent))
}

func (c *Client) handleMulti(packet *PacketMsg) {
	body := new(CMsgMulti)
	packet.ReadProtoMsg(body)

	payload := body.GetMessageBody()

	if body.GetSizeUnzipped() > 0 {
		archive, err := zip.NewReader(bytes.NewReader(payload), int64(len(payload)))
		if err != nil {
			panic(err)
		}

		for _, f := range archive.File {
			if f.Name == "z" {
				r, _ := f.Open()
				payload, _ = ioutil.ReadAll(r)
				goto okay
			}
		}

		c.Errorf("Invalid Multi packet %v: Could not find 'z' file!", packet)
		return

	okay: // jump over error
	}

	pr := bytes.NewReader(payload)
	for pr.Len() > 0 {
		var length uint32
		binary.Read(pr, binary.LittleEndian, &length)
		packetData := make([]byte, length)
		pr.Read(packetData)
		p, err := NewPacketMsg(packetData)
		if err != nil {
			c.Errorf("Error reading packet in Multi msg %v: %v", packet, err)
			continue
		}
		c.handlePacket(p)
	}

}
