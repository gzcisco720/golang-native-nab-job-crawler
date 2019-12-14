package server

import (
	"goweb/crawler/model"
	"goweb/crawler/persist"
)

type RpcItemService struct {}

func (r *RpcItemService)Save(item model.JobProfile, result *string) error {
	elasticService := persist.ElasticService{}
	elasticService.Init()
	err := elasticService.Save(item)
	if err == nil {
		*result = "ok"
	}
	return err
}