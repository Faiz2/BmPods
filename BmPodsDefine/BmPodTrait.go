package BmPodsDefine

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"github.com/alfredyang1986/BmPods/BmPanic"
	"github.com/alfredyang1986/BmPods/BmFactory"
	"github.com/alfredyang1986/BmPods/BmModel"
	"github.com/manyminds/api2go"
	"github.com/alfredyang1986/BmPods/BmResource"
	"github.com/alfredyang1986/BmPods/BmMockDataStorage"
)

type Pod struct {
	Name string
	Res map[string]interface{}
}

func (p *Pod) RegisterSerFromYAML(path string) {
	data, err := ioutil.ReadFile(path)
	if (err != nil) {
		fmt.Println("error")
	}
	//check(err)

	conf := Conf{}
	err = yaml.Unmarshal(data, &conf)
	if (err != nil) {
		fmt.Println("error")
		fmt.Println(err)
		panic(BmPanic.ALFRED_TEST_ERROR)
	}

	ins := BmFactory.GetInstanceByName(conf.Resource).(BmRes)
	fmt.Println(ins.GetResourceName())
}

func (p Pod) RegisterAllResource(api *api2go.API) {
	userStorage := BmMockDataStorage.NewUserStorage()
	chocStorage := BmMockDataStorage.NewChocolateStorage()
	api.AddResource(BmModel.User{}, BmResource.BmUserResource{ChocStorage: chocStorage, UserStorage: userStorage})
	api.AddResource(BmModel.Chocolate{}, BmResource.BmChocolateResource{ChocStorage: chocStorage, UserStorage: userStorage})
}
