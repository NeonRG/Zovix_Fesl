package fesl

import (
	"github.com/NeonRG/Zovix_Fesl/inter/network"
	"github.com/NeonRG/Zovix_Fesl/inter/mm"

	"github.com/NeonRG/Zovix_Fesl/inter/network/codec"
	"github.com/sirupsen/logrus"
)
const (
	pnow = "pnow"
)

type reqStart struct {
	TXN        string `fesl:"TXN"`
	debugLevel string `fesl:"debugLevel"`
	Version    int    `fesl:"version"`
	Partition 	string    `fesl:"partition"`

}

type ansStart struct {
	TXN    	string        `fesl:"TXN"`
	ID    	string        `fesl:"id.id"`//exclusive strings from the game
	Partition 	string    `fesl:"id.partition"`//exclusive strings from the game
}

// Start handles pnow.Start
func (fm *Fesl) Start(event network.EvProcess) {
	logrus.Println("---START---")

	event.Client.Answer(&codec.Packet{
		Content: ansStart{
			TXN:  "Start",
			ID:    "1",
			Partition: "eagames/bfwest-dedicated",
		},
		Send:    event.Process.HEX,
		Message: pnow,
	})
	fm.Status(event)

}

type ansStatus struct {
	Taxon        string                 `fesl:"TXN"`
	ID        string    `fesl:"id.id"`
	State string                 `fesl:"sessionState"`
	Properties   map[string]interface{} `fesl:"props"`
	Partition 	string				  `fesl:"partition"`
}

type stPartition struct {
	ID        string    `fesl:"id.id"`
	Partition string `fesl:"partition"`
}

type stGame struct {
	LobbyID int    `fesl:"lid"`
	Fit     int    `fesl:"fit"`
	GID  string `fesl:"gid"` //gameID to join
}
// Status comes after Start. tells info about desired server
func (fm *Fesl) Status(event network.EvProcess) {
	logrus.Println("--Status--")			


	for search := range mm.Games {  //is this crashing ?
		gid := search
		gamesArray := []stGame{
			{
				GID:     gid,
				Fit:     1001,
				LobbyID: 1,
			},
		}

	event.Client.Answer(&codec.Packet{
		Content: ansStatus{
			Taxon:        "Status",
			ID:           "1",
			State: "COMPLETE",
			Partition: "eagames/bfwest-dedicated",
			Properties: map[string]interface{}{
				"props.{}": "3", //the n of properties
				"resultType": "JOIN",
				"sessionType": "findServer",
				"games":      gamesArray},
		},		
		Send:    0x80000000,
		Message:    pnow,
	})

	}
}
