package steam

import (
	. "github.com/Philipp15b/go-steam/internal/steamlang"
	. "github.com/Philipp15b/go-steam/steamid"
)

type LoggedOnEvent struct {
	Result                    EResult
	ExtendedResult            EResult
	OutOfGameSecsPerHeartbeat int32
	InGameSecsPerHeartbeat    int32
	PublicIp                  uint32
	ServerTime                uint32
	AccountFlags              EAccountFlags
	ClientSteamId             SteamId `json:",string"`
	EmailDomain               string
	CellId                    uint32
	CellIdPingThreshold       uint32
	Steam2Ticket              []byte
	UsePics                   bool
	WebApiUserNonce           string
	IpCountryCode             string
	VanityUrl                 string
	NumLoginFailuresToMigrate int32
	NumDisconnectsToMigrate   int32
}

type LoginKeyEvent struct {
	UniqueId uint32
	LoginKey string
}

type LoggedOffEvent struct {
	Result EResult
}

type MachineAuthUpdateEvent struct {
	Hash []byte
}

type AccountInfoEvent struct {
	PersonaName          string
	Country              string
	PasswordSalt         []byte
	PasswordSHADisgest   []byte
	CountAuthedComputers int32
	LockedWithIpt        bool
	AccountFlags         EAccountFlags
	FacebookId           uint64 `json:",string"`
	FacebookName         string
}
