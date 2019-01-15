package BmFactory

import (
	"github.com/alfredyang1986/BmPods/BmResource"
	"github.com/alfredyang1986/BmPods/BmModel"
	"github.com/alfredyang1986/BmPods/BmDataStorage"
)

var BLACKMIRROR_MODEL_FACTORY = map[string]interface{}{
	"BmUser": BmModel.User{},
	"BmChocolate": BmModel.Chocolate{}}

var BLACKMIRROR_RESOURCE_FACTORY = map[string]interface{}{
	"BmUserResource": BmResource.BmUserResource{},
	"BmChocolateResource": BmResource.BmChocolateResource{} }

var BLACKMIRROR_STORAGE_FACTORY = map[string]interface{}{
	"BmUserStorage": BmDataStorage.UserStorage{},
	"BmChocolateStorage": BmDataStorage.ChocolateStorage{} }

func GetModelByName(name string) interface{}{
	return BLACKMIRROR_MODEL_FACTORY[name]
}

func GetResourceByName(name string) interface{}{
	return BLACKMIRROR_RESOURCE_FACTORY[name]
}

func GetStorageByName(name string) interface{}{
	return BLACKMIRROR_STORAGE_FACTORY[name]
}
