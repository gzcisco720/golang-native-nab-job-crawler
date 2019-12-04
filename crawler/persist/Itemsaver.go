package persist

import (
	"goweb/crawler/model"
)

func ItemSaver() chan interface{} {
	elasticService := ElasticService{}
	elasticService.Init()
	out := make(chan interface{})
	go func() {
		for {
			item := <-out
			if i, ok := item.(model.JobProfile); ok{
				elasticService.Save(i)
			}
		}
	}()
	return out
}
