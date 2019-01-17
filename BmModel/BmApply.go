package BmModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

// Apply is a generic database Apply
type Apply struct {
	ID  string        `json:"-"`
	Id_ bson.ObjectId `json:"-" bson:"_id"`

	Status       float64  `json:"status" bson:"status"` //0=未处理，1=已处理
	ApplyTime    float64  `json:"apply_time" bson:"apply_time"`
	ExceptTime   float64  `json:"except_time" bson:"except_time"`
	CreateTime   float64  `json:"create_time" bson:"create_time"`
	ApplyeeId    string   `json:"applyeeId" bson:"applyeeId"`
	BrandId      string   `json:"brandId" bson:"brandId"`
	ApplyFrom    string   `json:"applyFrom" bson:"applyFrom"`
	CourseType   float64  `json:"courseType" bson:"courseType"` //0活动 1体验课 2普通课程 -1预注册
	CourseName   string   `json:"courseName" bson:"courseName"`
	Contact      string   `json:"contact" bson:"contact"`
	ReservableId string   `json:"reservableId" bson:"reservableId"`
	Kids         []*Kid   `json:"-"`
	KidsIDs      []string `json:"-" bson:"kidsIds"`
}

//var ApplyName = strings.ToLower("Kid") + "s"
var ApplyName = "apply" + "s"

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (u Apply) GetID() string {
	return u.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (u *Apply) SetID(id string) error {
	u.ID = id
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u Apply) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: ApplyName,
			Name: ApplyName,
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u Apply) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}
	for _, chocolateID := range u.KidsIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   chocolateID,
			Type: ApplyName,
			Name: ApplyName,
		})
	}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u Apply) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}
	for key := range u.Kids {
		result = append(result, u.Kids[key])
	}

	return result
}

// SetToManyReferenceIDs sets the leafs reference IDs and satisfies the jsonapi.UnmarshalToManyRelations interface
func (u *Apply) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == ApplyName {
		u.KidsIDs = IDs
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

// AddToManyIDs adds some new leafs that a users loves so much
func (u *Apply) AddToManyIDs(name string, IDs []string) error {
	if name == ApplyName {
		u.KidsIDs = append(u.KidsIDs, IDs...)
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

// DeleteToManyIDs removes some leafs from a users because they made him very sick
func (u *Apply) DeleteToManyIDs(name string, IDs []string) error {
	if name == ApplyName {
		for _, ID := range IDs {
			for pos, oldID := range u.KidsIDs {
				if ID == oldID {
					// match, this ID must be removed
					u.KidsIDs = append(u.KidsIDs[:pos], u.KidsIDs[pos+1:]...)
				}
			}
		}
	}

	return errors.New("There is no to-many relationship with the name " + name)
}
