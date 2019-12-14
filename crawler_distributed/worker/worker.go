package main

import (
	"flag"
	"fmt"
	rpc_support "goweb/crawler_distributed/rpc-support"
	"goweb/crawler_distributed/worker/server"
	"log"
)

var port = flag.Int("port", 0, "the port to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("Please specify the port")
		return
	}
	log.Fatal(rpc_support.ServeRpc(fmt.Sprintf(":%d", *port), &server.RpcWorkerService{}))
}