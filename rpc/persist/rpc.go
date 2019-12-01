package persist

import (
	"goweb/model"
	"goweb/service"
)

type RpcItemService struct {
	
}

func (r *RpcItemService)Save(item model.JobProfile, result *string) error {
	elasticService := service.ElasticService{}
	elasticService.Init()

	err := elasticService.Save(item)
	if err == nil {
		*result = "ok"
	}
	return err
}