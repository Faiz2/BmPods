package BmFactory

import (
	"github.com/alfredyang1986/BmPods/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmPods/BmDataStorage"
	"github.com/alfredyang1986/BmPods/BmModel"
	"github.com/alfredyang1986/BmPods/BmResource"
)

var BLACKMIRROR_MODEL_FACTORY = map[string]interface{}{
	"BmUser":           BmModel.User{},
	"BmChocolate":      BmModel.Chocolate{},
	"BmModelExample":     BmModel.ModelExample{},
	"BmModelLeafExample": BmModel.ModelLeafExample{},
	"Kid": BmModel.Kid{},
}

var BLACKMIRROR_RESOURCE_FACTORY = map[string]interface{}{
	"BmUserResource":             BmResource.BmUserResource{},
	"BmChocolateResource":        BmResource.BmChocolateResource{},
	"BmModelResourceExample":     BmResource.BmModelResourceExample{},
	"BmModelLeafResourceExample": BmResource.BmModelLeafResourceExample{},
	"BmKidResource": BmResource.BmKidResource{},
}

var BLACKMIRROR_STORAGE_FACTORY = map[string]interface{}{
	"BmUserStorage":             BmDataStorage.UserStorage{},
	"BmChocolateStorage":        BmDataStorage.ChocolateStorage{},
	"BmModelStorageExample":     BmDataStorage.BmModelStorageExample{},
	"BmModelLeafStorageExample": BmDataStorage.BmModelLeafStorageExample{},
	"BmKidStorage": BmDataStorage.BmKidStorage{},
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
