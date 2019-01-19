package BmModel

import "gopkg.in/mgo.v2/bson"

type Account struct {
	ID  string        `json:"-"`
	Id_ bson.ObjectId `json:"-" bson:"_id"`

	Account string `json:"account" bson:"account"`
	// SecretWord string `json:"secretword" bson:"secretword"`
	Password string `json:"password" bson:"password"`

	Token string `json:"token"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (u Account) GetID() string {
	return u.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (u *Account) SetID(id string) error {
	u.ID = id
	return nil
}

func (u *Account) GetConditionsBsonM(parameters map[string][]string) bson.M {
	return bson.M{}
}
