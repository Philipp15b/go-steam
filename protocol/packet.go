package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/Philipp15b/go-steam/v3/protocol/steamlang"
	"google.golang.org/protobuf/proto"
)

// TODO: Headers are always deserialized twice.

// Represents an incoming, partially unread message.
type Packet struct {
	EMsg        steamlang.EMsg
	IsProto     bool
	TargetJobId JobId
	SourceJobId JobId
	Data        []byte
}

func NewPacket(data []byte) (*Packet, error) {
	var rawEMsg uint32
	err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &rawEMsg)
	if err != nil {
		return nil, err
	}
	eMsg := steamlang.NewEMsg(rawEMsg)
	buf := bytes.NewReader(data)
	if eMsg == steamlang.EMsg_ChannelEncryptRequest || eMsg == steamlang.EMsg_ChannelEncryptResult {
		header := steamlang.NewMsgHdr()
		header.Msg = eMsg
		err = header.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return &Packet{
			EMsg:        eMsg,
			IsProto:     false,
			TargetJobId: JobId(header.TargetJobID),
			SourceJobId: JobId(header.SourceJobID),
			Data:        data,
		}, nil
	} else if steamlang.IsProto(rawEMsg) {
		header := steamlang.NewMsgHdrProtoBuf()
		header.Msg = eMsg
		err = header.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return &Packet{
			EMsg:        eMsg,
			IsProto:     true,
			TargetJobId: JobId(header.Proto.GetJobidTarget()),
			SourceJobId: JobId(header.Proto.GetJobidSource()),
			Data:        data,
		}, nil
	} else {
		header := steamlang.NewExtendedClientMsgHdr()
		header.Msg = eMsg
		err = header.Deserialize(buf)
		if err != nil {
			return nil, err
		}
		return &Packet{
			EMsg:        eMsg,
			IsProto:     false,
			TargetJobId: JobId(header.TargetJobID),
			SourceJobId: JobId(header.SourceJobID),
			Data:        data,
		}, nil
	}
}

func (p *Packet) String() string {
	return fmt.Sprintf("Packet{EMsg = %v, Proto = %v, Len = %v, TargetJobId = %v, SourceJobId = %v}", p.EMsg, p.IsProto, len(p.Data), p.TargetJobId, p.SourceJobId)
}

func (p *Packet) ReadProtoMsg(body proto.Message) *ClientMsgProtobuf {
	header := steamlang.NewMsgHdrProtoBuf()
	buf := bytes.NewBuffer(p.Data)
	header.Deserialize(buf)
	proto.Unmarshal(buf.Bytes(), body)
	return &ClientMsgProtobuf{ // protobuf messages have no payload
		Header: header,
		Body:   body,
	}
}

func (p *Packet) ReadClientMsg(body MessageBody) *ClientMsg {
	header := steamlang.NewExtendedClientMsgHdr()
	buf := bytes.NewReader(p.Data)
	header.Deserialize(buf)
	body.Deserialize(buf)
	payload := make([]byte, buf.Len())
	buf.Read(payload)
	return &ClientMsg{
		Header:  header,
		Body:    body,
		Payload: payload,
	}
}

func (p *Packet) ReadMsg(body MessageBody) *Msg {
	header := steamlang.NewMsgHdr()
	buf := bytes.NewReader(p.Data)
	header.Deserialize(buf)
	body.Deserialize(buf)
	payload := make([]byte, buf.Len())
	buf.Read(payload)
	return &Msg{
		Header:  header,
		Body:    body,
		Payload: payload,
	}
}
