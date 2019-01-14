package BmFactory

import (
	"github.com/alfredyang1986/BmPods/BmResource"
)

var BLACKMIRROR_CONF_FACTORY = map[string]interface{}{
	"BmUser": BmResource.BmUserResource{} }

func GetInstanceByName(name string) interface{}{
	return BLACKMIRROR_CONF_FACTORY[name]
}