package steam

import (
	"bytes"
	"encoding/binary"
	"io"
	"sync"
	"time"

	"github.com/Philipp15b/go-steam/v2/protocol"
	"github.com/Philipp15b/go-steam/v2/protocol/protobuf"
	"github.com/Philipp15b/go-steam/v2/protocol/steamlang"
	"github.com/Philipp15b/go-steam/v2/rwu"
	"github.com/Philipp15b/go-steam/v2/socialcache"
	"github.com/Philipp15b/go-steam/v2/steamid"
	"google.golang.org/protobuf/proto"
)

// Provides access to social aspects of Steam.
type Social struct {
	mutex sync.RWMutex

	name         string
	avatar       []byte
	personaState steamlang.EPersonaState

	Friends *socialcache.FriendsList
	Groups  *socialcache.GroupsList
	Chats   *socialcache.ChatsList

	client *Client
}

func newSocial(client *Client) *Social {
	return &Social{
		Friends: socialcache.NewFriendsList(),
		Groups:  socialcache.NewGroupsList(),
		Chats:   socialcache.NewChatsList(),
		client:  client,
	}
}

// Gets the local user's avatar
func (s *Social) GetAvatar() []byte {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.avatar
}

// Gets the local user's persona name
func (s *Social) GetPersonaName() string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.name
}

// Sets the local user's persona name and broadcasts it over the network
func (s *Social) SetPersonaName(name string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.name = name
	s.client.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientChangeStatus, &protobuf.CMsgClientChangeStatus{
		PersonaState: proto.Uint32(uint32(s.personaState)),
		PlayerName:   proto.String(name),
	}))
}

// Gets the local user's persona state
func (s *Social) GetPersonaState() steamlang.EPersonaState {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.personaState
}

// Sets the local user's persona state and broadcasts it over the network
func (s *Social) SetPersonaState(state steamlang.EPersonaState) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.personaState = state
	s.client.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientChangeStatus, &protobuf.CMsgClientChangeStatus{
		PersonaState: proto.Uint32(uint32(state)),
	}))
}

// Sends a chat message to ether a room or friend
func (s *Social) SendMessage(to steamid.SteamId, entryType steamlang.EChatEntryType, message string) {
	// Friend
	if to.GetAccountType() == int32(steamlang.EAccountType_Individual) || to.GetAccountType() == int32(steamlang.EAccountType_ConsoleUser) {
		s.client.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientFriendMsg, &protobuf.CMsgClientFriendMsg{
			Steamid:       proto.Uint64(to.ToUint64()),
			ChatEntryType: proto.Int32(int32(entryType)),
			Message:       []byte(message),
		}))
		// Chat room
	} else if to.GetAccountType() == int32(steamlang.EAccountType_Clan) || to.GetAccountType() == int32(steamlang.EAccountType_Chat) {
		chatId := to.ClanToChat()
		s.client.Write(protocol.NewClientMsg(&steamlang.MsgClientChatMsg{
			ChatMsgType:     entryType,
			SteamIdChatRoom: chatId,
			SteamIdChatter:  s.client.SteamId(),
		}, []byte(message)))
	}
}

// Adds a friend to your friends list or accepts a friend. You'll receive a FriendStateEvent
// for every new/changed friend
func (s *Social) AddFriend(id steamid.SteamId) {
	s.client.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientAddFriend, &protobuf.CMsgClientAddFriend{
		SteamidToAdd: proto.Uint64(id.ToUint64()),
	}))
}

// Removes a friend from your friends list
func (s *Social) RemoveFriend(id steamid.SteamId) {
	s.client.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientRemoveFriend, &protobuf.CMsgClientRemoveFriend{
		Friendid: proto.Uint64(id.ToUint64()),
	}))
}

// Ignores or unignores a friend on Steam
func (s *Social) IgnoreFriend(id steamid.SteamId, setIgnore bool) {
	ignore := uint8(1) // True
	if !setIgnore {
		ignore = uint8(0) // False
	}
	s.client.Write(protocol.NewClientMsg(&steamlang.MsgClientSetIgnoreFriend{
		MySteamId:     s.client.SteamId(),
		SteamIdFriend: id,
		Ignore:        ignore,
	}, make([]byte, 0)))
}

