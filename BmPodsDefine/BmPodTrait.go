package BmPodsDefine

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"github.com/alfredyang1986/BmPods/BmPanic"
	"github.com/alfredyang1986/BmPods/BmFactory"
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

	ins := BmFactory.GetInstanceByName(conf.Resource).(BmResource)
	fmt.Println(ins.GetResourceName())
}
