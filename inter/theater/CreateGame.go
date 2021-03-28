package theater

import (
	"fmt"
	"net"

	"github.com/NeonRG/Zovix_Fesl/inter/mm"
	"github.com/NeonRG/Zovix_Fesl/inter/network"
	"github.com/NeonRG/Zovix_Fesl/inter/network/codec"
	"github.com/sirupsen/logrus"
)

type reqCGAM struct {
	Tid int `fesl:"TID"`
	LobbyID int `fesl:"LID"`
	ReserveHost int `fesl:"RESERVE-HOST"`
	Name string `fesl:"NAME"`
	Port int `fesl:"PORT"`
	Httype string `fesl:"HTTYPE"`
	Type string `fesl:"TYPE"`
	Qlen int `fesl:"QLEN"`
	DisableAutoDequeue int `fesl:"DISABLE-AUTO-DEQUEUE"`
	Hxfr int `fesl:"HXFR"`
	IntPort int `fesl:"INT-PORT"`
	IntIP string `fesl:"INT-IP"`
	MaxPlayers int `fesl:"MAX-PLAYERS"`
	BMaxObservers int `fesl:"B-maxObservers"`
	BNumObservers int `fesl:"B-numObservers"`
	Ugid string `fesl:"UGID"` /// Value passed in +guid
	Secret string `fesl:"SECRET"` // Value passed in +secret
	BUAlwaysQueue int `fesl:"B-U-alwaysQueue"`
	BUArmyBalance string `fesl:"B-U-army_balance"`
	BUArmyDistribution string `fesl:"B-U-army_distribution"`
	BUAvailSlotsNational string `fesl:"B-U-avail_slots_national"`
	BUAvailSlotsRoyal string `fesl:"B-U-avail_slots_royal"`
	BUAvgAllyRank string `fesl:"B-U-avg_ally_rank"`
	BUAvgAxisRank string `fesl:"B-U-avg_axis_rank"`
	BUCommunityName string `fesl:"B-U-community_name"`
	BUDataCenter string `fesl:"B-U-data_center"`
	BUEloRank string `fesl:"B-U-elo_rank"`
	BUMap string `fesl:"B-U-map"`
	BUPercentFull int `fesl:"B-U-percent_full"`
	BUServerIP string `fesl:"B-U-server_ip"`
	BUServerPort int `fesl:"B-U-server_port"`
	BUServerState string `fesl:"B-U-server_state"`
	BVersion string `fesl:"B-version"`
	Join string `fesl:"JOIN"`
	Rt string `fesl:"RT"`
}
type ansCGAM struct {
	TID        string `fesl:"TID"`
	LobbyID    int `fesl:"LID"`
	MaxPlayers string `fesl:"MAX-PLAYERS"`
	EKEY       string `fesl:"EKEY"`
	UGID       string `fesl:"UGID"`
	Secret     string `fesl:"SECRET"`
	JOIN       string `fesl:"JOIN"`
	J          string `fesl:"J"`
	GameID     string `fesl:"GID"`
	isRanked   bool   `fesl:"B-U-UNRANKED"`
}

// CGAM - CreateGameParameters
func (tm *Theater) CGAM(event network.EvProcess) {

	answer := event.Process.Msg

	addr, ok := event.Client.IpAddr.(*net.TCPAddr)
	if !ok {
		logrus.Errorln("Failed turning IpAddr to net.TCPAddr")
		return
	}

	res, err := tm.db.stmtCreateServer.Exec(
		answer["NAME"],
		answer["B-U-community_name"],
		answer["INT-IP"],
		answer["INT-PORT"],
		answer["B-version"],
	)
	if err != nil {
		logrus.Error("Cannot create New server", err)
		return
	}

	id, _ := res.LastInsertId()
	gameID := fmt.Sprintf("%d", id)

	// Store gameID for access later
	mm.Games[gameID] = event.Client

	var args []interface{}

	// Setup a new key for our game
	gameServer := tm.level.NewObject("gdata", gameID)

	keys := 0

	// Stores what we know about this game in the redis db
	for index, value := range answer {
		if index == "TID" {
			continue
		}

		keys++

		// Strip quotes
		if len(value) > 0 && value[0] == '"' {
			value = value[1:]
		}
		if len(value) > 0 && value[len(value)-1] == '"' {
			value = value[:len(value)-1]
		}
		gameServer.Set(index, value)

		args = append(args, gameID)
		args = append(args, index)
		args = append(args, value)
	}

	gameServer.Set("LID", "1")
	gameServer.Set("GID", gameID)
	gameServer.Set("IP", addr.IP.String())
	gameServer.Set("AP", "0")
	gameServer.Set("QUEUE-LENGTH", "0")

	event.Client.HashState.Set("gdata:GID", gameID)

	_, err = tm.db.setServerStatsStatement(keys).Exec(args...)
	if err != nil {
		logrus.Error("Failed setting stats for game server "+gameID, err.Error())
		return
	}
	logrus.Println("===========CGAM=============")
	event.Client.Answer(&codec.Packet{
		Message: "CGAM",
		Content: ansCGAM{
			TID:        answer["TID"],
			LobbyID:    1, //should not be hardcoded
			UGID:       answer["UGID"],
			MaxPlayers: answer["MAX-PLAYERS"],
			EKEY:       `O65zZ2D2A58mNrZw1hmuJw%3d%3d`,
			Secret:     "MargeSimpson",
			isRanked:   false,
			JOIN:       answer["JOIN"],
			J:          answer["JOIN"],
			GameID:     gameID,
		},
	})

	// Create game in database
	_, err = tm.db.stmtAddGame.Exec(gameID, addr.IP.String(), answer["PORT"], answer["B-version"], answer["JOIN"], answer["B-U-map"], 0, 0, answer["MAX-PLAYERS"], 0, 0, "")
	if err != nil {
		logrus.Println("Failed to add game: %v", err)
	}
	logrus.Println("Added GAMESERVER TO DB")
}
