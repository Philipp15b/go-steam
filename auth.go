package steam

import (
	"crypto/sha1"
	"sync/atomic"
	"time"

	"github.com/Philipp15b/go-steam/v3/protocol"
	"github.com/Philipp15b/go-steam/v3/protocol/protobuf"
	"github.com/Philipp15b/go-steam/v3/protocol/steamlang"
	"github.com/Philipp15b/go-steam/v3/steamid"
	"google.golang.org/protobuf/proto"
)

type Auth struct {
	client  *Client
	details *LogOnDetails
}

type SentryHash []byte

type LogOnDetails struct {
	Username string

	// If logging into an account without a login key, the account's password.
	Password string

	// If you have a Steam Guard email code, you can provide it here.
	AuthCode string

	// If you have a Steam Guard mobile two-factor authentication code, you can provide it here.
	TwoFactorCode  string
	SentryFileHash SentryHash
	LoginKey       string

	// true if you want to get a login key which can be used in lieu of
	// a password for subsequent logins. false or omitted otherwise.
	ShouldRememberPassword bool
}

// Log on with the given details. You must always specify username and
// password OR username and loginkey. For the first login, don't set an authcode or a hash and you'll
//
//	receive an error (EResult_AccountLogonDenied)
//
// and Steam will send you an authcode. Then you have to login again, this time with the authcode.
// Shortly after logging in, you'll receive a MachineAuthUpdateEvent with a hash which allows
// you to login without using an authcode in the future.
//
// If you don't use Steam Guard, username and password are enough.
//
// After the event EMsg_ClientNewLoginKey is received you can use the LoginKey
// to login instead of using the password.
func (a *Auth) LogOn(details *LogOnDetails) {
	if details.Username == "" {
		panic("Username must be set!")
	}
	if details.Password == "" && details.LoginKey == "" {
		panic("Password or LoginKey must be set!")
	}

	logon := new(protobuf.CMsgClientLogon)
	logon.AccountName = &details.Username
	logon.Password = &details.Password
	if details.AuthCode != "" {
		logon.AuthCode = proto.String(details.AuthCode)
	}
	if details.TwoFactorCode != "" {
		logon.TwoFactorCode = proto.String(details.TwoFactorCode)
	}
	logon.ClientLanguage = proto.String("english")
	logon.ProtocolVersion = proto.Uint32(steamlang.MsgClientLogon_CurrentProtocol)
	logon.ShaSentryfile = details.SentryFileHash
	if details.LoginKey != "" {
		logon.LoginKey = proto.String(details.LoginKey)
	}
	if details.ShouldRememberPassword {
		logon.ShouldRememberPassword = proto.Bool(details.ShouldRememberPassword)
	}

	atomic.StoreUint64(&a.client.steamId, uint64(steamid.NewIdAdv(0, 1, int32(steamlang.EUniverse_Public), int32(steamlang.EAccountType_Individual))))

	a.client.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientLogon, logon))
}

func (a *Auth) HandlePacket(packet *protocol.Packet) {
	switch packet.EMsg {
	case steamlang.EMsg_ClientLogOnResponse:
		a.handleLogOnResponse(packet)
	case steamlang.EMsg_ClientNewLoginKey:
		a.handleLoginKey(packet)
	case steamlang.EMsg_ClientSessionToken:
	case steamlang.EMsg_ClientLoggedOff:
		a.handleLoggedOff(packet)
	case steamlang.EMsg_ClientUpdateMachineAuth:
		a.handleUpdateMachineAuth(packet)
	case steamlang.EMsg_ClientAccountInfo:
		a.handleAccountInfo(packet)
	}
}