// Requests persona state for a list of specified SteamIds
func (s *Social) RequestFriendListInfo(ids []steamid.SteamId, requestedInfo steamlang.EClientPersonaStateFlag) {
	var friends []uint64
	for _, id := range ids {
		friends = append(friends, id.ToUint64())
	}
	s.client.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientRequestFriendData, &protobuf.CMsgClientRequestFriendData{
		PersonaStateRequested: proto.Uint32(uint32(requestedInfo)),
		Friends:               friends,
	}))
}

// Requests persona state for a specified SteamId
func (s *Social) RequestFriendInfo(id steamid.SteamId, requestedInfo steamlang.EClientPersonaStateFlag) {
	s.RequestFriendListInfo([]steamid.SteamId{id}, requestedInfo)
}

// Requests profile information for a specified SteamId
func (s *Social) RequestProfileInfo(id steamid.SteamId) {
	s.client.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientFriendProfileInfo, &protobuf.CMsgClientFriendProfileInfo{
		SteamidFriend: proto.Uint64(id.ToUint64()),
	}))
}

// Requests all offline messages and marks them as read
func (s *Social) RequestOfflineMessages() {
	s.client.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientChatGetFriendMessageHistoryForOfflineMessages, &protobuf.CMsgClientChatGetFriendMessageHistoryForOfflineMessages{}))
}

// Attempts to join a chat room
func (s *Social) JoinChat(id steamid.SteamId) {
	chatId := id.ClanToChat()
	s.client.Write(protocol.NewClientMsg(&steamlang.MsgClientJoinChat{
		SteamIdChat: chatId,
	}, make([]byte, 0)))
}

// Attempts to leave a chat room
func (s *Social) LeaveChat(id steamid.SteamId) {
	chatId := id.ClanToChat()
	payload := new(bytes.Buffer)
	binary.Write(payload, binary.LittleEndian, s.client.SteamId().ToUint64())                 // ChatterActedOn
	binary.Write(payload, binary.LittleEndian, uint32(steamlang.EChatMemberStateChange_Left)) // StateChange
	binary.Write(payload, binary.LittleEndian, s.client.SteamId().ToUint64())                 // ChatterActedBy
	s.client.Write(protocol.NewClientMsg(&steamlang.MsgClientChatMemberInfo{
		SteamIdChat: chatId,
		Type:        steamlang.EChatInfoType_StateChange,
	}, payload.Bytes()))
}

// Kicks the specified chat member from the given chat room
func (s *Social) KickChatMember(room steamid.SteamId, user steamid.SteamId) {
	chatId := room.ClanToChat()
	s.client.Write(protocol.NewClientMsg(&steamlang.MsgClientChatAction{
		SteamIdChat:        chatId,
		SteamIdUserToActOn: user,
		ChatAction:         steamlang.EChatAction_Kick,
	}, make([]byte, 0)))
}

// Bans the specified chat member from the given chat room
func (s *Social) BanChatMember(room steamid.SteamId, user steamid.SteamId) {
	chatId := room.ClanToChat()
	s.client.Write(protocol.NewClientMsg(&steamlang.MsgClientChatAction{
		SteamIdChat:        chatId,
		SteamIdUserToActOn: user,
		ChatAction:         steamlang.EChatAction_Ban,
	}, make([]byte, 0)))
}

// Unbans the specified chat member from the given chat room
func (s *Social) UnbanChatMember(room steamid.SteamId, user steamid.SteamId) {
	chatId := room.ClanToChat()
	s.client.Write(protocol.NewClientMsg(&steamlang.MsgClientChatAction{
		SteamIdChat:        chatId,
		SteamIdUserToActOn: user,
		ChatAction:         steamlang.EChatAction_UnBan,
	}, make([]byte, 0)))
}

func (s *Social) HandlePacket(packet *protocol.Packet) {
	switch packet.EMsg {
	case steamlang.EMsg_ClientPersonaState:
		s.handlePersonaState(packet)
	case steamlang.EMsg_ClientClanState:
		s.handleClanState(packet)
	case steamlang.EMsg_ClientFriendsList:
		s.handleFriendsList(packet)
	case steamlang.EMsg_ClientFriendMsgIncoming:
		s.handleFriendMsg(packet)
	case steamlang.EMsg_ClientAccountInfo:
		s.handleAccountInfo(packet)
	case steamlang.EMsg_ClientAddFriendResponse:
		s.handleFriendResponse(packet)
	case steamlang.EMsg_ClientChatEnter:
		s.handleChatEnter(packet)
	case steamlang.EMsg_ClientChatMsg:
		s.handleChatMsg(packet)
	case steamlang.EMsg_ClientChatMemberInfo:
		s.handleChatMemberInfo(packet)
	case steamlang.EMsg_ClientChatActionResult:
		s.handleChatActionResult(packet)
	case steamlang.EMsg_ClientChatInvite:
		s.handleChatInvite(packet)
	case steamlang.EMsg_ClientSetIgnoreFriendResponse:
		s.handleIgnoreFriendResponse(packet)
	case steamlang.EMsg_ClientFriendProfileInfoResponse:
		s.handleProfileInfoResponse(packet)
	case steamlang.EMsg_ClientFSGetFriendMessageHistoryResponse:
		s.handleFriendMessageHistoryResponse(packet)
	}
}

