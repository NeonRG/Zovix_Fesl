package fesl

import (
	"github.com/NeonRG/Zovix_Fesl/inter/network"
	"github.com/NeonRG/Zovix_Fesl/inter/network/codec"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)


type ansNuLogin struct {
	TXN       string `fesl:"TXN"`
	ProfileID string `fesl:"profileId"`
	UserID    string `fesl:"userId"`
	NucleusID string `fesl:"nuid"`
	Encrypt   int    `fesl:"returnEncryptedInfo"`
	Lkey      string `fesl:"lkey"`
}

//return  new uuid
func getlkey() (rand string, retErr error) {
	newRandom, _ := uuid.NewV4()
	newlkey := newRandom.String()

	return newlkey, nil
}

// NuLogin - First Login Command
func (fm *Fesl) NuLogin(event network.EvProcess) {

	if event.Client.HashState.Get("clientType") == "server" {
		// Server login
		fm.NuLoginServer(event)
		return
	}

	var id, username, email, birthday, language, country, gameToken string

	err := fm.db.stmtGetHeroByToken.QueryRow(event.Process.Msg["encryptedInfo"]).Scan(&id, &username,
		&email, &birthday, &language, &country, &gameToken) 

	if err != nil {
		logrus.Println("=nuLogin issue=")
		return
	}

	saveRedis := map[string]interface{}{
		"uID":       id,
		"username":  username,
		"sessionId": gameToken,
		"email":     email,
	}
	event.Client.HashState.SetM(saveRedis)

	// Setup a new key for our persona
	lkey, _ := getlkey()
	lkeyRedis := fm.level.NewObject("lkeys", lkey)
	lkeyRedis.Set("id", id)
	lkeyRedis.Set("userID", id)
	lkeyRedis.Set("name", username)

	event.Client.HashState.Set("lkeys", event.Client.HashState.Get("lkeys")+";"+lkey)
	event.Client.Answer(&codec.Packet{
		Content: ansNuLogin{
			TXN:       "NuLogin",
			ProfileID: id,
			UserID:    id,
			NucleusID: username,
			Lkey:      lkey,
		},
		Send:    event.Process.HEX,
		Message: acct,
	})
}

type ansNuLoginPersona struct {
	TXN       string `fesl:"TXN"`
	ProfileID string `fesl:"profileId"`
	UserID    string `fesl:"userId"`
	Lkey      string `fesl:"lkey"`
	Encrypt   int    `fesl:"returnEncryptedInfo"`
}

// User Login with selected Hero (persona)
func (fm *Fesl) NuLoginPersona(event network.EvProcess) {
	if !event.Client.IsActive {
		logrus.Println("C Left")
		return
	}

	if event.Client.HashState.Get("clientType") == "server" {
		logrus.Println("Server Login")
		fm.NuLoginPersonaServer(event)
		return
	}

	var id, userID, heroName, online string
	err := fm.db.stmtGetHeroByName.QueryRow(event.Process.Msg["name"]).Scan(&id, &userID, &heroName, &online)
	if err != nil {
		logrus.Println("Wrong Login")
		return
	}

	// Setup a new key for our persona
	lkey, _ := getlkey()
	lkeyRedis := fm.level.NewObject("lkeys", lkey)
	lkeyRedis.Set("id", id)
	lkeyRedis.Set("userID", userID)
	lkeyRedis.Set("name", heroName)

	saveRedis := make(map[string]interface{})
	saveRedis["heroID"] = id
	event.Client.HashState.SetM(saveRedis)

	event.Client.HashState.Set("lkeys", event.Client.HashState.Get("lkeys")+";"+lkey)

	event.Client.Answer(&codec.Packet{
		Content: ansNuLogin{ // todo check why its not nuLoginPersona struct
			TXN:       "NuLoginPersona",
			ProfileID: userID, // todo use PID
			UserID:    userID,
			Encrypt:   1,
			Lkey:      lkey,
		},
		Send:    event.Process.HEX,
		Message: acct,
	})
}