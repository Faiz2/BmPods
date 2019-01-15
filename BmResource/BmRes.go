package BmResource

import (
	"github.com/alfredyang1986/BmPods/BmDataStorage"
)

type BmRes interface {
	GetResourceName() string
	RegisterRelateStorage(n string, s BmDataStorage.BmStorage)
}