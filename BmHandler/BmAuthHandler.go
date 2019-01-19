package BmHandler

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"gopkg.in/mgo.v2/bson"

	"github.com/alfredyang1986/BmPods/BmModel"

	"github.com/julienschmidt/httprouter"

	"github.com/alfredyang1986/BmPods/BmDaemons"
	"github.com/alfredyang1986/BmPods/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/blackmirror/jsonapi/jsonapiobj"
)

type AuthHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
}

func (h AuthHandler) NewAuthHandler(args ...interface{}) AuthHandler {
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
				if tp.Name() == "BmMongodb" { //TODO: 这个地方有问题 BmMongodbDaemon
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
		} else {
		}
	}
	return AuthHandler{Method: md, HttpMethod: hm, Args: ag, db: m}
}

func (h AuthHandler) Validation(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	w.Header().Add("Content-Type", "application/json")
	var acc BmModel.Account
	json.NewDecoder(r.Body).Decode(&acc)

	out := BmModel.Account{}
	// TODO 实在是不敢乱动，又加了一个
	cond := bson.M{"account": acc.Account, "password": acc.Password}
	err := h.db.FindAccont(&acc, &out, cond)

	jso := jsonapiobj.JsResult{}
	response := map[string]interface{}{
		"status": "",
		"result": nil,
		"error":  nil,
	}
	if err == nil && out.Id_ != "" {
		hex := md5.New()
		out.Password = ""
		token := fmt.Sprintf("%x", hex.Sum(nil))
		out.Token = token

		response["status"] = "ok"
		response["result"] = out

		jso.Obj = response
		enc := json.NewEncoder(w)
		enc.Encode(jso.Obj)
		return 0
	} else {
		response["status"] = "error"
		response["error"] = "账户或密码错误！"
		jso.Obj = response
		enc := json.NewEncoder(w)
		enc.Encode(jso.Obj)
		return 1
	}

}

func (h AuthHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h AuthHandler) GetHandlerMethod() string {
	return h.Method
}
