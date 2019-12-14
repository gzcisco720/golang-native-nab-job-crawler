package persist

import (
	"goweb/crawler/model"
	"testing"
)

func TestElasticService_Save(t *testing.T) {
	elasticService := ElasticService{}
	elasticService.Init()
	item := model.JobProfile{
		JobNo:        "123",
		Title:        "test",
		Location:     "home",
		WorkType:     "full-time",
		BusinessUnit: "self",
		Date:         "0000-00-00",
	}
	err := elasticService.Save(item)
	if err != nil {
		t.Error(err)
	}
	res, err := elasticService.Client.
		Delete().
		Index("nab-career").
		Type("joblist").
		Id("123").
		Do(elasticService.Context)
	if err != nil {
		// Handle error
		t.Error(err)
	}
	if res != nil {
		t.Log("Test passed")
	}
}
