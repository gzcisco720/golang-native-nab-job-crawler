package engine

import (
	"goweb/fetcher"
	"goweb/model"
	"log"
)

type SimpleEngine struct {}
//Run func
func (e SimpleEngine) Run(seeds ...model.Request) {
	var requests []model.Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := e.worker(r)
		if err!=nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
		for _, res := range parseResult.Requests {
			log.Printf("Got item %v", res.URL)
		}
	}
}

func (e SimpleEngine) worker(r model.Request) (model.ParseResult, error){
	doc, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf("Error happens when fetch: %s", r.URL)
		return model.ParseResult{}, err
	}
	return r.ParserFunc(doc), nil
}