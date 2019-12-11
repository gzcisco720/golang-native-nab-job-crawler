package main

import (
	"goweb/crawler/engine"
	coreModel "goweb/crawler/model"
	"goweb/crawler/nab-career/parser"
	"goweb/crawler/scheduler"
	itemsaver "goweb/crawler_distributed/itemsaver/client"
	distributedModel "goweb/crawler_distributed/model"
	worker "goweb/crawler_distributed/worker/client"
)

func main() {
	itemSaver, err := itemsaver.RpcItemSaver(":8081")
	if err != nil {
		panic(err)
	}
	processor, err := worker.CreateWorker()
	if err != nil {
		panic(err)
	}
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
