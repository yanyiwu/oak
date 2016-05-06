package main

import (
	"flag"
	"strconv"

	"github.com/golang/glog"
	"github.com/yanyiwu/settledb/server"
)

var port = flag.Int("port", 8080, "listen port")

func main() {
	flag.Parse()
	addr := ":" + strconv.Itoa(*port)
	server.Init()
	glog.Info("Server started, listen " + addr)
	server.Serve(addr)
}
