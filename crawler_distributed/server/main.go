package main

import (
	"goweb/crawler_distributed/persist"
	rpc_support "goweb/crawler_distributed/rpc-support"
	"log"
)

func main() {
	log.Fatal(rpc_support.ServeRpc(":8081", &persist.RpcItemService{}))
}