func (s *Social) handleAccountInfo(packet *protocol.Packet) {
	// Just fire the personainfo, Auth handles the callback
	flags := steamlang.EClientPersonaStateFlag_PlayerName | steamlang.EClientPersonaStateFlag_Presence | steamlang.EClientPersonaStateFlag_SourceID
	s.RequestFriendInfo(s.client.SteamId(), steamlang.EClientPersonaStateFlag(flags))
}

func (s *Social) handleFriendsList(packet *protocol.Packet) {
	list := new(protobuf.CMsgClientFriendsList)
	packet.ReadProtoMsg(list)
	var friends []steamid.SteamId
	for _, friend := range list.GetFriends() {
		steamId := steamid.SteamId(friend.GetUlfriendid())
		isClan := steamId.GetAccountType() == int32(steamlang.EAccountType_Clan)

		if isClan {
			rel := steamlang.EClanRelationship(friend.GetEfriendrelationship())
			if rel == steamlang.EClanRelationship_None {
				s.Groups.Remove(steamId)
			} else {
				s.Groups.Add(socialcache.Group{
					SteamId:      steamId,
					Relationship: rel,
				})

			}
			if list.GetBincremental() {
				s.client.Emit(&GroupStateEvent{steamId, rel})
			}
		} else {
			rel := steamlang.EFriendRelationship(friend.GetEfriendrelationship())
			if rel == steamlang.EFriendRelationship_None {
				s.Friends.Remove(steamId)
			} else {
				s.Friends.Add(socialcache.Friend{
					SteamId:      steamId,
					Relationship: rel,
				})

			}
			if list.GetBincremental() {
				s.client.Emit(&FriendStateEvent{steamId, rel})
			}
		}
		if !list.GetBincremental() {
			friends = append(friends, steamId)
		}
	}
	if !list.GetBincremental() {
		s.RequestFriendListInfo(friends, protocol.EClientPersonaStateFlag_DefaultInfoRequest)
		s.client.Emit(&FriendsListEvent{})
	}
}

