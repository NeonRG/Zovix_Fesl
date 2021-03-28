package theater

import (
	"github.com/Synaxis/bfheroesFesl/inter/network"
	"github.com/Synaxis/bfheroesFesl/inter/network/codec"

	//"github.com/sirupsen/logrus"
)

type reqPENT struct {
	TID      int `fesl:"TID"`
	LobbyID  int `fesl:"LID"`
	GameID   int `fesl:"GID"`
	PlayerID int `fesl:"PID"`
}

type ansPENT struct {
	TID      string `fesl:"TID"`
	PlayerID string `fesl:"PID"`
}

// PENT - PlayerEntered
func (tM *Theater) PENT(event network.EvProcess) {

	// pid := event.Process.Msg["PID"]

	// // Get 4 stats for PID
	// rows, err := tM.db.getStatsStatement(4).Query(pid, "c_kit", "c_team", "elo", "level")
	// if err != nil {
	// 	logrus.Errorln("Failed gettings stats for hero "+pid, err.Error())
	// }

	// stats := make(map[string]string)

	// for rows.Next() {
	// 	var userID, heroID, heroName, statsKey, statsValue string
	// 	err := rows.Scan(&userID, &heroID, &heroName, &statsKey, &statsValue)
	// 	if err != nil {
	// 		logrus.Errorln("Issue with database:", err.Error())
	// 	}
	// 	stats[statsKey] = statsValue
	// }

	// switch stats["c_team"] {
	// case "1":
	// 	_, err = tM.db.stmtGameIncreaseTeam1.Exec(event.Process.Msg["GID"])
	// 	if err != nil {
	// 		logrus.Error("PENT ", err)
	// 	}
	// case "2":
	// 	_, err = tM.db.stmtGameIncreaseTeam2.Exec(event.Process.Msg["GID"])
	// 	if err != nil {
	// 		logrus.Error("PENT ", err)
	// 	}
	// default:
	// 	logrus.Errorln("Invalid team " + stats["c_team"] + " for " + pid)
	// }

	event.Client.Answer(&codec.Packet{
		Message: "PENT",
		Content: ansPENT{
			event.Process.Msg["TID"],
			event.Process.Msg["PID"],
		},
	})
}
