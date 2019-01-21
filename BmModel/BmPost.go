package BmModel

import (
	"errors"

	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/mgo.v2/bson"
)

type Post struct {
	ID  string        `json:"id"`
	Id_ bson.ObjectId `json:"-" bson:"_id"`

	Title    string `json:"title" bson:"title"`
	SubTitle string `json:"sub-title" bson:"sub-title"`
	Content  string `json:"content" bson:"content"`

	CommentsIDs []string   `json:"-" bson:"comment-ids"`
	Comments    []*Comment `json:"-" bson:""`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (u Post) GetID() string {
	return u.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (u *Post) SetID(id string) error {
	u.ID = id
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u Post) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "comment",
			Name: "comments",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u Post) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}
	for _, kID := range u.CommentsIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "comment",
			Name: "comments",
		})
	}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u Post) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}
	for key := range u.CommentsIDs {
		result = append(result, u.Comments[key])
	}

	return result
}

// SetToManyReferenceIDs sets the leafs reference IDs and satisfies the jsonapi.UnmarshalToManyRelations interface
func (u *Post) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "comments" {
		u.CommentsIDs = IDs
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

// AddToManyIDs adds some new leafs that a users loves so much
func (u *Post) AddToManyIDs(name string, IDs []string) error {
	if name == "comments" {
		u.CommentsIDs = append(u.CommentsIDs, IDs...)
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

// DeleteToManyIDs removes some leafs from a users because they made him very sick
func (u *Post) DeleteToManyIDs(name string, IDs []string) error {
	if name == "comments" {
		for _, ID := range IDs {
			for pos, oldID := range u.CommentsIDs {
				if ID == oldID {
					// match, this ID must be removed
					u.CommentsIDs = append(u.CommentsIDs[:pos], u.CommentsIDs[pos+1:]...)
				}
			}
		}
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

func (u *Post) GetConditionsBsonM(parameters map[string][]string) bson.M {
	return bson.M{}
}
