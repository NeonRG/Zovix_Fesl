package fesl

import (
	"github.com/sirupsen/logrus"
	"github.com/Synaxis/bfheroesFesl/inter/network"
	"github.com/Synaxis/bfheroesFesl/inter/network/codec"
)

type reqMemCheck struct {
	TXN string `fesl:"TXN"`
	Result string `fesl:"result"`
}

type ansMemCheck struct {	
	TXN string `fesl:"TXN"`
	MemChecks string `fesl:"memcheck.[]"`
	Salt      string `fesl:"salt"`
	Type      string `fesl:"type"`
	Result 	  string `fesl:"result"`
}


func (fm *Fesl) fsysMemCheck(event *network.EventNewClient) {
	logrus.Println("Sending MemCheck")
	event.Client.Answer(&codec.Packet{
		Message: fsys,
		Content: ansMemCheck{
			TXN:  "MemCheck",
			Salt: "5",
			Result: "",
			Type: "0",
			MemChecks: "0",
		},
		Send: 0xC0000000,
	})
}