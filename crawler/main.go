package main

import (
	"goweb/crawler/engine"
	"goweb/crawler/model"
	"goweb/crawler/nab-career/parser"
	"goweb/crawler/persist"
	"goweb/crawler/scheduler"
	model2 "goweb/crawler_distributed/model"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChannel: persist.ItemSaver(),
	}
	//e := engine.SimpleEngine{}
	e.Run(
		model.Request{
			URL:        "http://careers.nab.com.au/aw/en/listing/?page=1&page-items=9999",
			Parser: 	model2.FuncParserFactory(parser.ParseJobList, "ParseJobList"),
		})
}
