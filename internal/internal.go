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

//Extra consts

// the default details to request in most situations
const EClientPersonaStateFlag_DefaultInfoRequest = EClientPersonaStateFlag_PlayerName |
	EClientPersonaStateFlag_Presence | EClientPersonaStateFlag_SourceID |
	EClientPersonaStateFlag_GameExtraInfo

const DefaultAvatar = "fef49e7fa7e1997310d705b2a6158ff8dc1cdfeb"

//Helper functions

func ValidAvatar(avatar string) bool {
	return !(avatar == "0000000000000000000000000000000000000000" || len(avatar) != 40)
}

// Helpers for reading

func ReadByte2Bool(r io.Reader) (bool, error) {
	var c uint8
	err := binary.Read(r, binary.LittleEndian, &c)
	return c != 0, err
}

func ReadUint8(r io.Reader) (uint8, error) {
	var c uint8
	err := binary.Read(r, binary.LittleEndian, &c)
	return c, err
}

func ReadUint16(r io.Reader) (uint16, error) {
	var c uint16
	err := binary.Read(r, binary.LittleEndian, &c)
	return c, err
}

func ReadUint32(r io.Reader) (uint32, error) {
	var c uint32
	err := binary.Read(r, binary.LittleEndian, &c)
	return c, err
}

func ReadUint64(r io.Reader) (uint64, error) {
	var c uint64
	err := binary.Read(r, binary.LittleEndian, &c)
	return c, err
}

func ReadInt8(r io.Reader) (int8, error) {
	var c int8
	err := binary.Read(r, binary.LittleEndian, &c)
	return c, err
}

func ReadInt16(r io.Reader) (int16, error) {
	var c int16
	err := binary.Read(r, binary.LittleEndian, &c)
	return c, err
}

func ReadInt32(r io.Reader) (int32, error) {
	var c int32
	err := binary.Read(r, binary.LittleEndian, &c)
	return c, err
}

func ReadInt64(r io.Reader) (int64, error) {
	var c int64
	err := binary.Read(r, binary.LittleEndian, &c)
	return c, err
}

func ReadString(r io.Reader) (string, error) {
	c := make([]byte, 0)
	var err error
	for {
		var temp byte
		err = binary.Read(r, binary.LittleEndian, &temp)
		if temp == byte(0x0) || err != nil {
			break
		}
		c = append(c, temp)
	}
	return string(c), err
}

func ReadByte(r io.Reader) (byte, error) {
	var c byte
	err := binary.Read(r, binary.LittleEndian, &c)
	return c, err
}

func ReadBytes(r io.Reader, num int32) ([]byte, error) {
	c := make([]byte, num)
	err := binary.Read(r, binary.LittleEndian, &c)
	return c, err
}

func WriteBool2Byte(w io.Writer, b bool) error {
	var err error
	if b {
		_, err = w.Write([]byte{1})
	} else {
		_, err = w.Write([]byte{0})
	}
	return err
}
