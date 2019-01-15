package BmPodsDefine

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"github.com/alfredyang1986/BmPods/BmPanic"
	"github.com/manyminds/api2go"
)

type Pod struct {
	Name string
	Res map[string]interface{}
	conf Conf
}

func (p *Pod) RegisterSerFromYAML(path string) {
	data, err := ioutil.ReadFile(path)
	if (err != nil) {
		fmt.Println("error")
	}
	//check(err)

	p.conf = Conf{}
	err = yaml.Unmarshal(data, &p.conf)
	if (err != nil) {
		fmt.Println("error")
		fmt.Println(err)
		panic(BmPanic.ALFRED_TEST_ERROR)
	}

	//ins := BmFactory.GetResourceByName(p.conf.Resource).(BmResource.BmRes)
	//fmt.Println(ins.GetResourceName())
}

func (p Pod) RegisterAllResource(api *api2go.API) {
	//res := BmFactory.GetResourceByName(p.conf.Resource).(BmResource.BmRes)
	//stor := BmFactory.GetStorageByName(p.conf.Storage).(BmDataStorage.BmStorage)
	//res.RegisterRelateStorage("self", stor)
	//for _, o2o := range p.conf.Relationships.One2one {
	//	storage := BmFactory.GetStorageByName(o2o["storage"]).(BmDataStorage.BmStorage)
	//	res.RegisterRelateStorage(o2o["storage"], storage)
	//}
	//
	//for _, o2m := range p.conf.Relationships.One2many {
	//	storage := BmFactory.GetStorageByName(o2m["storage"]).(BmDataStorage.BmStorage)
	//	res.RegisterRelateStorage(o2m["storage"], storage)
	//}
	//
	//fmt.Println(res)
	//api.AddResource(BmModel.User{}, res)

	//userStorage := BmMockDataStorage.NewUserStorage()
	//chocStorage := BmMockDataStorage.NewChocolateStorage()
	//api.AddResource(BmModel.User{}, BmResource.BmUserResource{ChocStorage: chocStorage, UserStorage: userStorage})
	//api.AddResource(BmModel.Chocolate{}, BmResource.BmChocolateResource{ChocStorage: chocStorage, UserStorage: userStorage})
}
