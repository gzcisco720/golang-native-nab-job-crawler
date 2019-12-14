package engine

import (
	"goweb/crawler/fetcher"
	"goweb/crawler/model"
	"log"
)

func Worker(r model.Request) (model.ParseResult, error){
	doc, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf("Error happens when fetch: %s", r.URL)
		return model.ParseResult{}, err
	}
	return r.Parser.Parse(doc), nil
}