func (s *Social) handlePersonaState(packet *protocol.Packet) {
	list := new(protobuf.CMsgClientPersonaState)
	packet.ReadProtoMsg(list)
	flags := steamlang.EClientPersonaStateFlag(list.GetStatusFlags())
	for _, friend := range list.GetFriends() {
		id := steamid.SteamId(friend.GetFriendid())
		if id == s.client.SteamId() { // this is our client id
			s.mutex.Lock()
			if friend.GetPlayerName() != "" {
				s.name = friend.GetPlayerName()
			}
			avatar := friend.GetAvatarHash()
			if protocol.ValidAvatar(avatar) {
				s.avatar = avatar
			}
			s.mutex.Unlock()
		} else if id.GetAccountType() == int32(steamlang.EAccountType_Individual) {
			if (flags & steamlang.EClientPersonaStateFlag_PlayerName) == steamlang.EClientPersonaStateFlag_PlayerName {
				if friend.GetPlayerName() != "" {
					s.Friends.SetName(id, friend.GetPlayerName())
				}
			}
			if (flags & steamlang.EClientPersonaStateFlag_Presence) == steamlang.EClientPersonaStateFlag_Presence {
				avatar := friend.GetAvatarHash()
				if protocol.ValidAvatar(avatar) {
					s.Friends.SetAvatar(id, avatar)
				}
				s.Friends.SetPersonaState(id, steamlang.EPersonaState(friend.GetPersonaState()))
				s.Friends.SetPersonaStateFlags(id, steamlang.EPersonaStateFlag(friend.GetPersonaStateFlags()))
			}
			if (flags & steamlang.EClientPersonaStateFlag_GameDataBlob) == steamlang.EClientPersonaStateFlag_GameDataBlob {
				s.Friends.SetGameAppId(id, friend.GetGamePlayedAppId())
				s.Friends.SetGameId(id, friend.GetGameid())
				s.Friends.SetGameName(id, friend.GetGameName())
			}
		} else if id.GetAccountType() == int32(steamlang.EAccountType_Clan) {
			if (flags & steamlang.EClientPersonaStateFlag_PlayerName) == steamlang.EClientPersonaStateFlag_PlayerName {
				if friend.GetPlayerName() != "" {
					s.Groups.SetName(id, friend.GetPlayerName())
				}
			}
			if (flags & steamlang.EClientPersonaStateFlag_Presence) == steamlang.EClientPersonaStateFlag_Presence {
				avatar := friend.GetAvatarHash()
				if protocol.ValidAvatar(avatar) {
					s.Groups.SetAvatar(id, avatar)
				}
			}
		}
		s.client.Emit(&PersonaStateEvent{
			StatusFlags:            flags,
			FriendId:               id,
			State:                  steamlang.EPersonaState(friend.GetPersonaState()),
			StateFlags:             steamlang.EPersonaStateFlag(friend.GetPersonaStateFlags()),
			GameAppId:              friend.GetGamePlayedAppId(),
			GameId:                 friend.GetGameid(),
			GameName:               friend.GetGameName(),
			GameServerIp:           friend.GetGameServerIp(),
			GameServerPort:         friend.GetGameServerPort(),
			QueryPort:              friend.GetQueryPort(),
			SourceSteamId:          steamid.SteamId(friend.GetSteamidSource()),
			GameDataBlob:           friend.GetGameDataBlob(),
			Name:                   friend.GetPlayerName(),
			Avatar:                 friend.GetAvatarHash(),
			LastLogOff:             friend.GetLastLogoff(),
			LastLogOn:              friend.GetLastLogon(),
			ClanRank:               friend.GetClanRank(),
			ClanTag:                friend.GetClanTag(),
			OnlineSessionInstances: friend.GetOnlineSessionInstances(),
			PersonaSetByUser:       friend.GetPersonaSetByUser(),
		})
	}
}

func (s *Social) handleClanState(packet *protocol.Packet) {
	body := new(protobuf.CMsgClientClanState)
	packet.ReadProtoMsg(body)
	var name string
	var avatar []byte
	if body.GetNameInfo() != nil {
		name = body.GetNameInfo().GetClanName()
		avatar = body.GetNameInfo().GetShaAvatar()
	}
	var totalCount, onlineCount, chattingCount, ingameCount uint32
	if body.GetUserCounts() != nil {
		usercounts := body.GetUserCounts()
		totalCount = usercounts.GetMembers()
		onlineCount = usercounts.GetOnline()
		chattingCount = usercounts.GetChatting()
		ingameCount = usercounts.GetInGame()
	}
	var events, announcements []ClanEventDetails
	for _, event := range body.GetEvents() {
		events = append(events, ClanEventDetails{
			Id:         event.GetGid(),
			EventTime:  event.GetEventTime(),
			Headline:   event.GetHeadline(),
			GameId:     event.GetGameId(),
			JustPosted: event.GetJustPosted(),
		})
	}
	for _, announce := range body.GetAnnouncements() {
		announcements = append(announcements, ClanEventDetails{
			Id:         announce.GetGid(),
			EventTime:  announce.GetEventTime(),
			Headline:   announce.GetHeadline(),
			GameId:     announce.GetGameId(),
			JustPosted: announce.GetJustPosted(),
		})
	}

	// Add stuff to group
	clanid := steamid.SteamId(body.GetSteamidClan())
	if body.NameInfo != nil {
		info := body.NameInfo
		s.Groups.SetName(clanid, info.GetClanName())
		s.Groups.SetAvatar(clanid, info.GetShaAvatar())
	}
	if body.GetUserCounts() != nil {
		s.Groups.SetMemberTotalCount(clanid, totalCount)
		s.Groups.SetMemberOnlineCount(clanid, onlineCount)
		s.Groups.SetMemberChattingCount(clanid, chattingCount)
		s.Groups.SetMemberInGameCount(clanid, ingameCount)
	}
	s.client.Emit(&ClanStateEvent{
		ClanId:              clanid,
		AccountFlags:        steamlang.EAccountFlags(body.GetClanAccountFlags()),
		ClanName:            name,
		Avatar:              avatar,
		MemberTotalCount:    totalCount,
		MemberOnlineCount:   onlineCount,
		MemberChattingCount: chattingCount,
		MemberInGameCount:   ingameCount,
		Events:              events,
		Announcements:       announcements,
	})
}

