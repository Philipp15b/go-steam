package protocol

import (
	"encoding/binary"
	"io"
)

type MsgGCSetItemPosition struct {
	AssetId, Position uint64
}

func (m *MsgGCSetItemPosition) Serialize(w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, m)
}

type MsgGCCraft struct {
	Recipe   int16 // -2 = wildcard
	numItems int16
	Items    []uint64
}

func (m *MsgGCCraft) Serialize(w io.Writer) error {
	err := binary.Write(w, binary.LittleEndian, int16(-2))
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, int16(len(m.Items)))
	if err != nil {
		return err
	}
	for _, id := range m.Items {
		err = binary.Write(w, binary.LittleEndian, id)
		if err != nil {
			return err
		}
	}
	return nil
}

type MsgGCDeleteItem struct {
	ItemId uint64
}

func (m *MsgGCDeleteItem) Serialize(w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, m.ItemId)
}

type MsgGCNameItem struct {
	Tool, Target uint64
	Name         string
}

func (m *MsgGCNameItem) Serialize(w io.Writer) error {
	err := binary.Write(w, binary.LittleEndian, m.Tool)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, m.Target)
	if err != nil {
		return err
	}
	for _, char := range m.Name {
		err = binary.Write(w, binary.LittleEndian, int8(char))
		if err != nil {
			return err
		}
	}
	return err
}
