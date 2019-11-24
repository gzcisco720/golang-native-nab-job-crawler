package engine

import (
	"goweb/fetcher"
	"goweb/model"
	"goweb/service"
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(model.Request)
	ConfigWorkerChan(chan model.Request)
	WorkerReady(chan model.Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...model.Request) {
	out := make(chan model.ParseResult)
	e.Scheduler.Run()

	elasticService := service.ElasticService{}
	elasticService.Init()

	for i:=0; i<e.WorkerCount; i++ {
		createWorker(out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	for {
		result := <- out
		for _, item := range result.Items {
			elasticService.Save(item)
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(out chan model.ParseResult, scheduler Scheduler)  {
	in := make(chan model.Request)
	go func() {
		for {
			scheduler.WorkerReady(in)
			r := <-in
			doc, err := fetcher.Fetch(r.URL)
			if err != nil {
				log.Printf("Error happens when fetch: %s", r.URL)
				continue
			}
			out <- r.ParserFunc(doc)
		}
	}()
}