func (s *Social) handleFriendResponse(packet *protocol.Packet) {
	body := new(protobuf.CMsgClientAddFriendResponse)
	packet.ReadProtoMsg(body)
	s.client.Emit(&FriendAddedEvent{
		Result:      steamlang.EResult(body.GetEresult()),
		SteamId:     steamid.SteamId(body.GetSteamIdAdded()),
		PersonaName: body.GetPersonaNameAdded(),
	})
}

func (s *Social) handleFriendMsg(packet *protocol.Packet) {
	body := new(protobuf.CMsgClientFriendMsgIncoming)
	packet.ReadProtoMsg(body)
	message := string(bytes.Split(body.GetMessage(), []byte{0x0})[0])
	s.client.Emit(&ChatMsgEvent{
		ChatterId: steamid.SteamId(body.GetSteamidFrom()),
		Message:   message,
		EntryType: steamlang.EChatEntryType(body.GetChatEntryType()),
		Timestamp: time.Unix(int64(body.GetRtime32ServerTimestamp()), 0),
	})
}

func (s *Social) handleChatMsg(packet *protocol.Packet) {
	body := new(steamlang.MsgClientChatMsg)
	payload := packet.ReadClientMsg(body).Payload
	message := string(bytes.Split(payload, []byte{0x0})[0])
	s.client.Emit(&ChatMsgEvent{
		ChatRoomId: steamid.SteamId(body.SteamIdChatRoom),
		ChatterId:  steamid.SteamId(body.SteamIdChatter),
		Message:    message,
		EntryType:  steamlang.EChatEntryType(body.ChatMsgType),
	})
}

func (s *Social) handleChatEnter(packet *protocol.Packet) {
	body := new(steamlang.MsgClientChatEnter)
	payload := packet.ReadClientMsg(body).Payload
	reader := bytes.NewBuffer(payload)
	name, _ := rwu.ReadString(reader)
	rwu.ReadByte(reader) // 0
	count := body.NumMembers
	chatId := steamid.SteamId(body.SteamIdChat)
	clanId := steamid.SteamId(body.SteamIdClan)
	s.Chats.Add(socialcache.Chat{SteamId: chatId, GroupId: clanId})
	for i := 0; i < int(count); i++ {
		id, chatPerm, clanPerm := readChatMember(reader)
		rwu.ReadBytes(reader, 6) // No idea what this is
		s.Chats.AddChatMember(chatId, socialcache.ChatMember{
			SteamId:         steamid.SteamId(id),
			ChatPermissions: chatPerm,
			ClanPermissions: clanPerm,
		})
	}
	s.client.Emit(&ChatEnterEvent{
		ChatRoomId:    steamid.SteamId(body.SteamIdChat),
		FriendId:      steamid.SteamId(body.SteamIdFriend),
		ChatRoomType:  steamlang.EChatRoomType(body.ChatRoomType),
		OwnerId:       steamid.SteamId(body.SteamIdOwner),
		ClanId:        steamid.SteamId(body.SteamIdClan),
		ChatFlags:     byte(body.ChatFlags),
		EnterResponse: steamlang.EChatRoomEnterResponse(body.EnterResponse),
		Name:          name,
	})
}

func (s *Social) handleChatMemberInfo(packet *protocol.Packet) {
	body := new(steamlang.MsgClientChatMemberInfo)
	payload := packet.ReadClientMsg(body).Payload
	reader := bytes.NewBuffer(payload)
	chatId := steamid.SteamId(body.SteamIdChat)
	if body.Type == steamlang.EChatInfoType_StateChange {
		actedOn, _ := rwu.ReadUint64(reader)
		state, _ := rwu.ReadInt32(reader)
		actedBy, _ := rwu.ReadUint64(reader)
		rwu.ReadByte(reader) // 0
		stateChange := steamlang.EChatMemberStateChange(state)
		if stateChange == steamlang.EChatMemberStateChange_Entered {
			_, chatPerm, clanPerm := readChatMember(reader)
			s.Chats.AddChatMember(chatId, socialcache.ChatMember{
				SteamId:         steamid.SteamId(actedOn),
				ChatPermissions: chatPerm,
				ClanPermissions: clanPerm,
			})
		} else if stateChange == steamlang.EChatMemberStateChange_Banned || stateChange == steamlang.EChatMemberStateChange_Kicked ||
			stateChange == steamlang.EChatMemberStateChange_Disconnected || stateChange == steamlang.EChatMemberStateChange_Left {
			s.Chats.RemoveChatMember(chatId, steamid.SteamId(actedOn))
		}
		stateInfo := StateChangeDetails{
			ChatterActedOn: steamid.SteamId(actedOn),
			StateChange:    steamlang.EChatMemberStateChange(stateChange),
			ChatterActedBy: steamid.SteamId(actedBy),
		}
		s.client.Emit(&ChatMemberInfoEvent{
			ChatRoomId:      steamid.SteamId(body.SteamIdChat),
			Type:            steamlang.EChatInfoType(body.Type),
			StateChangeInfo: stateInfo,
		})
	}
}

