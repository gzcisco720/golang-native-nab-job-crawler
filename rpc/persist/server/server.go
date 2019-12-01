package main

import (
	"goweb/rpc/persist"
	"goweb/service"
)

func main() {
	service.ServeRpc(":8081", persist.RpcItemService{})
}
