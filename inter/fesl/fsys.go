package fesl

import (
	"fmt"
	"time"

	"github.com/Synaxis/bfheroesFesl/config"
	"github.com/Synaxis/bfheroesFesl/inter/network"
	"github.com/Synaxis/bfheroesFesl/inter/network/codec"
	"github.com/sirupsen/logrus"
)

//THIS IS THE HELLO PACKET
const (
	fsys             = "fsys"
	fsysHello        = "Hello"
	fsysGetPingSites = "GetPingSites"
	fsysPing         = "Ping"
)

var firstLogin = true

// reqHello is definition of the fsys.Hello request call
type reqHello struct {
	TXN             string `fesl:"TXN"`
	clientString    string `fesl:"clientString"`
	Sku             int    `fesl:"sku"`
	Locale          string `fesl:"locale"`
	clientPlatform  string `fesl:"clientPlatform"`
	clientVersion   string `fesl:"clientVersion"`
	sdkVersion      string `fesl:"SDKVersion"`
	protocolVersion string `fesl:"protocolVersion"`
	fragmentSize    int    `fesl:"fragmentSize"`
	clientType      string `fesl:"clientType"`
}

type ansHello struct {
	TXN           string          `fesl:"TXN"`
	Domain        domainPartition `fesl:"domainPartition"`
	ConnTTL       int             `fesl:"activityTimeoutSecs"`
	ConnectedAt   string          `fesl:"curTime"`
	MessengerIP   string          `fesl:"messengerIp"`
	MessengerPort int             `fesl:"messengerPort"`
	TheaterIP     string          `fesl:"theaterIp"`
	TheaterPort   int             `fesl:"theaterPort"`
}

type domainPartition struct {
	Name    string `fesl:"domain"`
	SubName string `fesl:"subDomain"`
}

func (fm *Fesl) hello(event network.EvProcess) {
	Active := event.Client.IsActive
	if !Active {
		logrus.Println("Cli Left")
		return
	}	

	redisState := fm.createState(fmt.Sprintf(
		"%s-%s",
		event.Process.Msg["clientType"],
		event.Client.IpAddr.String(),
	))

	event.Client.HashState = redisState

	if !fm.server {
		fm.GetSessionId(event)
	}

	saveRedis := map[string]interface{}{
		"SDKVersion":     event.Process.Msg["SDKVersion"],
		"clientPlatform": event.Process.Msg["clientPlatform"],
		"clientString":   event.Process.Msg["clientString"],
		"clientType":     event.Process.Msg["clientType"],
		"clientVersion":  event.Process.Msg["clientVersion"],
		"locale":         event.Process.Msg["locale"],
		"sku":            event.Process.Msg["sku"],
	}
	event.Client.HashState.SetM(saveRedis)

	answer := ansHello{
		TXN:         fsysHello,
		ConnTTL:     int((60 * time.Hour).Seconds()),
		ConnectedAt: time.Now().Format("Jan-02-2006 15:04:05 MST"),
		TheaterIP:   config.General.ThtrAddr,
		MessengerIP: config.General.MessengerAddr,
	}

	if fm.server {
		answer.Domain = domainPartition{"eagames", "bfwest-server"}
		answer.TheaterPort = config.General.ThtrServerPort
	} else {
		answer.Domain = domainPartition{"eagames", "bfwest-dedicated"}
		answer.TheaterPort = config.General.ThtrClientPort
	}

	event.Client.Answer(&codec.Packet{
		Content: answer,
		Message: fsys,
		Send:    0xC0000001,
	})
	if !Active {
		fm.Goodbye(event)
		return
	}
}

type GetSessionId struct {
	TXN string `fesl:"TXN"`
}

func (fm *Fesl) GetSessionId(event network.EvProcess) {

	//Check if its a Server
	if event.Client.HashState.Get("clientType") == "server" {
		fm.NuLoginServer(event)
		return
	}

	event.Client.Answer(&codec.Packet{
		Content: GetSessionId{
			TXN: "GetSessionId", //case sensitive
		},
		Message: "gsum",
	})
}
type ansGoodbye struct {
	TXN        string `fesl:"TXN"`
	Reason     string `fesl:"reason"`
	messageArr string `fesl:"message"`
}

// Goodbye - Handle Client Close
func (fm *Fesl) Goodbye(event network.EvProcess) {
	logrus.Println("Client Disconnected")
	event.Client.Answer(&codec.Packet{
		Message: fsys,
		Send:    event.Process.HEX,
		Content: ansGoodbye{
			TXN:        "Goodbye",
			Reason:     "GOODBYE_CLIENT_NORMAL",
			messageArr: "n/a",
		},
	})
}

///////////////////////////////
type ansGetPingSites struct {
	TXN       string     `fesl:"TXN"`
	MinPings  int        `fesl:"minPingSitesToPing"`
	PingSites []pingSite `fesl:"pingSites"`
}

type pingSite struct {
	Addr    string `fesl:"addr"`
	Name    string `fesl:"name"`
	Message int    `fesl:"type"`
}

// GetPingSites - used as LoadBalancer/Not working Now(but is requested)
func (fm *Fesl) GetPingSites(event network.EvProcess) {

	event.Client.Answer(&codec.Packet{
		Message: fsys,
		Send:    event.Process.HEX,
		Content: ansGetPingSites{
			TXN:      fsysGetPingSites,
			MinPings: 0,
			PingSites: []pingSite{
				{"localhost", "iad", 1},
			},
		},
	})
}