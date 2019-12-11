package engine

import (
	"fmt"
	model2 "goweb/crawler/model"
	parser2 "goweb/crawler/nab-career/parser"
	persist2 "goweb/crawler/persist"
	"log"
)

type Processor func(model2.Request) (model2.ParseResult, error)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChannel chan interface{}
	Processor Processor
}

type Scheduler interface {
	Submit(model2.Request)
	ConfigWorkerChan(chan model2.Request)
	WorkerReady(chan model2.Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...model2.Request) {
	out := make(chan model2.ParseResult)
	e.Scheduler.Run()
	elasticService := persist2.ElasticService{}
	elasticService.Init()

	jobLinks, err := parser2.GetJobLinks()
	if err!=nil {
		panic(err)
	}

	numOfJobs := len(jobLinks)

	for i:=0; i<e.WorkerCount; i++ {
		e.createWorker(out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	jobTh := 0
	fmt.Println("There are ",numOfJobs," Jobs")
	fmt.Println("=============================")
	for {
		if jobTh == numOfJobs {
			fmt.Println("=============================")
			break
		}
		result := <- out
		for _, item := range result.Items {
			jobTh++
			fmt.Println("JobNo:", item.JobNo, "JobTitle:", item.Title)
			go func() {
				e.ItemChannel <- item
			}()
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func (e *ConcurrentEngine) createWorker(out chan model2.ParseResult, scheduler Scheduler)  {
	in := make(chan model2.Request)
	go func() {
		for {
			scheduler.WorkerReady(in)
			r := <-in
			result, err := e.Processor(r)
			if err != nil {
				log.Printf("Error happens when fetch: %s", r.URL)
				continue
			}
			out <- result
		}
	}()
}