package main

import (
	"fmt"

	"github.com/yanyiwu/settledb/server"
)

func main() {
	server := server.NewServer()
	server.Start()
	fmt.Println("Hello SettleDB")
}
