package server

import (
	"encoding/json"
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/yanyiwu/settledb/engine"
)

var indexdir = flag.String("indexdir", "./indexdir", "")

var (
	ng *engine.Engine
)

func Init() {
	var err error
	ng, err = engine.NewEngine(*indexdir)
	if err != nil {
		glog.Error(err)
		panic(err)
	}
}

func Serve(addr string) error {
	if ng == nil {
		glog.Fatal()
	}
	http.HandleFunc("/status", StatusHandler)
	http.HandleFunc("/get", GetHandler)
	http.HandleFunc("/set", SetHandler)
	return http.ListenAndServe(addr, nil)
}

func StatusHandler(w http.ResponseWriter, req *http.Request) {
	//status := Status{
	//	Name: ServerName,
	//}
	req.ParseForm()
	b, _ := json.Marshal(req)
	w.Write(b)
}

func GetHandler(w http.ResponseWriter, req *http.Request) {
}

func SetHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		ResponseError(w, "http method should be POST")
		return
	}
	decoder := json.NewDecoder(req.Body)
	var t SetRequest
	err := decoder.Decode(&t)
	if err != nil {
		ResponseError(w, err.Error())
		return
	}
}
