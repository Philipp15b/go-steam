package appticket

import (
	"encoding/binary"
	"net"
	"time"
)

type AppTicket struct {
	Version                   uint32
	SteamId                   uint64
	AppId                     uint32
	OwnershipTicketExternalIP net.IP
	OwnershipTicketInternalIP net.IP
	OwnershipFlags            uint32
	OwnershipTicketGenerated  time.Time
	OwnershipTicketExpires    time.Time

	originalBuffer []byte
}

func (at *AppTicket) IsExpired(currentDate time.Time) bool {
	return at.OwnershipTicketExpires.Before(currentDate)
}

func (at *AppTicket) OriginalBuffer() []byte {
	return at.originalBuffer
}

func NewAppTicket(buffer []byte) (*AppTicket, error) {
	at := &AppTicket{
		originalBuffer: make([]byte, len(buffer)),
	}
	copy(at.originalBuffer, buffer)

	initialLength := binary.LittleEndian.Uint32(buffer[0:])

	if initialLength == 20 {
		panic("Unipmlemented")
	}

	at.Version = binary.LittleEndian.Uint32(buffer[4:])
	at.SteamId = binary.LittleEndian.Uint64(buffer[8:])
	at.AppId = binary.LittleEndian.Uint32(buffer[16:])
	at.OwnershipTicketExternalIP = bytesToIp4LE(buffer[20:])
	at.OwnershipTicketInternalIP = bytesToIp4LE(buffer[24:])
	at.OwnershipFlags = binary.LittleEndian.Uint32(buffer[28:])
	at.OwnershipTicketGenerated = time.Unix(int64(binary.LittleEndian.Uint32(buffer[32:])), 0)
	at.OwnershipTicketExpires = time.Unix(int64(binary.LittleEndian.Uint32(buffer[36:])), 0)

	return at, nil
}

func bytesToIp4LE(b []byte) net.IP {
	return net.IPv4(b[3], b[2], b[1], b[0])
}
