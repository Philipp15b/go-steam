package internal

import (
	"encoding/binary"
	. "github.com/Philipp15b/go-steam/internal/steamlang"
	"io"
)

type JobId uint64

type Serializer interface {
	Serialize(w io.Writer) error
}

type Deserializer interface {
	Deserialize(r io.Reader) error
}

type Serializable interface {
	Serializer
	Deserializer
}

type MessageBody interface {
	Serializable
	GetEMsg() EMsg
}

// the default details to request in most situations
const EClientPersonaStateFlag_DefaultInfoRequest = EClientPersonaStateFlag_PlayerName |
	EClientPersonaStateFlag_Presence | EClientPersonaStateFlag_SourceID |
	EClientPersonaStateFlag_GameExtraInfo

const DefaultAvatar = "fef49e7fa7e1997310d705b2a6158ff8dc1cdfeb"

func ValidAvatar(avatar string) bool {
	return !(avatar == "0000000000000000000000000000000000000000" || len(avatar) != 40)
}
