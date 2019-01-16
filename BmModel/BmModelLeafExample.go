package BmModel

import "gopkg.in/mgo.v2/bson"

// ModelLeafExample is the ModelLeafExample that a user consumes in order to get fat and happy
type ModelLeafExample struct {
	ID    string `json:"-"`
	Id_ bson.ObjectId `json:"-" bson:"_id"`
	Attr2  string `json:"attr2" bson:"attr2"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c ModelLeafExample) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *ModelLeafExample) SetID(id string) error {
	c.ID = id
	return nil
}