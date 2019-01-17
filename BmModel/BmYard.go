package BmModel

import (
	"gopkg.in/mgo.v2/bson"
	"errors"
	"github.com/manyminds/api2go/jsonapi"
)

type Yard struct {
	ID          string        `json:"-"`
	Id_         bson.ObjectId `json:"-" bson:"_id"`

	BrandId     string        `json:"brand-id" bson:"brand-id"`
	Title       string        `json:"title" bson:"title"`
	Cover       string        `json:"cover" bson:"cover"`
	Description string        `json:"description" bson:"description"`
	Around      string        `json:"around" bson:"around"`

	//Address address.BmAddress `json:"address" bson:"relationships"`
	/**
	 * 在构建过程中，yard可能成为地址搜索的条件
	 */
	Province       string        `json:"province" bson:"province"`
	City           string        `json:"city" bson:"city"`
	District       string        `json:"district" bson:"district"`
	Address        string        `json:"address" bson:"address"`
	TrafficInfo    string        `json:"traffic_info" bson:"traffic_info"`
	Attribute      string        `json:"attribute" bson:"attribute"`
	Scenario       string        `json:"scenario" bson:"scenario"`
	OpenTime       string        `json:"openTime" bson:"openTime"`
	ServiceContact string        `json:"serviceContact" bson:"serviceContact"`
	Facilities     []string 	 `json:"facilities" bson:"facilities"`
	//Friendly       []string                   `json:"friendly" bson:"friendly"`

	//RoomCount float64 `json:"room_count"`
	/**
	 * 在构建过程中，除了排课逻辑，不会通过query到Room
	 */
	//TODO:Certifications合并成TagImgs,添加category做区分.
	ImagesIDs []string 						`json:"-"`
	Images	 []string `json:"-"`
	RoomsIDs	 []string						`json:"-"`
	Rooms 	 []string `json:"-"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (u Yard) GetID() string {
	return u.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (u *Yard) SetID(id string) error {
	u.ID = id
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (u Yard) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "Image",
			Name: "images",
		},
		{
			Type: "Room",
			Name: "rooms",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (u Yard) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}
	for _, kID := range u.ImagesIDs {
		result = append(result, jsonapi.ReferenceID{
			ID:   kID,
			Type: "image",
			Name: "images",
		})
	}

	//if u.CategoryID != "" {
	//	result = append(result, jsonapi.ReferenceID{
	//		ID: u.CategoryID,
	//		Type: "Category",
	//		Name: "category",
	//	})
	//}

	return result
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u Yard) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}
	//for key := range u.Images {
		//result = append(result, u.Images[key])
	//}

	//if u.CategoryID != "" {
		//result = append(result, u.Category)
	//}

	return result
}

//func (u *Yard) SetToOneReferenceID(name, ID string) error {
//	if name == "category" {
//		u.CategoryID = ID
//		return nil
//	}
//
//	return errors.New("There is no to-one relationship with the name " + name)
//}

// SetToManyReferenceIDs sets the leafs reference IDs and satisfies the jsonapi.UnmarshalToManyRelations interface
func (u *Yard) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "images" {
		u.ImagesIDs = IDs
		return nil
	}

	if name == "rooms" {
		u.RoomsIDs = IDs
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

// AddToManyIDs adds some new leafs that a users loves so much
func (u *Yard) AddToManyIDs(name string, IDs []string) error {
	if name == "images" {
		u.ImagesIDs = append(u.ImagesIDs, IDs...)
		return nil
	}

	if name == "rooms" {
		u.RoomsIDs = append(u.RoomsIDs, IDs...)
		return nil
	}

	return errors.New("There is no to-many relationship with the name " + name)
}

// DeleteToManyIDs removes some leafs from a users because they made him very sick
func (u *Yard) DeleteToManyIDs(name string, IDs []string) error {
	if name == "images" {
		for _, ID := range IDs {
			for pos, oldID := range u.ImagesIDs {
				if ID == oldID {
					// match, this ID must be removed
					u.ImagesIDs = append(u.ImagesIDs[:pos], u.ImagesIDs[pos+1:]...)
				}
			}
		}
	}

	if name == "rooms" {
		for _, ID := range IDs {
			for pos, oldID := range u.RoomsIDs {
				if ID == oldID {
					// match, this ID must be removed
					u.RoomsIDs = append(u.RoomsIDs[:pos], u.RoomsIDs[pos+1:]...)
				}
			}
		}
	}

	return errors.New("There is no to-many relationship with the name " + name)
}