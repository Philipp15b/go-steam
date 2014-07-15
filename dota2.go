package steam

import (
	"fmt"
	. "github.com/smithfox/go-steam/internal/gamecoordinator"
	. "github.com/smithfox/go-steam/internal/protobuf/dota"
	"time"
)

const DOTA2_APP_ID = 570

type Dota2Client struct {
	gc          *GameCoordinator
	accountId   uint32
	gcReady     bool
	helloTicket *time.Ticker
	League      *Dota2League
}

func ToAccountId(steamid uint64) uint32 {
	return uint32(steamid - 76561197960265728)
}

func NewDota2Client(gc *GameCoordinator) *Dota2Client {
	dota2client := &Dota2Client{
		gc: gc,
	}
	dota2client.accountId = ToAccountId(uint64(gc.client.SteamId()))
	gc.RegisterPacketHandler(dota2client)

	dota2client.League = NewDota2League(dota2client)
	gc.RegisterPacketHandler(dota2client.League)
	return dota2client
}

func (d *Dota2Client) Launch() {
	d.gc.SetGamesPlayed(DOTA2_APP_ID)
	go d.SendClientHello()
}

func (d *Dota2Client) Exit() {
	if d.helloTicket != nil {
		d.helloTicket.Stop()
	}

	d.gcReady = false
	d.gc.SetGamesPlayed()
}

type Dota2ClientReadyEvent struct {
	Ready bool
}

func (d *Dota2Client) Emit(event interface{}) {
	d.gc.client.Emit(event)
}

func (d *Dota2Client) HandleGCPacket(packet *GCPacket) {
	fmt.Printf("Dota2Client HandleGCPacket\n")
	switch packet.MsgType {
	case uint32(EGCBaseClientMsg_k_EMsgGCClientWelcome):
		if d.helloTicket != nil {
			d.helloTicket.Stop()
		}

		d.gcReady = true
		d.Emit(&Dota2ClientReadyEvent{Ready: d.gcReady})
	case uint32(EGCBaseClientMsg_k_EMsgGCClientConnectionStatus):
		body := new(CMsgConnectionStatus)
		packet.ReadProtoMsg(body)

		switch body.GetStatus() {
		case GCConnectionStatus_GCConnectionStatus_HAVE_SESSION:
			//log("GC Connection Status regained.");
			if d.helloTicket != nil {
				d.helloTicket.Stop()
			}

			d.Emit(&Dota2ClientReadyEvent{Ready: d.gcReady})
		default:
			//log("GC Connection Status unreliable - " + status);
			if d.helloTicket == nil {
				go d.SendClientHello()
				d.Emit(&Dota2ClientReadyEvent{Ready: false})
			}
		}
	}
}

func (d *Dota2Client) _sendClientHello() {
	d.gc.Write(NewGCMsgProtobuf(DOTA2_APP_ID, uint32(EGCBaseClientMsg_k_EMsgGCClientHello), new(CMsgClientHello)))
}

func (d *Dota2Client) SendClientHello() {
	if d.helloTicket != nil {
		d.helloTicket.Stop()
	}
	d.gcReady = false
	d.helloTicket = time.NewTicker(3 * time.Second)
	for {
		_, ok := <-d.helloTicket.C
		if !ok {
			break
		}
		d._sendClientHello()
	}
	d.helloTicket = nil
}
