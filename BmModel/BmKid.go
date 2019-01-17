package BmModel

import "gopkg.in/mgo.v2/bson"

// Kid is the Kid that a user consumes in order to get fat and happy
type Kid struct {
	ID           string        `json:"-"`
	Id_          bson.ObjectId `json:"-" bson:"_id"`
	Name         string        `json:"name" bson:"name"`
	NickName     string        `json:"nickname" bson:"nickname"`
	Gender       float64       `json:"gender" bson:"gender"`
	Dob          float64       `json:"dob" bson:"dob"`
	GuardianRole string        `json:"guardian-role" bson:"guardian-role"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (c Kid) GetID() string {
	return c.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (c *Kid) SetID(id string) error {
	c.ID = id
	return nil
}
