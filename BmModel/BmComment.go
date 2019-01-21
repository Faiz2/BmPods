package BmModel

import (
	"gopkg.in/mgo.v2/bson"
)

type Comment struct {
	ID  string        `json:"id"`
	Id_ bson.ObjectId `json:"-" bson:"_id"`

	AccountName string `json:"account-name" bson:"account-name"`
	Content     string `json:"content" bson:"content"`
	// PostID      string `json:"post-id" bson:"post-id"`
	// Post        *Post  `json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (u Comment) GetID() string {
	return u.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (u *Comment) SetID(id string) error {
	u.ID = id
	return nil
}

// // GetReferences to satisfy the jsonapi.MarshalReferences interface
// func (u Comment) GetReferences() []jsonapi.Reference {
// 	return []jsonapi.Reference{
// 		{
// 			Type: "post",
// 			Name: "post",
// 		},
// 	}
// }

// // GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
// func (u Comment) GetReferencedIDs() []jsonapi.ReferenceID {
// 	result := []jsonapi.ReferenceID{}
// 	if u.PostID != "" {
// 		result = append(result, jsonapi.ReferenceID{
// 			ID:   u.PostID,
// 			Type: "post",
// 			Name: "post",
// 		})
// 	}

// 	return result
// }

// // GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
// func (u Comment) GetReferencedStructs() []jsonapi.MarshalIdentifier {
// 	result := []jsonapi.MarshalIdentifier{}
// 	if u.PostID != "" {
// 		result = append(result, u.Post)
// 	}

// 	return result
// }

// func (u *Comment) SetToOneReferenceID(name, ID string) error {
// 	if name == "post" {
// 		u.PostID = ID
// 		return nil
// 	}

// 	return errors.New("There is no to-one relationship with the name " + name)
// }

func (u *Comment) GetConditionsBsonM(parameters map[string][]string) bson.M {
	return bson.M{}
}
