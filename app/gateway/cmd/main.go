package main

import (
	"github.com/whuyfk/godolist/app/gateway/routes"
	"github.com/whuyfk/godolist/app/gateway/rpc"
)

func main() {
	rpc.InitRpc()
	router := routes.InitRouter()
	err := router.Run()
	if err != nil {
		return
	}
}
