package BmPodsDefine

import (
	"fmt"
	"github.com/alfredyang1986/BmPods/BmDaemons"
	"github.com/alfredyang1986/BmPods/BmDataStorage"
	"github.com/alfredyang1986/BmPods/BmFactory"
	"github.com/alfredyang1986/BmPods/BmPanic"
	"github.com/alfredyang1986/BmPods/BmResource"
	"github.com/alfredyang1986/BmPods/BmSingleton"
	"github.com/manyminds/api2go"
	"github.com/manyminds/api2go/jsonapi"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Pod struct {
	Name string
	Res  map[string]interface{}
	conf Conf

	Storages  map[string]BmDataStorage.BmStorage
	Resources map[string]BmResource.BmRes
	Daemons   map[string]BmDaemons.BmDaemon
}

func (p *Pod) RegisterSerFromYAML(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("error")
	}
	//check(err)

	p.conf = Conf{}
	err = yaml.Unmarshal(data, &p.conf)
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		panic(BmPanic.ALFRED_TEST_ERROR)
	}

	p.CreateDaemonInstances()
	p.CreateStorageInstances()
	p.CreateResourceInstances()
}

func (p *Pod) CreateDaemonInstances() {
	if p.Daemons == nil {
		p.Daemons = make(map[string]BmDaemons.BmDaemon)
	}

	for _, d := range p.conf.Daemons {
		any := BmFactory.GetDaemonByName(d.Name)
		name := d.Method
		args := d.Args
		inc, _ := BmSingleton.GetFactoryInstance().ReflectFunctionCall(any, name, args)
		p.Daemons[d.Name] = inc.Interface()
	}
}

func (p *Pod) CreateStorageInstances() {

	if p.Storages == nil {
		p.Storages = make(map[string]BmDataStorage.BmStorage)
	}

	for _, s := range p.conf.Storages {
		any := BmFactory.GetStorageByName(s.Name)
		name := s.Method
		var args []BmDaemons.BmDaemon
		for _, d := range s.Daemons {
			tmp := p.Daemons[d]
			args = append(args, tmp)
		}

		inc, _ := BmSingleton.GetFactoryInstance().ReflectFunctionCall(any, name, args)
		p.Storages[s.Name] = inc.Interface()
	}
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
}

func (p Pod) RegisterAllResource(api *api2go.API) {
	for _, ser := range p.conf.Services {
		md := BmFactory.GetModelByName(ser.Model)
		res := p.Resources[ser.Resource]
		api.AddResource(md.(jsonapi.MarshalIdentifier), res)
	}
}
