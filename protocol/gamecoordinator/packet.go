package gamecoordinator

import (
	"bytes"

	"github.com/Philipp15b/go-steam/v3/protocol"
	"github.com/Philipp15b/go-steam/v3/protocol/protobuf"
	"github.com/Philipp15b/go-steam/v3/protocol/steamlang"
	"google.golang.org/protobuf/proto"
)

// An incoming, partially unread message from the Game Coordinator.
type GCPacket struct {
	AppId       uint32
	MsgType     uint32
	IsProto     bool
	GCName      string
	Body        []byte
	TargetJobId protocol.JobId
}

func NewGCPacket(wrapper *protobuf.CMsgGCClient) (*GCPacket, error) {
	packet := &GCPacket{
		AppId:   wrapper.GetAppid(),
		MsgType: wrapper.GetMsgtype(),
		GCName:  wrapper.GetGcname(),
	}

	r := bytes.NewReader(wrapper.GetPayload())
	if steamlang.IsProto(wrapper.GetMsgtype()) {
		packet.MsgType = packet.MsgType & steamlang.EMsgMask
		packet.IsProto = true

		header := steamlang.NewMsgGCHdrProtoBuf()
		err := header.Deserialize(r)
		if err != nil {
			return nil, err
		}
		packet.TargetJobId = protocol.JobId(header.Proto.GetJobidTarget())
	} else {
		header := steamlang.NewMsgGCHdr()
		err := header.Deserialize(r)
		if err != nil {
			return nil, err
		}
		packet.TargetJobId = protocol.JobId(header.TargetJobID)
	}

	body := make([]byte, r.Len())
	r.Read(body)
	packet.Body = body

	return packet, nil
}

func (g *GCPacket) ReadProtoMsg(body proto.Message) {
	proto.Unmarshal(g.Body, body)
}

func (g *GCPacket) ReadMsg(body protocol.MessageBody) {
	body.Deserialize(bytes.NewReader(g.Body))
}
