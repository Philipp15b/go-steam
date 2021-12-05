package steam

import (
	"time"

	"github.com/Philipp15b/go-steam/v2/protocol/protobuf"
	"github.com/Philipp15b/go-steam/v2/protocol/steamlang"
	"github.com/Philipp15b/go-steam/v2/steamid"
)

type FriendsListEvent struct{}

type FriendStateEvent struct {
	SteamId      steamid.SteamId `json:",string"`
	Relationship steamlang.EFriendRelationship
}

func (f *FriendStateEvent) IsFriend() bool {
	return f.Relationship == steamlang.EFriendRelationship_Friend
}

type GroupStateEvent struct {
	SteamId      steamid.SteamId `json:",string"`
	Relationship steamlang.EClanRelationship
}

func (g *GroupStateEvent) IsMember() bool {
	return g.Relationship == steamlang.EClanRelationship_Member
}

// Fired when someone changing their friend details
type PersonaStateEvent struct {
	StatusFlags            steamlang.EClientPersonaStateFlag
	FriendId               steamid.SteamId `json:",string"`
	State                  steamlang.EPersonaState
	StateFlags             steamlang.EPersonaStateFlag
	GameAppId              uint32
	GameId                 uint64 `json:",string"`
	GameName               string
	GameServerIp           uint32
	GameServerPort         uint32
	QueryPort              uint32
	SourceSteamId          steamid.SteamId `json:",string"`
	GameDataBlob           []byte
	Name                   string
	Avatar                 []byte
	LastLogOff             uint32
	LastLogOn              uint32
	ClanRank               uint32
	ClanTag                string
	OnlineSessionInstances uint32
	PersonaSetByUser       bool
	RichPresence           []*protobuf.CMsgClientPersonaState_Friend_KV
}

// Fired when a clan's state has been changed
type ClanStateEvent struct {
	ClanId              steamid.SteamId `json:",string"`
	AccountFlags        steamlang.EAccountFlags
	ClanName            string
	Avatar              []byte
	MemberTotalCount    uint32
	MemberOnlineCount   uint32
	MemberChattingCount uint32
	MemberInGameCount   uint32
	Events              []ClanEventDetails
	Announcements       []ClanEventDetails
}

type ClanEventDetails struct {
	Id         uint64 `json:",string"`
	EventTime  uint32
	Headline   string
	GameId     uint64 `json:",string"`
	JustPosted bool
}

// Fired in response to adding a friend to your friends list
type FriendAddedEvent struct {
	Result      steamlang.EResult
	SteamId     steamid.SteamId `json:",string"`
	PersonaName string
}

// Fired when the client receives a message from either a friend or a chat room
type ChatMsgEvent struct {
	ChatRoomId steamid.SteamId `json:",string"` // not set for friend messages
	ChatterId  steamid.SteamId `json:",string"`
	Message    string
	EntryType  steamlang.EChatEntryType
	Timestamp  time.Time
	Offline    bool
}

// Whether the type is ChatMsg
func (c *ChatMsgEvent) IsMessage() bool {
	return c.EntryType == steamlang.EChatEntryType_ChatMsg
}

// Fired in response to joining a chat
type ChatEnterEvent struct {
	ChatRoomId    steamid.SteamId `json:",string"`
	FriendId      steamid.SteamId `json:",string"`
	ChatRoomType  steamlang.EChatRoomType
	OwnerId       steamid.SteamId `json:",string"`
	ClanId        steamid.SteamId `json:",string"`
	ChatFlags     byte
	EnterResponse steamlang.EChatRoomEnterResponse
	Name          string
}

// Fired in response to a chat member's info being received
type ChatMemberInfoEvent struct {
	ChatRoomId      steamid.SteamId `json:",string"`
	Type            steamlang.EChatInfoType
	StateChangeInfo StateChangeDetails
}

type StateChangeDetails struct {
	ChatterActedOn steamid.SteamId `json:",string"`
	StateChange    steamlang.EChatMemberStateChange
	ChatterActedBy steamid.SteamId `json:",string"`
}

// Fired when a chat action has completed
type ChatActionResultEvent struct {
	ChatRoomId steamid.SteamId `json:",string"`
	ChatterId  steamid.SteamId `json:",string"`
	Action     steamlang.EChatAction
	Result     steamlang.EChatActionResult
}

// Fired when a chat invite is received
type ChatInviteEvent struct {
	InvitedId    steamid.SteamId `json:",string"`
	ChatRoomId   steamid.SteamId `json:",string"`
	PatronId     steamid.SteamId `json:",string"`
	ChatRoomType steamlang.EChatRoomType
	FriendChatId steamid.SteamId `json:",string"`
	ChatRoomName string
	GameId       uint64 `json:",string"`
}

// Fired in response to ignoring a friend
type IgnoreFriendEvent struct {
	Result steamlang.EResult
}

// Fired in response to requesting profile info for a user
type ProfileInfoEvent struct {
	Result      steamlang.EResult
	SteamId     steamid.SteamId `json:",string"`
	TimeCreated uint32
	RealName    string
	CityName    string
	StateName   string
	CountryName string
	Headline    string
	Summary     string
}
