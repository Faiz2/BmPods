package BmModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

// ReservableItem is a generic database ReservableItem
type ReservableItem struct {
	ID  string        `json:"-"`
	Id_ bson.ObjectId `json:"-" bson:"_id"`

	BrandId    string  `json:"brandId" bson:"brandId"`
	Status     float64 `json:"status" bson:"status"` //0活动 1体验课 2普通课程
	StartDate  float64 `json:"start-date" bson:"start-date"`
	EndDate    float64 `json:"end-date" bson:"end-date"`
	CreateTime float64 `json:"create-time" bson:"create-time"`

	SessioninfoID string      `json:"sessioninfo-id" bson:"sessioninfo-id"`
	Sessioninfo   Sessioninfo `json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (u ReservableItem) GetID() string {
	return u.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (u *ReservableItem) SetID(id string) error {
	u.ID = id
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u ReservableItem) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "Sessioninfo",
			Name: "sessioninfo",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u ReservableItem) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	if u.SessioninfoID != "" {
		result = append(result, jsonapi.ReferenceID{
			ID:   u.SessioninfoID,
			Type: "sessioninfo",
			Name: "sessioninfo",
		})
	}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u ReservableItem) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	if u.SessioninfoID != "" {
		result = append(result, u.Sessioninfo)
	}

	return result
}

func (u *ReservableItem) SetToOneReferenceID(name, ID string) error {
	if name == "sessioninfo" {
		u.SessioninfoID = ID
		return nil
	}

	return errors.New("There is no to-one relationship with the name " + name)
}
