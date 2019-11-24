package main

import (
	"goweb/engine"
	"goweb/model"
	"goweb/nab-career/parser"
	"goweb/scheduler"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 10,
	}
	//e := engine.SimpleEngine{}
	e.Run(
		model.Request{
			URL: "http://careers.nab.com.au/aw/en/listing/?page=1&page-items=9999",
			ParserFunc: parser.ParseJobList,
		})
}
