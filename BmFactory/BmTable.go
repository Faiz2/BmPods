package BmFactory

import (
	"github.com/alfredyang1986/BmPods/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmPods/BmDaemons/BmRedis"
	"github.com/alfredyang1986/BmPods/BmDataStorage"
	"github.com/alfredyang1986/BmPods/BmHandler"
	"github.com/alfredyang1986/BmPods/BmModel"
	"github.com/alfredyang1986/BmPods/BmResource"
)

var BLACKMIRROR_MODEL_FACTORY = map[string]interface{}{
	"BmKid":            BmModel.Kid{},
	"BmApply":          BmModel.Apply{},
	"BmApplicant":      BmModel.Applicant{},
	"BmCategory":       BmModel.Category{},
	"BmImage":          BmModel.Image{},
	"BmSessioninfo":    BmModel.Sessioninfo{},
	"BmBrand":          BmModel.Brand{},
	"BmReservableitem": BmModel.Reservableitem{},
	"BmStudent":        BmModel.Student{},
	"BmGuardian":       BmModel.Guardian{},
	"BmTeacher":        BmModel.Teacher{},
	"BmRoom":           BmModel.Room{},
	"BmYard":           BmModel.Yard{},
	"BmUnit":           BmModel.Unit{},
	"BmClass":          BmModel.Class{},
	"BmPost":           BmModel.Post{},
	"BmComment":        BmModel.Comment{},
}

var BLACKMIRROR_RESOURCE_FACTORY = map[string]interface{}{
	"BmKidResource":            BmResource.BmKidResource{},
	"BmApplyResource":          BmResource.BmApplyResource{},
	"BmApplicantResource":      BmResource.BmApplicantResource{},
	"BmCategoryResource":       BmResource.BmCategoryResource{},
	"BmImageResource":          BmResource.BmImageResource{},
	"BmSessioninfoResource":    BmResource.BmSessioninfoResource{},
	"BmReservableitemResource": BmResource.BmReservableitemResource{},
	"BmBrandResource":          BmResource.BmBrandResource{},
	"BmStudentResource":        BmResource.BmStudentResource{},
	"BmGuardianResource":       BmResource.BmGuardianResource{},
	"BmTeacherResource":        BmResource.BmTeacherResource{},
	"BmRoomResource":           BmResource.BmRoomResource{},
	"BmYardResource":           BmResource.BmYardResource{},
	"BmUnitResource":           BmResource.BmUnitResource{},
	"BmClassResource":          BmResource.BmClassResource{},
	"BmPostResource":           BmResource.BmPostResource{},
	"BmCommentResource":        BmResource.BmCommentResource{},
}

var BLACKMIRROR_STORAGE_FACTORY = map[string]interface{}{
	"BmKidStorage":            BmDataStorage.BmKidStorage{},
	"BmApplyStorage":          BmDataStorage.BmApplyStorage{},
	"BmApplicantStorage":      BmDataStorage.BmApplicantStorage{},
	"BmCategoryStorage":       BmDataStorage.BmCategoryStorage{},
	"BmImageStorage":          BmDataStorage.BmImageStorage{},
	"BmSessioninfoStorage":    BmDataStorage.BmSessioninfoStorage{},
	"BmReservableitemStorage": BmDataStorage.BmReservableitemStorage{},
	"BmBrandStorage":          BmDataStorage.BmBrandStorage{},
	"BmStudentStorage":        BmDataStorage.BmStudentStorage{},
	"BmGuardianStorage":       BmDataStorage.BmGuardianStorage{},
	"BmTeacherStorage":        BmDataStorage.BmTeacherStorage{},
	"BmRoomStorage":           BmDataStorage.BmRoomStorage{},
	"BmYardStorage":           BmDataStorage.BmYardStorage{},
	"BmUnitStorage":           BmDataStorage.BmUnitStorage{},
	"BmClassStorage":          BmDataStorage.BmClassStorage{},
	"BmPostStorage":           BmDataStorage.BmPostStorage{},
	"BmCommentStorage":        BmDataStorage.BmCommentStorage{},
}

var BLACKMIRROR_DAEMON_FACTORY = map[string]interface{}{
	"BmMongodbDaemon": BmMongodb.BmMongodb{},
	"BmRedisDaemon":   BmRedis.BmRedis{},
}

var BLACKMIRROR_FUNCTION_FACTORY = map[string]interface{}{
	"BmProvinceHandler":    BmHandler.ProvinceHandler{},
	"BmCityHandler":        BmHandler.CityHandler{},
	"BmDistrictHandler":    BmHandler.DistrictHandler{},
	"BmUploadToOssHandler": BmHandler.UploadToOssHandler{},
	"BmAccountHandler":     BmHandler.AccountHandler{},
	"BmApplicantHandler":   BmHandler.ApplicantHandler{},
	"BmWeChatHandler":      BmHandler.WeChatHandler{},
}

func GetModelByName(name string) interface{} {
	return BLACKMIRROR_MODEL_FACTORY[name]
}

func GetResourceByName(name string) interface{} {
	return BLACKMIRROR_RESOURCE_FACTORY[name]
}

func GetStorageByName(name string) interface{} {
	return BLACKMIRROR_STORAGE_FACTORY[name]
}

func GetDaemonByName(name string) interface{} {
	return BLACKMIRROR_DAEMON_FACTORY[name]
}

func GetFunctionByName(name string) interface{} {
	return BLACKMIRROR_FUNCTION_FACTORY[name]
}
