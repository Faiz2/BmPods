package BmModel

import "gopkg.in/mgo.v2/bson"

type Teacher struct {
	ID  string        `json:"-"`
	Id_ bson.ObjectId `json:"-" bson:"_id"`

	Intro string `json:"intro" bson:"intro"`

	BrandId string `json:"brand-id" bson:"brand-id"`

	Name        string  `json:"name" bson:"name"`
	Nickname    string  `json:"nickname" bson:"nickname"`
	Icon        string  `json:"icon" bson:"icon"`
	Dob         float64 `json:"dob" bson:"dob"`
	Gender      float64 `json:"gender" bson:"gender"`
	RegDate     float64 `json:"reg-date" bson:"reg-date"`
	Contact     string  `json:"contact" bson:"contact"`
	WeChat      string  `json:"wechat" bson:"wechat"`
	JobTitle    string  `json:"job-title" bson:"job-title"`
	JobType     float64 `json:"job-type" bson:"job-type"` //0-兼职, 1-全职
	IdCardNo    string  `json:"idCardNo" bson:"idCardNo"`
	Major       string  `json:"major" bson:"major"`
	TeachYears  float64 `json:"teachYears" bson:"teachYears"`
	Province    string  `json:"province" bson:"province"`
	City        string  `json:"city" bson:"city"`
	District    string  `json:"district" bson:"district"`
	Address     string  `json:"address" bson:"address"`
	NativePlace string  `json:"nativePlace" bson:"nativePlace"`
	CreateTime  float64 `json:"create_time" bson:"create_time"`
}