func readChatMember(r io.Reader) (steamid.SteamId, steamlang.EChatPermission, steamlang.EClanPermission) {
	rwu.ReadString(r) // MessageObject
	rwu.ReadByte(r)   // 7
	rwu.ReadString(r) // steamid
	id, _ := rwu.ReadUint64(r)
	rwu.ReadByte(r)   // 2
	rwu.ReadString(r) // Permissions
	chat, _ := rwu.ReadInt32(r)
	rwu.ReadByte(r)   // 2
	rwu.ReadString(r) // Details
	clan, _ := rwu.ReadInt32(r)
	return steamid.SteamId(id), steamlang.EChatPermission(chat), steamlang.EClanPermission(clan)
}

func (s *Social) handleChatActionResult(packet *protocol.Packet) {
	body := new(steamlang.MsgClientChatActionResult)
	packet.ReadClientMsg(body)
	s.client.Emit(&ChatActionResultEvent{
		ChatRoomId: steamid.SteamId(body.SteamIdChat),
		ChatterId:  steamid.SteamId(body.SteamIdUserActedOn),
		Action:     steamlang.EChatAction(body.ChatAction),
		Result:     steamlang.EChatActionResult(body.ActionResult),
	})
}

func (s *Social) handleChatInvite(packet *protocol.Packet) {
	body := new(protobuf.CMsgClientChatInvite)
	packet.ReadProtoMsg(body)
	s.client.Emit(&ChatInviteEvent{
		InvitedId:    steamid.SteamId(body.GetSteamIdInvited()),
		ChatRoomId:   steamid.SteamId(body.GetSteamIdChat()),
		PatronId:     steamid.SteamId(body.GetSteamIdPatron()),
		ChatRoomType: steamlang.EChatRoomType(body.GetChatroomType()),
		FriendChatId: steamid.SteamId(body.GetSteamIdFriendChat()),
		ChatRoomName: body.GetChatName(),
		GameId:       body.GetGameId(),
	})
}

func (s *Social) handleIgnoreFriendResponse(packet *protocol.Packet) {
	body := new(steamlang.MsgClientSetIgnoreFriendResponse)
	packet.ReadClientMsg(body)
	s.client.Emit(&IgnoreFriendEvent{
		Result: steamlang.EResult(body.Result),
	})
}

func (s *Social) handleProfileInfoResponse(packet *protocol.Packet) {
	body := new(protobuf.CMsgClientFriendProfileInfoResponse)
	packet.ReadProtoMsg(body)
	s.client.Emit(&ProfileInfoEvent{
		Result:      steamlang.EResult(body.GetEresult()),
		SteamId:     steamid.SteamId(body.GetSteamidFriend()),
		TimeCreated: body.GetTimeCreated(),
		RealName:    body.GetRealName(),
		CityName:    body.GetCityName(),
		StateName:   body.GetStateName(),
		CountryName: body.GetCountryName(),
		Headline:    body.GetHeadline(),
		Summary:     body.GetSummary(),
	})
}

func (s *Social) handleFriendMessageHistoryResponse(packet *protocol.Packet) {
	body := new(protobuf.CMsgClientChatGetFriendMessageHistoryResponse)
	packet.ReadProtoMsg(body)
	steamid := steamid.SteamId(body.GetSteamid())
	for _, message := range body.GetMessages() {
		if !message.GetUnread() {
			continue // Skip already read messages
		}
		s.client.Emit(&ChatMsgEvent{
			ChatterId: steamid,
			Message:   message.GetMessage(),
			EntryType: steamlang.EChatEntryType_ChatMsg,
			Timestamp: time.Unix(int64(message.GetTimestamp()), 0),
			Offline:   true, // GetUnread is true
		})
	}
}
