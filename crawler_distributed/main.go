package main

import (
	"flag"
	"goweb/crawler/engine"
	coreModel "goweb/crawler/model"
	"goweb/crawler/nab-career/parser"
	"goweb/crawler/scheduler"
	itemsaver "goweb/crawler_distributed/itemsaver/client"
	distributedModel "goweb/crawler_distributed/model"
	rpc_support "goweb/crawler_distributed/rpc-support"
	worker "goweb/crawler_distributed/worker/client"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemSaverHost", "", "itemSaver host")
	workerHosts = flag.String("workerHosts", "", "worker hosts separated by comma")
)

func main() {
	flag.Parse()
	itemSaver, err := itemsaver.RpcItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}
	hosts := strings.Split(*workerHosts,",")
	clientChan := createClientPool(hosts)
	processor := worker.CreateWorker(clientChan)

	e := engine.ConcurrentEngine{
		Scheduler:   	&scheduler.QueuedScheduler{},
		WorkerCount: 	10,
		ItemChannel: 	itemSaver,
		Processor: 		processor,
	}
	e.Run(
		coreModel.Request{
			URL:        "http://careers.nab.com.au/aw/en/listing/?page=1&page-items=9999",
			Parser: 	distributedModel.FuncParserFactory(parser.ParseJobList, "ParseJobList"),
		})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpc_support.NewRpcClient(h)
		if err==nil {
			clients = append(clients, client)
			log.Printf("Connect to %s", h)
		} else {
			log.Printf("Error connecting to %s: %v", h, err)
		}
	}

	out := make(chan *rpc.Client)

	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()

	return out
}