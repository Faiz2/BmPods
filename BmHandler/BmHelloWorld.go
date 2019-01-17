package BmHandler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"github.com/alfredyang1986/BmPods/BmDaemons/BmMongodb"
	"reflect"
	"github.com/alfredyang1986/BmPods/BmDaemons"
)

type HelloWorld struct {
	Method string
	HttpMethod string
	Args []string
	db *BmMongodb.BmMongodb
}

func (h HelloWorld) NewHelloWorld(args ...interface{}) HelloWorld {
	var m *BmMongodb.BmMongodb
	var hm string
	var md string
	var ag []string
	//sts := args[0].([]BmDaemons.BmDaemon)
	for i, arg := range args {
		if i == 0 {
			sts := arg.([]BmDaemons.BmDaemon)
			for _, dm := range sts {
				tp := reflect.ValueOf(dm).Elem().Type()
				if tp.Name() == "BmMongodbDaemon" {
					m = dm.(*BmMongodb.BmMongodb)
				}
			}
		} else if i == 1 {
			md = arg.(string)
		} else if i == 2 {
			hm = arg.(string)
		} else if i == 3 {
			lst := arg.([]string)
			for _, str := range lst {
				ag = append(ag, str)
			}
		} else {}
	}
	return HelloWorld{Method:md, HttpMethod:hm, Args:ag, db:m}
}

func (h HelloWorld) HelloWorldHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	fmt.Fprintf(w, "hello-world")
	return 0
}

func (h HelloWorld) GetHttpMethod() string {
	return h.HttpMethod
}

func (h HelloWorld) GetHandlerMethod() string {
	return h.Method
}
