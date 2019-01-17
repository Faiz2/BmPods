package BmFactory

import (
	"github.com/alfredyang1986/BmPods/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmPods/BmDataStorage"
	"github.com/alfredyang1986/BmPods/BmModel"
	"github.com/alfredyang1986/BmPods/BmResource"
)

var BLACKMIRROR_MODEL_FACTORY = map[string]interface{}{
	"BmKid":          BmModel.Kid{},
	"BmApply":        BmModel.Apply{},
	"BmApplicant":    BmModel.Applicant{},
	"BmCategory":     BmModel.Category{},
	"BmImage":        BmModel.Image{},
	"BmSessioninfo":  BmModel.Sessioninfo{},
	"BmReservableItem": BmModel.ReservableItem{},
	"BmBrand":        BmModel.Brand{},
}

var BLACKMIRROR_RESOURCE_FACTORY = map[string]interface{}{
	"BmKidResource":            BmResource.BmKidResource{},
	"BmApplyResource":          BmResource.BmApplyResource{},
	"BmApplicantResource":      BmResource.BmApplicantResource{},
	"BmCategoryResource":       BmResource.BmCategoryResource{},
	"BmImageResource":          BmResource.BmImageResource{},
	"BmSessioninfoResource":    BmResource.BmSessioninfoResource{},
	"BmReservableItemResource": BmResource.BmReservableItemResource{},
	"BmBrandResource":          BmResource.BmBrandResource{},
}

var BLACKMIRROR_STORAGE_FACTORY = map[string]interface{}{
	"BmKidStorage":            BmDataStorage.BmKidStorage{},
	"BmApplyStorage":          BmDataStorage.BmApplyStorage{},
	"BmApplicantStorage":      BmDataStorage.BmApplicantStorage{},
	"BmCategoryStorage":       BmDataStorage.BmCategoryStorage{},
	"BmImageStorage":          BmDataStorage.BmImageStorage{},
	"BmSessioninfoStorage":    BmDataStorage.BmSessioninfoStorage{},
	"BmReservableItemStorage": BmDataStorage.BmReservableItemStorage{},
	"BmBrandStorage":          BmDataStorage.BmBrandStorage{},
}

var BLACKMIRROR_DAEMON_FACTORY = map[string]interface{}{
	"BmMongodbDaemon": BmMongodb.BmMongodb{}}

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
