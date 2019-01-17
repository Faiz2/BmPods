package main

import (
	"fmt"
	"github.com/alfredyang1986/BmPods/BmApiResolver"
	"github.com/alfredyang1986/BmPods/BmPodsDefine"
	"github.com/julienschmidt/httprouter"
	"github.com/manyminds/api2go"
	"net/http"
)

func main() {
	fmt.Println("pod archi begins")
	var pod = BmPodsDefine.Pod{Name: "alfred test"}
	pod.RegisterSerFromYAML("Resources/alfredtest.yaml")

	port := 31415
	api := api2go.NewAPIWithResolver("v0", &BmApiResolver.RequestURL{Port: port})
	pod.RegisterAllResource(api)

	fmt.Printf("Listening on :%d", port)
	pod.RegisterAllFunctions(api)

	handler := api.Handler().(*httprouter.Router)
	http.ListenAndServe(fmt.Sprintf(":%d", port), handler)

	fmt.Println("pod archi ends")
}
