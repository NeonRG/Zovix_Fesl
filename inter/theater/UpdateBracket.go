package theater

import (
	"github.com/NeonRG/Zovix_Fesl/inter/network"
	"github.com/NeonRG/Zovix_Fesl/inter/network/codec"

	"github.com/sirupsen/logrus"
)

type reqUBRA struct {
	TID int `fesl:"TID"`
	LobbyID int `fesl:"LID"`
	GameID int `fesl:"GID"`
	START int `fesl:"START"`
}

type ansUBRA struct {
	TID   string `fesl:"TID"`
	LID   string `fesl:"LID"`
	START int    `fesl:"START"`
}

// UBRA - "UpdateBracket" updates players connected (AP)
func (tM *Theater) UBRA(event network.EvProcess) {
	logrus.Println("===UBRA==")

	
	gdata := tM.level.NewObject("gdata", event.Process.Msg["GID"])

	if event.Process.Msg["START"] == "1" {
		gdata.Set("AP", "0") // AP = ActivePlayer (individual)
		// If Player Entered -> Reset AP
	}

	event.Client.Answer(&codec.Packet{
		Message: "UBRA",
		Content: ansUBRA{
			TID:   event.Process.Msg["TID"],
			LID:   event.Process.Msg["LID"],
			START: 1,
		}},
	)

}
