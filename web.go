package steam

import (
	"github.com/smithfox/goprotobuf/proto"
	"crypto/aes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"github.com/smithfox/go-steam/cryptoutil"
	. "github.com/smithfox/go-steam/internal"
	. "github.com/smithfox/go-steam/internal/protobuf"
	. "github.com/smithfox/go-steam/internal/steamlang"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"sync/atomic"
)

type Web struct {
	// The `sessionid` cookie required to use the steam website.
	WebSessionId string
	// The `steamLogin` cookie required to use the steam website.
	// It is only available after calling LogOn().
	SteamLogin string

	webLoginKey  string
	relogOnNonce uint32

	client *Client
}

func (w *Web) HandlePacket(packet *Packet) {
	switch packet.EMsg {
	case EMsg_ClientNewLoginKey:
		w.handleNewLoginKey(packet)
	case EMsg_ClientRequestWebAPIAuthenticateUserNonceResponse:
		w.handleAuthNonceResponse(packet)
	}
}

type WebLoggedOnEvent struct{}

// Fetches the `steamLogin` cookie. This may only be called after the first
// WebSessionIdEvent or it will panic.
func (w *Web) LogOn() {
	if w.webLoginKey == "" {
		panic("SteamWeb: webLoginKey not initialized!")
	}

	go func() {
		// retry three times. yes, I know about loops.
		err := w.apiLogOn()
		if err != nil {
			err = w.apiLogOn()
			if err != nil {
				err = w.apiLogOn()
			}
		}
		if err != nil {
			w.client.Errorf("web: Error logging on: %v", err)
			return
		}
	}()
}

func (w *Web) apiLogOn() error {
	sessionKey := make([]byte, 32)
	rand.Read(sessionKey)

	cryptedSessionKey := cryptoutil.RSAEncrypt(GetPublicKey(EUniverse_Public), sessionKey)
	ciph, _ := aes.NewCipher(sessionKey)
	cryptedLoginKey := cryptoutil.SymmetricEncrypt(ciph, []byte(w.WebSessionId))
	data := make(url.Values)
	data.Add("format", "json")
	data.Add("steamid", strconv.FormatUint(uint64(w.client.SteamId()), 10))
	data.Add("sessionkey", string(cryptedSessionKey))
	data.Add("encrypted_loginkey", string(cryptedLoginKey))
	resp, err := http.PostForm("http://api.steampowered.com/ISteamUserAuth/AuthenticateUser/v0001", data)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		// our web session id has expired, request a new one
		w.client.Write(NewClientMsgProtobuf(EMsg_ClientRequestWebAPIAuthenticateUserNonce, new(CMsgClientRequestWebAPIAuthenticateUserNonce)))
//		atomic.StoreUint32(&w.relogOnNonce, 1)
		w.relogOnNonce = 1
		return nil
	}

	result := new(struct {
		Authenticateuser struct {
			Token string
		}
	})
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, result)
	if err != nil {
		return err
	}

	w.SteamLogin = result.Authenticateuser.Token

	w.client.Emit(new(WebLoggedOnEvent))
	return nil
}

type WebSessionIdEvent struct{}

func (w *Web) handleNewLoginKey(packet *Packet) {
	msg := new(CMsgClientNewLoginKey)
	packet.ReadProtoMsg(msg)

	w.client.Write(NewClientMsgProtobuf(EMsg_ClientNewLoginKeyAccepted, &CMsgClientNewLoginKeyAccepted{
		UniqueId: proto.Uint32(msg.GetUniqueId()),
	}))

	w.webLoginKey = msg.GetLoginKey()
	// number -> string -> bytes -> base64
	w.WebSessionId = base64.StdEncoding.EncodeToString([]byte(strconv.FormatUint(uint64(msg.GetUniqueId()), 10)))

	w.client.Emit(new(WebSessionIdEvent))
}

func (w *Web) handleAuthNonceResponse(packet *Packet) {
	// this has to be the best name for a message yet.
	msg := new(CMsgClientRequestWebAPIAuthenticateUserNonceResponse)
	packet.ReadProtoMsg(msg)
	w.WebSessionId = msg.GetWebapiAuthenticateUserNonce()

	// if the nonce was specifically requested in apiLogOn(),
	// don't emit an event.
	if atomic.CompareAndSwapUint32(&w.relogOnNonce, 1, 0) {
		w.LogOn()
	} else {
		w.client.Emit(new(WebSessionIdEvent))
	}
}
