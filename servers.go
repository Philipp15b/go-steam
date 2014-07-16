package steam

import (
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

var CMServers = []string{
	// Qwest, Seattle
	"72.165.61.174:27017",
	"72.165.61.174:27018",
	"72.165.61.175:27017",
	"72.165.61.175:27018",
	"72.165.61.176:27017",
	"72.165.61.176:27018",
	"72.165.61.185:27017",
	"72.165.61.185:27018",
	"72.165.61.187:27017",
	"72.165.61.187:27018",
	"72.165.61.188:27017",
	"72.165.61.188:27018",
	// Inteliquent, Luxembourg, cm-[01-04].lux.valve.net
	"146.66.152.12:27017",
	"146.66.152.12:27018",
	"146.66.152.12:27019",
	"146.66.152.13:27017",
	"146.66.152.13:27018",
	"146.66.152.13:27019",
	"146.66.152.14:27017",
	"146.66.152.14:27018",
	"146.66.152.14:27019",
	"146.66.152.15:27017",
	"146.66.152.15:27018",
	"146.66.152.15:27019",
	/* Highwinds, Netherlands (not live)
	"81.171.115.5":27017",
	"81.171.115.5":27018",
	"81.171.115.5":27019",
	"81.171.115.6":27017",
	"81.171.115.6":27018",
	"81.171.115.6":27019",
	"81.171.115.7":27017",
	"81.171.115.7":27018",
	"81.171.115.7":27019",
	"81.171.115.8":27017",
	"81.171.115.8":27018",
	"81.171.115.8":27019",*/
	// Highwinds, Kaysville
	"209.197.29.196:27017",
	"209.197.29.197:27017",
	/* Starhub, Singapore (non-optimal route)
	"103.28.54.10":27017",
	"103.28.54.11":27017}*/
}

// An addr that is neither restricted to TCP nor UDP, but has an IP and a port.
type PortAddr struct {
	IP   net.IP
	Port uint16
}

// Parses an IP address with a port, for example "209.197.29.196:27017".
// If the given string is not valid, this function returns nil.
func ParsePortAddr(addr string) *PortAddr {
	parts := strings.Split(addr, ":")
	if len(parts) != 2 {
		return nil
	}
	ip := net.ParseIP(parts[0])
	if ip == nil {
		return nil
	}
	port, err := strconv.ParseUint(parts[1], 10, 16)
	if err != nil {
		return nil
	}
	return &PortAddr{ip, uint16(port)}
}

func (p *PortAddr) ToTCPAddr() *net.TCPAddr {
	return &net.TCPAddr{p.IP, int(p.Port), ""}
}

func (p *PortAddr) ToUDPAddr() *net.UDPAddr {
	return &net.UDPAddr{p.IP, int(p.Port), ""}
}

func (p *PortAddr) String() string {
	return p.IP.String() + ":" + strconv.FormatUint(uint64(p.Port), 10)
}

func GetRandomCM() *PortAddr {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	addr := ParsePortAddr(CMServers[rng.Int31n(int32(len(CMServers)))])
	if addr == nil {
		panic("invalid address in CMServers slice")
	}
	return addr
}
