package BmModel

import "gopkg.in/mgo.v2/bson"

// Room
type Room struct {
	ID  string        `json:"-"`
	Id_ bson.ObjectId `json:"-" bson:"_id"`

	YardId   string  `json:"yardId" bson:"yardId"`
	Title    string  `json:"title" bson:"title"`
	RoomType float64 `json:"roomType" bson:"roomType"`
	Capacity float64 `json:"capacity" bson:"capacity"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (r Room) GetID() string {
	return r.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (r *Room) SetID(id string) error {
	r.ID = id
	return nil
}
