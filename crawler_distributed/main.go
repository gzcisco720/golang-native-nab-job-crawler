package main

import (
	"goweb/crawler/engine"
	"goweb/crawler/model"
	"goweb/crawler/nab-career/parser"
	"goweb/crawler/scheduler"
	"goweb/crawler_distributed/client"
	model2 "goweb/crawler_distributed/model"
)

func main() {
	itemSaver, err := client.RpcItemSaver(":8081")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChannel: itemSaver,
	}
	e.Run(
		model.Request{
			URL:        "http://careers.nab.com.au/aw/en/listing/?page=1&page-items=9999",
			Parser: 	model2.FuncParserFactory(parser.ParseJobList, "ParseJobList"),
		})
}
