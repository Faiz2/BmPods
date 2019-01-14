package main

import (
	"fmt"
	"github.com/alfredyang1986/BmPods/BmPodsDefine"
	"github.com/manyminds/api2go"
	"github.com/manyminds/api2go/examples/resolver"
	"github.com/manyminds/api2go/examples/storage"
	"github.com/manyminds/api2go/examples/model"
	"github.com/manyminds/api2go/examples/resource"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	fmt.Println("pod archi begins")
	var pod = BmPodsDefine.Pod{ Name: "alfred test" }
	pod.RegisterSerFromYAML("Resources/alfredtest.yaml")

	port := 31415
	api := api2go.NewAPIWithResolver("v0", &resolver.RequestURL{Port: port})
	userStorage := storage.NewUserStorage()
	chocStorage := storage.NewChocolateStorage()
	api.AddResource(model.User{}, resource.UserResource{ChocStorage: chocStorage, UserStorage: userStorage})
	api.AddResource(model.Chocolate{}, resource.ChocolateResource{ChocStorage: chocStorage, UserStorage: userStorage})

	fmt.Printf("Listening on :%d", port)
	handler := api.Handler().(*httprouter.Router)
	// It is also possible to get the instance of julienschmidt/httprouter and add more custom routes!
	handler.GET("/hello-world", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Hello World!\n")
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), handler)

	fmt.Println("pod archi ends")
}