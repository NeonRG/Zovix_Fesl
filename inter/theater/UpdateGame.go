package theater

import (
	"github.com/Synaxis/bfheroesFesl/inter/network"
	"github.com/sirupsen/logrus"
	"strings"
)

type reqUGAM struct {
	TID int `fesl:"TID"`
	LobbyID int `fesl:"LID"`
	GameID int `fesl:"GID"`
	JoinMode string `fesl:"JOIN"`
	MaxPlayers int `fesl:"MAX-PLAYERS"`
	MaxObservers int `fesl:"B-maxObservers"`
	NumObservers int `fesl:"B-numObservers"`
}

func (tM *Theater) UGAM(event network.EvProcess) {
	
	logrus.Println("===UPDATE GAME==")
	gameID := event.Process.Msg["GID"] // TODO gameID := mm.FindGids()

	gdata := tM.level.NewObject("gdata", gameID)

	logrus.Println("Updating GameServer " + gameID)

	var args []interface{}
	keys := 0
	for index, value := range event.Process.Msg {
		if index == "TID" {
			continue
		}

		keys++

		// Strip quotes
		value = strings.Trim(value, `"`)

		gdata.Set(index, value)
		args = append(args, gameID)
		args = append(args, index)
		args = append(args, value)
	}
	_, err := tM.db.stmtUpdateGame.Exec(gameID)
	if err != nil {
		logrus.Println("======UGAM  Error==== ", err)
	}

	_, err = tM.db.setServerStatsStatement(keys).Exec(args...)
	if err != nil {
		logrus.Println("Failed to update stats for game server "+gameID, err.Error())
		if err.Error() == "Error 1213: Deadlock found when trying to get lock; try restarting transaction" {
			_, err = tM.db.setServerStatsStatement(keys).Exec(args...)
			if err != nil {
				logrus.Println("Failed to update stats for game server "+gameID+" on the second try", err.Error())
			}
		}
	}
}
