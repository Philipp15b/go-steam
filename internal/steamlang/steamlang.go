/*
Contains code generated from SteamKit's SteamLanguage data.
*/
package steamlang

import (
	"encoding/binary"
	"io"
)

const (
	ProtoMask uint32 = 0x80000000
	EMsgMask         = ^ProtoMask
)

func NewEMsg(e uint32) EMsg {
	return EMsg(e & EMsgMask)
}

func IsProto(e uint32) bool {
	return e&ProtoMask > 0
}

func readByte2Bool(r io.Reader) (bool, error) {
	var c uint8
	err := binary.Read(r, binary.LittleEndian, &c)
	return c != 0, err
}

func readUint8(r io.Reader) (uint8, error) {
	var c uint8
	err := binary.Read(r, binary.LittleEndian, &c)
	return c, err
}

func readUint16(r io.Reader) (uint16, error) {
	var c uint16
	err := binary.Read(r, binary.LittleEndian, &c)
	return c, err
}

func readUint32(r io.Reader) (uint32, error) {
	var c uint32
	err := binary.Read(r, binary.LittleEndian, &c)
	return c, err
}

func readUint64(r io.Reader) (uint64, error) {
	var c uint64
	err := binary.Read(r, binary.LittleEndian, &c)
	return c, err
}

func readInt8(r io.Reader) (int8, error) {
	var c int8
	err := binary.Read(r, binary.LittleEndian, &c)
	return c, err
}

func readInt16(r io.Reader) (int16, error) {
	var c int16
	err := binary.Read(r, binary.LittleEndian, &c)
	return c, err
}

func readInt32(r io.Reader) (int32, error) {
	var c int32
	err := binary.Read(r, binary.LittleEndian, &c)
	return c, err
}

func readInt64(r io.Reader) (int64, error) {
	var c int64
	err := binary.Read(r, binary.LittleEndian, &c)
	return c, err
}

func writeBool2Byte(w io.Writer, b bool) error {
	var err error
	if b {
		_, err = w.Write([]byte{1})
	} else {
		_, err = w.Write([]byte{0})
	}
	return err
}
