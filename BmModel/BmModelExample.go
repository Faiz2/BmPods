package BmModel

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

// ModelExample is a generic database ModelExample
type ModelExample struct {
	ID  string        `json:"-"`
	Id_ bson.ObjectId `json:"-" bson:"_id"`

	Attr1         string              `json:"attr1" bson:"attr1"`
	NoStoreAttr   string              `json:"-"`
	ModelLeafs    []*ModelLeafExample `json:"-"`
	ModelLeafsIDs []string            `json:"-" bson:"modelLeafsIds"`
}

//var ModelLeafName = strings.ToLower("ModelLeafExample") + "s"
var ModelLeafName = "modelLeafExample" + "s"

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (u ModelExample) GetID() string {
	return u.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (u *ModelExample) SetID(id string) error {
	u.ID = id
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u ModelExample) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: ModelLeafName,
			Name: ModelLeafName,
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u ModelExample) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}
	for _, chocolateID := range u.ModelLeafsIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   chocolateID,
			Type: ModelLeafName,
			Name: ModelLeafName,
		})
	}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u ModelExample) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}
	for key := range u.ModelLeafs {
		result = append(result, u.ModelLeafs[key])
	}

	return result
}

// SetToManyReferenceIDs sets the leafs reference IDs and satisfies the jsonapi.UnmarshalToManyRelations interface
func (u *ModelExample) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == ModelLeafName {
		u.ModelLeafsIDs = IDs
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

// AddToManyIDs adds some new leafs that a users loves so much
func (u *ModelExample) AddToManyIDs(name string, IDs []string) error {
	if name == ModelLeafName {
		u.ModelLeafsIDs = append(u.ModelLeafsIDs, IDs...)
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

// DeleteToManyIDs removes some leafs from a users because they made him very sick
func (u *ModelExample) DeleteToManyIDs(name string, IDs []string) error {
	if name == ModelLeafName {
		for _, ID := range IDs {
			for pos, oldID := range u.ModelLeafsIDs {
				if ID == oldID {
					// match, this ID must be removed
					u.ModelLeafsIDs = append(u.ModelLeafsIDs[:pos], u.ModelLeafsIDs[pos+1:]...)
				}
			}
		}
	}

	return errors.New("There is no to-many relationship with the name " + name)
}
