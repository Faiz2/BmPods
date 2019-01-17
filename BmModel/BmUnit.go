package BmModel

import (
	"gopkg.in/mgo.v2/bson"
	"errors"
	"github.com/manyminds/api2go/jsonapi"
)

type Unit struct {
	ID  string        `json:"-"`
	Id_ bson.ObjectId `json:"-" bson:"_id"`

	Status        float64                   `json:"status" bson:"status"`
	StartDate     float64                   `json:"start-date" bson:"start-date"`
	EndDate       float64                   `json:"end-date" bson:"end-date"`
	CourseTime    float64                   `json:"course-time" bson:"course-time"`		//课时

	//SessionableId string                    `json:"sessionableId" bson:"sessionableId"`
	//Sessionable   sessionable.BmSessionable `json:"Sessionable" jsonapi:"relationships"`

	TeacherID	string `json:'-' bson:"teacher-id"`
	Teacher       *Teacher `json:"-"`

	RoomID	string `json:"-" bson:"room-id"`
	Room 	*Room	`json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (u Unit) GetID() string {
	return u.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (u *Unit) SetID(id string) error {
	u.ID = id
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u Unit) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "Teacher",
			Name: "teacher",
		},
		{
			Type: "Room",
			Name: "room",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u Unit) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	if u.TeacherID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.TeacherID,
			Type: "Teacher",
			Name: "teacher",
		})
	}

	if u.RoomID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID: u.RoomID,
			Type: "Room",
			Name: "room",
		})
	}

	// TODO: sessionable

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u Unit) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	if u.TeacherID != "" {
		result = append(result, u.Teacher)
	}

	if u.RoomID  != "" {
		result = append(result, u.Room)
	}

	// TODO: sessionable

	return result
}

func (u *Unit) SetToOneReferenceID(name, ID string) error {
	if name == "teacher" {
		u.TeacherID = ID
		return nil
	}

	if name == "room" {
		u.RoomID = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}
