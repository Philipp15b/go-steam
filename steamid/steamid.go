package steamid

import (
	"fmt"
	"strconv"
)

type ChatInstanceFlag uint64

const (
	Clan     ChatInstanceFlag = 0x100000 >> 1
	Lobby    ChatInstanceFlag = 0x100000 >> 2
	MMSLobby ChatInstanceFlag = 0x100000 >> 3
)

type SteamId uint64

func New(accountId, instance uint32, universe int32, accountType int32) SteamId {
	s := SteamId(0)
	s = s.SetAccountId(accountId)
	s = s.SetAccountInstance(instance)
	s = s.SetAccountUniverse(universe)
	return s.SetAccountType(accountType)
}

func (s SteamId) ToUint64() uint64 {
	return uint64(s)
}

func (s SteamId) String() string {
	switch s.GetAccountType() {
	case 0: // EAccountType_Invalid
		fallthrough
	case 1: // EAccountType_Individual
		if s.GetAccountUniverse() <= 1 { // EUniverse_Public
			return fmt.Sprintf("STEAM_0:%d:%d", s.GetAccountId()&1, s.GetAccountId()>>1)
		} else {
			return fmt.Sprintf("STEAM_%d:%d:%d", s.GetAccountUniverse(), s.GetAccountId()&1, s.GetAccountId()>>1)
		}
	default:
		return strconv.FormatUint(uint64(s), 10)
	}
}

func (s SteamId) get(offset uint, mask uint64) uint64 {
	return (uint64(s) >> offset) & mask
}

func (s SteamId) set(offset uint, mask, value uint64) SteamId {
	return SteamId((uint64(s) & ^(mask << offset)) | (value&mask)<<offset)
}

func (s SteamId) GetAccountId() uint32 {
	return uint32(s.get(0, 0xFFFFFFFF))
}

func (s SteamId) SetAccountId(id uint32) SteamId {
	return s.set(0, 0xFFFFFFFF, uint64(id))
}

func (s SteamId) GetAccountInstance() uint32 {
	return uint32(s.get(32, 0xFFFFF))
}

func (s SteamId) SetAccountInstance(value uint32) SteamId {
	return s.set(32, 0xFFFFF, uint64(value))
}

func (s SteamId) GetAccountType() int32 {
	return int32(s.get(52, 0xF))
}

func (s SteamId) SetAccountType(t int32) SteamId {
	return s.set(52, 0xF, uint64(t))
}

func (s SteamId) GetAccountUniverse() int32 {
	return int32(s.get(56, 0xF))
}

func (s SteamId) SetAccountUniverse(universe int32) SteamId {
	return s.set(56, 0xF, uint64(universe))
}
