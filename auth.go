package steam

import (
	"crypto/sha1"
	. "github.com/smithfox/go-steam/internal"
	. "github.com/smithfox/go-steam/internal/protobuf"
	. "github.com/smithfox/go-steam/internal/steamlang"
	"github.com/smithfox/go-steam/steamid"
	"github.com/smithfox/goprotobuf/proto"
	// "sync/atomic"
	"time"
)

type Auth struct {
	client  *Client
	details *LogOnDetails
}

func (a *Auth) HandlePacket(packet *Packet) {
	switch packet.EMsg {
	case EMsg_ClientLogOnResponse:
		a.handleClientLogOnResponse(packet)
	case EMsg_ClientUpdateMachineAuth:
		a.handleClientUpdateMachineAuth(packet)
	}
}

type SentryHash []byte

type LogOnDetails struct {
	Username       string
	Password       string
	AuthCode       string
	SentryFileHash SentryHash
}

// Log on with the given details. You must always specify username and
// password. For the first login, don't set an authcode or a hash and you'll receive an error
// and Steam will send you an authcode. Then you have to login again, this time with the authcode.
// Shortly after logging in, you'll receive a MachineAuthUpdateEvent with a hash which allows
// you to login without using an authcode in the future.
//
// If you don't use Steam Guard, username and password are enough.
func (a *Auth) LogOn(details *LogOnDetails) {
	if len(details.Username) == 0 || len(details.Password) == 0 {
		panic("Username and password must be set!")
	}

	uuu := uint64(steamid.New(0, 1, int32(EUniverse_Public), EAccountType_Individual))

	a.client.steamId = uuu
	//atomic.StoreUint64(&(a.client.steamId), uuu)
	logon := new(CMsgClientLogon)
	logon.AccountName = &details.Username
	logon.Password = &details.Password
	if details.AuthCode != "" {
		logon.AuthCode = proto.String(details.AuthCode)
	}
	logon.ShaSentryfile = details.SentryFileHash

	logon.ProtocolVersion = proto.Uint32(MsgClientLogon_CurrentProtocol)

	a.client.Write(NewClientMsgProtobuf(EMsg_ClientLogon, logon))
}

type LoggedOnEvent struct{}

func (a *Auth) handleClientLogOnResponse(packet *Packet) {
	if !packet.IsProto {
		a.client.Fatalf("Got non-proto logon response!")
		return
	}

	body := new(CMsgClientLogonResponse)
	msg := packet.ReadProtoMsg(body)

	result := EResult(body.GetEresult())
	if result == EResult_OK {
		a.client.sessionId = msg.Header.Proto.GetClientSessionid()
		//atomic.StoreInt32(&a.client.sessionId, msg.Header.Proto.GetClientSessionid())
		a.client.steamId = msg.Header.Proto.GetSteamid()
		//atomic.StoreUint64(&a.client.steamId, msg.Header.Proto.GetSteamid())

		go a.client.heartbeatLoop(time.Duration(body.GetOutOfGameHeartbeatSeconds()))

		a.client.Emit(&LoggedOnEvent{})
	} else if result == EResult_Fail || result == EResult_ServiceUnavailable || result == EResult_TryAnotherCM {
		// some error on Steam's side, we'll get an EOF later
	} else {
		a.client.Fatalf("Login error: %v", result)
	}
}

type MachineAuthUpdateEvent struct {
	Hash SentryHash
}

func (a *Auth) handleClientUpdateMachineAuth(packet *Packet) {
	body := new(CMsgClientUpdateMachineAuth)
	packet.ReadProtoMsg(body)

	hash := sha1.New()
	hash.Write(body.GetBytes())
	sha := hash.Sum(nil)

	msg := NewClientMsgProtobuf(EMsg_ClientUpdateMachineAuthResponse, &CMsgClientUpdateMachineAuthResponse{
		ShaFile: sha,
	})
	msg.SetTargetJobId(packet.SourceJobId)
	a.client.Write(msg)

	a.client.Emit(&MachineAuthUpdateEvent{sha})
}
