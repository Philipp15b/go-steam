package internal

import (
	. "github.com/smithfox/go-steam/internal/steamlang"
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
