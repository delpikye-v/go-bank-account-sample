package service

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/go-kit/kit/bankaccount/pkg/model"
)

var accountStore model.AccountStore

func CreateSession() {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: []string{ENV_CONFIG.Mgourl},
	})
	// Timeout: 60 * time.Second,

	if err != nil {
		log.Fatalf("[MongoDB Session]: %s\n", err)
	}

	collection := session.DB(ENV_CONFIG.Mgodb).C("account")
	// collection.RemoveAll(nil)
	accountStore = model.AccountStore{
		C: collection,
	}
}
