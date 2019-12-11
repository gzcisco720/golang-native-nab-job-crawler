package main

import (
	"goweb/crawler_distributed/itemsaver/server"
	"goweb/crawler_distributed/rpc-support"
	"log"
)

func main() {
	log.Fatal(rpc_support.ServeRpc(":8081", &server.RpcItemService{}))
}
