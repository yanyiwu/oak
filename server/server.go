package server

import (
	"encoding/json"
	"net/http"

	"github.com/golang/glog"
	"github.com/yanyiwu/settledb/engine"
)

var (
	ng *engine.Engine
)

func Init() {
	ng = engine.NewEngine()
}

func Serve(addr string) error {
	if ng == nil {
		glog.Fatal()
	}
	http.HandleFunc("/status", StatusHandler)
	http.HandleFunc("/keys", KeysHandler)
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

func KeysHandler(w http.ResponseWriter, req *http.Request) {
	keys := ng.GetKeys()
	b, _ := json.Marshal(keys)
	w.Write(b)
}
