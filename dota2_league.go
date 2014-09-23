package steam

import (
	"code.google.com/p/goprotobuf/proto"
	"fmt"
	. "github.com/Philipp15b/go-steam/internal/gamecoordinator"
	. "github.com/Philipp15b/go-steam/internal/protobuf/dota"
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
	case uint32(EDOTAGCMsg_k_EMsgGCLeagueScheduleResponse):
		s.handleLeagueScheduleResponse(packet)
	}
}

func (s *Dota2League) handleMatchesResponse(packet *GCPacket) {
	list := new(CMsgDOTARequestMatchesResponse)
	packet.ReadProtoMsg(list)

	for _, match := range list.GetMatches() {
		//		mid := match.GetMatchId()
		fmt.Printf("---------------------- match ---------------\n")
		fmt.Printf("Mid=%d\n", match.GetMatchId())
		fmt.Printf("Leagueid=%v\n", match.GetLeagueid())
		fmt.Printf("StartTime=%d\n", match.GetStartTime())
		fmt.Printf("RadiantTeamId=%v\n", match.GetRadiantTeamId())
		fmt.Printf("DireTeamId=%v\n", match.GetDireTeamId())
		fmt.Printf("RadiantTeamName=%v\n", match.GetRadiantTeamName())
		fmt.Printf("DireTeamName=%v\n", match.GetDireTeamName())
		fmt.Printf("RadiantTeamTag=%v\n", match.GetRadiantTeamTag())
		fmt.Printf("DireTeamTag=%v\n", match.GetDireTeamTag())
		fmt.Printf("---------------------- end match ---------------\n")
	}
	for _, series := range list.GetSeries() {
		seriesid := series.GetSeriesId()
		seriestype := series.GetSeriesType()
		fmt.Printf("Dota2League seriesid=%d,seriestype=%d\n", seriesid, seriestype)
		for _, match := range series.GetMatches() {
			fmt.Printf("---------------------- match ---------------\n")
			fmt.Printf("Mid=%d\n", match.GetMatchId())
			fmt.Printf("Leagueid=%v\n", match.GetLeagueid())
			fmt.Printf("StartTime=%d\n", match.GetStartTime())
			fmt.Printf("RadiantTeamId=%v\n", match.GetRadiantTeamId())
			fmt.Printf("DireTeamId=%v\n", match.GetDireTeamId())
			fmt.Printf("RadiantTeamName=%v\n", match.GetRadiantTeamName())
			fmt.Printf("DireTeamName=%v\n", match.GetDireTeamName())
			fmt.Printf("RadiantTeamTag=%v\n", match.GetRadiantTeamTag())
			fmt.Printf("DireTeamTag=%v\n", match.GetDireTeamTag())
			fmt.Printf("---------------------- end match ---------------\n")
		}
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

func (s *Dota2League) RequestSchedules(leagueid uint32) {
	s.dota2client.gc.Write(NewGCMsgProtobuf(
		DOTA2_APP_ID,
		uint32(EDOTAGCMsg_k_EMsgGCLeagueScheduleRequest),
		&CMsgDOTALeagueScheduleRequest{
			LeagueId: proto.Uint32(leagueid),
		}),
	)
}

func (s *Dota2League) handleLeagueScheduleResponse(packet *GCPacket) {
	list := new(CMsgDOTALeagueScheduleResponse)
	packet.ReadProtoMsg(list)

	league := list.GetLeague()
	if league == nil {
		fmt.Printf("response schedule no league")
		return
	}
	for _, schedule := range league.GetSchedule() {
		fmt.Printf("-------------- schedule -------------\n")
		fmt.Printf("BlockId=%v\n", schedule.GetBlockId())
		fmt.Printf("StartTime=%v\n", schedule.GetStartTime())
		fmt.Printf("Finals=%v\n", schedule.GetFinals())
		fmt.Printf("Comment=%v\n", schedule.GetComment())
		teams := schedule.GetTeams()
		for _, team := range teams {
			fmt.Printf("TeamId=%v\n", team.GetTeamId())
			fmt.Printf("TeamName=%v\n", team.GetName())
			fmt.Printf("TeamLogo=%v\n", team.GetLogo())
		}
		fmt.Printf("-------------- end schedule -------------\n")
	}
}
