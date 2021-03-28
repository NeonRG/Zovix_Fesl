package theater

import (
	"github.com/Synaxis/bfheroesFesl/inter/network"
	"github.com/Synaxis/bfheroesFesl/inter/network/codec"

	"github.com/sirupsen/logrus"
)

type reqECNL struct {
	TID int `fesl:"TID"`
	GameID int `fesl:"GID"`
	LobbyID int `fesl:"LID"`
}

type ansECNL struct {
	TID     string `fesl:"TID"`
	GameID  string `fesl:"GID"`
	LobbyID string `fesl:"LID"`
}

// ECNL - EnterConnectionLost
func (tm *Theater) ECNL(event network.EvProcess) {
	logrus.Println("======ECNL======")
	logrus.Println("========SENT Leave Announcement to Player======")

	event.Client.Answer(&codec.Packet{
		Message: "ECNL",
		Content: ansECNL{
			event.Process.Msg["TID"],
			event.Process.Msg["GID"],
			event.Process.Msg["LID"],
		},
	})
}
