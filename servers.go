package steam

import (
	"github.com/Philipp15b/go-steam/netutil"
	"math/rand"
	"time"
)

var CMServers = [][]string{
	{ //North American Servers
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
		// Highwinds, Kaysville
		"209.197.29.196:27017",
		"209.197.29.197:27017",
	},
	{ //Europe Servers
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
	},
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
	/* Starhub, Singapore (non-optimal route)
	"103.28.54.10":27017",
	"103.28.54.11":27017,*/
}

func GetRandomCM() *netutil.PortAddr {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	servers := append(CMServers[0], CMServers[1]...)
	addr := netutil.ParsePortAddr(servers[rng.Int31n(int32(len(servers)))])
	if addr == nil {
		panic("invalid address in CMServers slice")
	}
	return addr
}

func GetRandomNorthAmericaCM() *netutil.PortAddr {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	addr := netutil.ParsePortAddr(CMServers[0][rng.Int31n(int32(len(CMServers[0])))])
	if addr == nil {
		panic("invalid address in CMServers slice")
	}
	return addr
}

func GetRandomEuropeCM() *netutil.PortAddr {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	addr := netutil.ParsePortAddr(CMServers[1][rng.Int31n(int32(len(CMServers[0])))])
	if addr == nil {
		panic("invalid address in CMServers slice")
	}
	return addr
}
