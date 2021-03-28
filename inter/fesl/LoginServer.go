package fesl

import (
	"github.com/Synaxis/bfheroesFesl/inter/network"
	"github.com/Synaxis/bfheroesFesl/inter/network/codec"
	"github.com/sirupsen/logrus"
)

// NuLoginServer - NuLogin for gameServer.exe
func (fm *Fesl) NuLoginServer(event network.EvProcess) {
	active := event.Client.IsActive
	if !active {
		logrus.Println("C Left")
		return
	}

	if event.Client.HashState.Get("clientType") != "server" {
		return
	}

	var id, userID, servername, secretKey, username string
	err := fm.db.stmtGetServerBySecret.QueryRow(event.Process.Msg["password"]).Scan(&id,
		&userID, &servername, &secretKey, &username)

	if err != nil {
		logrus.Println("===NuLogin issue=")
		return
	}

	Redis := make(map[string]interface{})
	Redis["uID"] = userID
	Redis["sID"] = id
	Redis["username"] = username
	Redis["apikey"] = event.Process.Msg["encryptedInfo"]
	Redis["keyHash"] = event.Process.Msg["password"]
	event.Client.HashState.SetM(Redis)

	//Setup new key for our persona	
	lkey, _ := getlkey()
	lkeyRedis := fm.level.NewObject("lkeys", lkey)
	lkeyRedis.Set("id", id)
	lkeyRedis.Set("userID", userID)
	lkeyRedis.Set("name", username)

	if !active {
		logrus.Println("AFK")
		return
	}

	event.Client.HashState.Set("lkeys", event.Client.HashState.Get("lkeys")+";"+lkey)
	event.Client.Answer(&codec.Packet{
		Content: ansNuLogin{
			TXN:       "NuLogin",
			ProfileID: userID,
			UserID:    userID,
			NucleusID: username,
			Lkey:      lkey,
		},
		Send:    event.Process.HEX,
		Message: acct,
	})
}



//NuLoginPersonaServer The Login is based on the Name
//there's only 1 persona(hero) for the server, so it works like a password
func (fm *Fesl) NuLoginPersonaServer(event network.EvProcess) {
	active := event.Client.IsActive
	/////Checks///////
	if !active {
		logrus.Println("AFK")
		return
	}

	logrus.Println("==LoginPersonaServer==")

	if event.Client.HashState.Get("clientType") != "server" {
		logrus.Println("==PossibleExploit==")
		fm.Goodbye(event)
		return
	}


	var id, userID, servername string
	//err := fm.db.stmtGetServerByName.QueryRow(event.Process.Msg["name"]).Scan(id, userID, servername, secretKey, username)

	// Setup a new key for our persona	
	lkey, _ := getlkey()
	lkeyRedis := fm.level.NewObject("lkeys", lkey)
	lkeyRedis.Set("id", userID)
	lkeyRedis.Set("userID", userID)
	lkeyRedis.Set("name", servername)

	event.Client.HashState.Set("lkeys", event.Client.HashState.Get("lkeys")+";"+lkey)
	event.Client.Answer(&codec.Packet{
		Content: ansNuLogin{
			TXN:       "NuLoginPersona",
			ProfileID: id,
			UserID:    userID,
			Lkey:      lkey,
		},
		Send:    event.Process.HEX,
		Message: acct,
	})

	logrus.Println("=== Server  Login OK===")
}
