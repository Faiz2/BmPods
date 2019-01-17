package BmModel

import (
	"gopkg.in/mgo.v2/bson"
)

type Applicant struct {
	ID  string        `json:"-"`
	Id_ bson.ObjectId `json:"-" bson:"_id"`

	Name            string  `json:"name" bson:"name"`
	Gender          float64 `json:"gender" bson:"gender"`
	Pic             string  `json:"pic" bson:"pic"`
	RegisterPhone   string  `json:"regi-phone" bson:"regi-phone"`
	WeChatOpenid    string  `json:"wechat-openid" bson:"wechat-openid"`
	WeChatBindPhone string  `json:"wechat-bind-phone" bson:"wechat-bind-phone"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (u Applicant) GetID() string {
	return u.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (u *Applicant) SetID(id string) error {
	u.ID = id
	return nil
}
