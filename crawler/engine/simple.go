package engine

import (
	model2 "goweb/crawler/model"
	"log"
)

type SimpleEngine struct {}
//Run func
func (e SimpleEngine) Run(seeds ...model2.Request) {
	var requests []model2.Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := Worker(r)
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