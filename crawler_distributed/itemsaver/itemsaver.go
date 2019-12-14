package main

import (
	"flag"
	"fmt"
	"goweb/crawler_distributed/itemsaver/server"
	"goweb/crawler_distributed/rpc-support"
	"log"
)

var port = flag.Int("port", 0, "the port to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("Please specify the port")
		return
	}
	log.Fatal(rpc_support.ServeRpc(fmt.Sprintf(":%d", *port), &server.RpcItemService{}))
}
