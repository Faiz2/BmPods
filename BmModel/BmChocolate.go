package BmModel

import "gopkg.in/mgo.v2/bson"

// Chocolate is the chocolate that a user consumes in order to get fat and happy
type Chocolate struct {
	ID    string `json:"-"`
	Id_ bson.ObjectId `json:"-" bson:"_id"`
	Name  string `json:"name" bson:"name"`
	Taste string `json:"taste" bson:"taste"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Chocolate) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Chocolate) SetID(id string) error {
	c.ID = id
	return nil
}