package main

import (
	rpc_support "goweb/crawler_distributed/rpc-support"
	"goweb/crawler_distributed/worker/server"
	"log"
)

func main() {
	log.Fatal(rpc_support.ServeRpc(":8082", &server.RpcWorkerService{}))
}