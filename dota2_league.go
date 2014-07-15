package steam

import (
	// . "github.com/smithfox/go-steam/internal"
	// . "github.com/smithfox/go-steam/internal/protobuf"
	. "github.com/smithfox/go-steam/internal/gamecoordinator"
	. "github.com/smithfox/go-steam/internal/protobuf/dota"
	// steamlang "github.com/smithfox/go-steam/internal/steamlang"
	// . "github.com/smithfox/go-steam/steamid"
	"fmt"
	"github.com/smithfox/goprotobuf/proto"
)

type Dota2League struct {
	//mutex sync.RWMutex
	dota2client *Dota2Client
}

func NewDota2League(client *Dota2Client) *Dota2League {
	return &Dota2League{
		dota2client: client,
	}
}

func (s *Dota2League) HandleGCPacket(packet *GCPacket) {
	fmt.Printf("Dota2League HandleGCPacket\n")
	switch packet.MsgType {
	case uint32(EDOTAGCMsg_k_EMsgGCRequestMatchesResponse):
		s.handleMatchesResponse(packet)
	}
}

func (s *Dota2League) handleMatchesResponse(packet *GCPacket) {
	list := new(CMsgDOTARequestMatchesResponse)
	packet.ReadProtoMsg(list)
	fmt.Printf("handleMatchesResponse %s\n", list.String())
	for _, match := range list.GetMatches() {
		mid := match.GetMatchId()
		fmt.Printf("Dota2League mid=%d\n", mid)
	}
	for _, series := range list.GetSeries() {
		seriesid := series.GetSeriesId()
		fmt.Printf("Dota2League seriesid=%d\n", seriesid)
	}
	s.dota2client.Emit(new(Dota2LeagueEvent))
}

type Dota2LeagueEvent struct{}

func (s *Dota2League) RequestMatchs(leagueid uint32) {
	s.dota2client.gc.Write(NewGCMsgProtobuf(DOTA2_APP_ID, uint32(EDOTAGCMsg_k_EMsgGCRequestMatches), &CMsgDOTARequestMatches{
		LeagueId:         proto.Uint32(leagueid),
		MatchesRequested: proto.Uint32(20),
	}))
}
