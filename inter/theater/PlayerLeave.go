package theater

import (
	"github.com/NeonRG/Zovix_Fesl/inter/network"
	"github.com/NeonRG/Zovix_Fesl/inter/network/codec"

	"github.com/sirupsen/logrus"
)

type reqPLVT struct {
	TID int `fesl:"TID"`
	LobbyID int `fesl:"LID"`
	GameID int `fesl:"GID"`
	PlayerID int `fesl:"PID"`
}

type ansKICK struct {
	TID     string `fesl:"TID"`
	LobbyID string `fesl:"LID"`
	GID     string `fesl:"GID"`
}

type ansPLVT struct {
	TID      string `fesl:"TID"`
	PlayerID string `fesl:"PID"`
}

// PLVT - PlayerLeaveTeam
func (tM *Theater) PLVT(event network.EvProcess) {

	pid := event.Process.Msg["PID"]
	// Get 4 stats for PID
	rows, err := tM.db.getStatsStatement(4).Query(pid, "c_kit", "c_team", "elo", "level")
	if err != nil {
		logrus.Errorln("Wrong stats for hero "+pid, err.Error())
	}

	stats := make(map[string]string)

	for rows.Next() {
		var userID, heroID, heroName, statsKey, statsValue string
		err := rows.Scan(&userID, &heroID, &heroName, &statsKey, &statsValue)
		if err != nil {
			logrus.Println("Issue with database:", err.Error())
		}
		stats[statsKey] = statsValue
	}

	switch stats["c_team"] {
	case "1":
		_, err = tM.db.stmtGameDecreaseTeam1.Exec(event.Process.Msg["GID"])
		if err != nil {
			logrus.Println("PLVT ", err)
		}
	case "2":
		_, err = tM.db.stmtGameDecreaseTeam2.Exec(event.Process.Msg["GID"])
		if err != nil {
			logrus.Println("PLVT ", err)
		}
	default:
		logrus.Println("Wrong team " + stats["c_team"] + " for " + pid)
	}

	event.Client.Answer(&codec.Packet{ // need to check this
		Message: "PLVT",
		Content: ansPLVT{
			event.Process.Msg["TID"],
			event.Process.Msg["PID"],
		},
	})

	event.Client.Answer(&codec.Packet{ // need to check this
		Message: "KICK",
		Content: ansKICK{
			event.Process.Msg["PID"],
			event.Process.Msg["LID"],
			event.Process.Msg["TID"],
		},
	})

}
