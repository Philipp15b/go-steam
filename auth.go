package steam

import (
	"code.google.com/p/goprotobuf/proto"
	"crypto/sha1"
	. "github.com/Philipp15b/go-steam/internal"
	"github.com/Philipp15b/go-steam/steamid"
	"log"
	"sync/atomic"
	"time"
)

type Auth struct {
	client  *Client
	details *LogOnDetails
}

func (a *Auth) HandlePacket(packet *PacketMsg) {
	switch packet.EMsg {
	case EMsg_ClientLogOnResponse:
		a.handleClientLogOnResponse(packet)
	case EMsg_ClientUpdateMachineAuth:
		a.handleClientUpdateMachineAuth(packet)
	}
}

type LogOnDetails struct {
	Username       string
	Password       string
	AuthCode       string
	SentryFileHash []byte
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

	logon := new(CMsgClientLogon)
	logon.AccountName = &details.Username
	logon.Password = &details.Password
	if details.AuthCode != "" {
		logon.AuthCode = proto.String(details.AuthCode)
	}
	logon.ShaSentryfile = details.SentryFileHash

	logon.ProtocolVersion = proto.Uint32(MsgClientLogon_CurrentProtocol)

	atomic.StoreUint64(&a.client.steamId, uint64(steamid.New(0, 1, int32(EUniverse_Public), EAccountType_Individual)))

	a.client.Write(NewClientMsgProtobuf(EMsg_ClientLogon, logon))
}

type LoggedOnEvent struct{}

func (a *Auth) handleClientLogOnResponse(packet *PacketMsg) {
	if !packet.IsProto {
		a.client.Fatalf("Got non-proto logon response!")
		return
	}

	body := new(CMsgClientLogonResponse)
	msg := packet.ReadProtoMsg(body)

	result := EResult(body.GetEresult())
	log.Println(result)
	if result == EResult_OK {
		atomic.StoreInt32(&a.client.sessionId, msg.Header.Proto.GetClientSessionid())
		atomic.StoreUint64(&a.client.steamId, msg.Header.Proto.GetSteamid())

		go a.client.heartbeatLoop(time.Duration(body.GetOutOfGameHeartbeatSeconds()))

		a.client.Emit(&LoggedOnEvent{})
	} else if result == EResult_Fail || result == EResult_ServiceUnavailable || result == EResult_TryAnotherCM {
		// some error on Steam's side, we'll get an EOF later
	} else {
		a.client.Fatalf("Login error: %v", result)
	}
}

type MachineAuthUpdateEvent struct {
	Hash []byte
}

func (a *Auth) handleClientUpdateMachineAuth(packet *PacketMsg) {
	body := new(CMsgClientUpdateMachineAuth)
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