func (a *Auth) handleLogOnResponse(packet *protocol.Packet) {
	if !packet.IsProto {
		a.client.Fatalf("Got non-proto logon response!")
		return
	}

	body := new(protobuf.CMsgClientLogonResponse)
	msg := packet.ReadProtoMsg(body)

	result := steamlang.EResult(body.GetEresult())
	if result == steamlang.EResult_OK {
		atomic.StoreInt32(&a.client.sessionId, msg.Header.Proto.GetClientSessionid())
		atomic.StoreUint64(&a.client.steamId, msg.Header.Proto.GetSteamid())
		if body.WebapiAuthenticateUserNonce != nil {
			a.client.Web.webLoginKey = *body.WebapiAuthenticateUserNonce
		}

		go a.client.heartbeatLoop(time.Duration(body.GetOutOfGameHeartbeatSeconds()))

		a.client.Emit(&LoggedOnEvent{
			Result:                    steamlang.EResult(body.GetEresult()),
			ExtendedResult:            steamlang.EResult(body.GetEresultExtended()),
			OutOfGameSecsPerHeartbeat: body.GetOutOfGameHeartbeatSeconds(),
			InGameSecsPerHeartbeat:    body.GetInGameHeartbeatSeconds(),
			PublicIp:                  body.GetDeprecatedPublicIp(),
			ServerTime:                body.GetRtime32ServerTime(),
			AccountFlags:              steamlang.EAccountFlags(body.GetAccountFlags()),
			ClientSteamId:             steamid.SteamId(body.GetClientSuppliedSteamid()),
			EmailDomain:               body.GetEmailDomain(),
			CellId:                    body.GetCellId(),
			CellIdPingThreshold:       body.GetCellIdPingThreshold(),
			Steam2Ticket:              body.GetSteam2Ticket(),
			UsePics:                   body.GetDeprecatedUsePics(),
			WebApiUserNonce:           body.GetWebapiAuthenticateUserNonce(),
			IpCountryCode:             body.GetIpCountryCode(),
			VanityUrl:                 body.GetVanityUrl(),
			NumLoginFailuresToMigrate: body.GetCountLoginfailuresToMigrate(),
			NumDisconnectsToMigrate:   body.GetCountDisconnectsToMigrate(),
		})
	} else if result == steamlang.EResult_Fail || result == steamlang.EResult_ServiceUnavailable || result == steamlang.EResult_TryAnotherCM {
		// some error on Steam's side, we'll get an EOF later
	} else {
		a.client.Emit(&LogOnFailedEvent{
			Result: steamlang.EResult(body.GetEresult()),
		})
		a.client.Disconnect()
	}
}

func (a *Auth) handleLoginKey(packet *protocol.Packet) {
	body := new(protobuf.CMsgClientNewLoginKey)
	packet.ReadProtoMsg(body)
	a.client.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientNewLoginKeyAccepted, &protobuf.CMsgClientNewLoginKeyAccepted{
		UniqueId: proto.Uint32(body.GetUniqueId()),
	}))
	a.client.Emit(&LoginKeyEvent{
		UniqueId: body.GetUniqueId(),
		LoginKey: body.GetLoginKey(),
	})
}

func (a *Auth) handleLoggedOff(packet *protocol.Packet) {
	result := steamlang.EResult_Invalid
	if packet.IsProto {
		body := new(protobuf.CMsgClientLoggedOff)
		packet.ReadProtoMsg(body)
		result = steamlang.EResult(body.GetEresult())
	} else {
		body := new(steamlang.MsgClientLoggedOff)
		packet.ReadClientMsg(body)
		result = body.Result
	}
	a.client.Emit(&LoggedOffEvent{Result: result})
}

func (a *Auth) handleUpdateMachineAuth(packet *protocol.Packet) {
	body := new(protobuf.CMsgClientUpdateMachineAuth)
	packet.ReadProtoMsg(body)
	hash := sha1.New()
	hash.Write(packet.Data)
	sha := hash.Sum(nil)

	msg := protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientUpdateMachineAuthResponse, &protobuf.CMsgClientUpdateMachineAuthResponse{
		ShaFile: sha,
	})
	msg.SetTargetJobId(packet.SourceJobId)
	a.client.Write(msg)

	a.client.Emit(&MachineAuthUpdateEvent{sha})
}

func (a *Auth) handleAccountInfo(packet *protocol.Packet) {
	body := new(protobuf.CMsgClientAccountInfo)
	packet.ReadProtoMsg(body)
	a.client.Emit(&AccountInfoEvent{
		PersonaName:          body.GetPersonaName(),
		Country:              body.GetIpCountry(),
		CountAuthedComputers: body.GetCountAuthedComputers(),
		AccountFlags:         steamlang.EAccountFlags(body.GetAccountFlags()),
		FacebookId:           body.GetFacebookId(),
		FacebookName:         body.GetFacebookName(),
	})
}
