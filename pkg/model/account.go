package model

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

/**
JSON Account Collection
{
	uuid: '',
	type: '',
	currency: '',
	balance: '',
	number: '',
}
*/

type Account struct {
	Uuid     string  `json:"uuid"`
	Type     string  `json:"type"`
	Currency string  `json:"currency"`
	Balance  float64 `json:"balance"`
	Number   int64   `json:"number"`
}

type AccountStore struct {
	C *mgo.Collection
}

func (store AccountStore) CreateAccount(b *Account) error {
	// Assign a new bson.ObjectId
	// b.ID = bson.NewObjectId()
	err := store.C.Insert(b)
	return err
}

func (store AccountStore) GetAccount(uuid string) (Account, error) {
	var acc Account
	err := store.C.Find(bson.M{"uuid": uuid}).One(&acc)
	return acc, err
}

func (store AccountStore) GetAccountByNumber(number int64) (Account, error) {
	var acc Account
	err := store.C.Find(bson.M{"number": number}).One(&acc)
	return acc, err
}

func (store AccountStore) UpdateAccount(b *Account) error {
	error := store.C.Update(
		bson.M{"number": b.Number},
		bson.M{"$set": bson.M{
			"type":     b.Type,
			"currency": b.Currency,
			"balance":  b.Balance,
			"number":   b.Number,
		},
		},
	)
	return error
}
