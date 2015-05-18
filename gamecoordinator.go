package steam

import (
	"bytes"
	"github.com/golang/protobuf/proto"
	. "github.com/vincentserpoul/go-steam/internal"
	. "github.com/vincentserpoul/go-steam/internal/gamecoordinator"
	. "github.com/vincentserpoul/go-steam/internal/protobuf"
	. "github.com/vincentserpoul/go-steam/internal/steamlang"
)

type GameCoordinator struct {
	client   *Client
	handlers []GCPacketHandler
}

func newGC(client *Client) *GameCoordinator {
	return &GameCoordinator{
		client:   client,
		handlers: make([]GCPacketHandler, 0),
	}
}

type GCPacketHandler interface {
	HandleGCPacket(*GCPacket)
}

func (g *GameCoordinator) RegisterPacketHandler(handler GCPacketHandler) {
	g.handlers = append(g.handlers, handler)
}

func (g *GameCoordinator) HandlePacket(packet *Packet) {
	if packet.EMsg != EMsg_ClientFromGC {
		return
	}

	msg := new(CMsgGCClient)
	packet.ReadProtoMsg(msg)

	p, err := NewGCPacket(msg)
	if err != nil {
		g.client.Errorf("Error reading GC message: %v", err)
		return
	}

	for _, handler := range g.handlers {
		handler.HandleGCPacket(p)
	}
}

func (g *GameCoordinator) Write(msg IGCMsg) {
	buf := new(bytes.Buffer)
	msg.Serialize(buf)

	msgType := msg.GetMsgType()
	if msg.IsProto() {
		msgType = msgType | 0x80000000 // mask with protoMask
	}

	g.client.Write(NewClientMsgProtobuf(EMsg_ClientToGC, &CMsgGCClient{
		Msgtype: proto.Uint32(msgType),
		Appid:   proto.Uint32(msg.GetAppId()),
		Payload: buf.Bytes(),
	}))
}

// Sets you in the given games. Specify none to quit all games.
func (g *GameCoordinator) SetGamesPlayed(appIds ...uint64) {
	games := make([]*CMsgClientGamesPlayed_GamePlayed, 0)
	for _, appId := range appIds {
		games = append(games, &CMsgClientGamesPlayed_GamePlayed{
			GameId: proto.Uint64(appId),
		})
	}

	g.client.Write(NewClientMsgProtobuf(EMsg_ClientGamesPlayed, &CMsgClientGamesPlayed{
		GamesPlayed: games,
	}))
}
