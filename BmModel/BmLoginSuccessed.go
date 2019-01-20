package BmModel

import (
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2/bson"
)

type BmLoginSucceed struct {
	ID  string `json:"id"`
	Id_ bson.ObjectId

	// ApplicantID string     `json:"-" bson:"applicant-id"`
	// Applicant   *Applicant `json:"-"`
	Applicant Applicant `json:"Applicant" jsonapi:"relationships"`
	Token     string    `json:"token"`
}

// // GetID to satisfy jsonapi.MarshalIdentifier interface
// func (u BmLoginSucceed) GetID() string {
// 	return u.ID
// }

// // SetID to satisfy jsonapi.UnmarshalIdentifier interface
// func (u *BmLoginSucceed) SetID(id string) error {
// 	u.ID = id
// 	return nil
// }

// GetReferences to satisfy the jsonapi.MarshalReferences interface
// func (u BmLoginSucceed) GetReferences() []jsonapi.Reference {
// 	return []jsonapi.Reference{
// 		{
// 			Type: "applicant",
// 			Name: "applicant",
// 		},
// 	}
// }

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
// func (u BmLoginSucceed) GetReferencedIDs() []jsonapi.ReferenceID {
// 	result := []jsonapi.ReferenceID{}

// 	if u.ApplicantID != "" {
// 		result = append(result, jsonapi.ReferenceID{
// 			ID:   u.ApplicantID,
// 			Type: "applicant",
// 			Name: "applicant",
// 		})
// 	}

// 	return result
// }

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
// func (u BmLoginSucceed) GetReferencedStructs() []jsonapi.MarshalIdentifier {
// 	result := []jsonapi.MarshalIdentifier{}

// 	if u.ApplicantID != "" {
// 		result = append(result, u.Applicant)
// 	}

// 	return result
// }

// func (u *BmLoginSucceed) SetToOneReferenceID(name, ID string) error {
// 	if name == "applicant" {
// 		u.ApplicantID = ID
// 		return nil
// 	}

// 	return errors.New("There is no to-one relationship with the name " + name)
// }

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *BmLoginSucceed) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *BmLoginSucceed) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *BmLoginSucceed) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *BmLoginSucceed) QueryId() string {
	return bd.ID
}

func (bd *BmLoginSucceed) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *BmLoginSucceed) SetId(id string) {
	bd.ID = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd BmLoginSucceed) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "Applicant":
		bd.Applicant = v.(Applicant)
	}
	return bd
}

func (bd BmLoginSucceed) QueryConnect(tag string) interface{} {
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *BmLoginSucceed) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *BmLoginSucceed) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *BmLoginSucceed) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}
