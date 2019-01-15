package BmPodsDefine

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"github.com/alfredyang1986/BmPods/BmPanic"
	"github.com/manyminds/api2go"
	"github.com/alfredyang1986/BmPods/BmFactory"
	"github.com/alfredyang1986/BmPods/BmDataStorage"
	"github.com/alfredyang1986/BmPods/BmResource"
	"github.com/alfredyang1986/BmPods/BmSingleton"
	"github.com/manyminds/api2go/jsonapi"
)

type Pod struct {
	Name string
	Res map[string]interface{}
	conf Conf

	Storages map[string]BmDataStorage.BmStorage
	Resources map[string]BmResource.BmRes
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

	p.CreateStorageInstances()
	p.CreateResourceInstances()
}

func (p *Pod) CreateStorageInstances() {

	if p.Storages == nil {
		p.Storages = make(map[string]BmDataStorage.BmStorage)
	}

	for _, s := range p.conf.Storages {
		any := BmFactory.GetStorageByName(s.Name)
		name := s.Method
		inc, _ := BmSingleton.GetFactoryInstance().ReflectFunctionCall(any, name)
		p.Storages[s.Name] = inc.Interface()
	}

	fmt.Println(p.Storages)
}

func (p *Pod) CreateResourceInstances() {
	if p.Resources == nil {
		p.Resources = make(map[string]BmResource.BmRes)
	}

	for _, r := range p.conf.Resources {
		any := BmFactory.GetResourceByName(r.Name)
		name := r.Method
		var args []BmDataStorage.BmStorage
		for _, s := range r.Storages {
			tmp := p.Storages[s] //BmFactory.GetStorageByName(s)
			args = append(args, tmp)
		}

		inc, _ := BmSingleton.GetFactoryInstance().ReflectFunctionCall(any, name, args)
		p.Resources[r.Name] = inc.Interface()
	}

	fmt.Println(p.Resources)
}

func (p Pod) RegisterAllResource(api *api2go.API) {
	for _, ser := range p.conf.Services {
		md := BmFactory.GetModelByName(ser.Model)
		res := p.Resources[ser.Resource]
		api.AddResource(md.(jsonapi.MarshalIdentifier), res)
	}

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
