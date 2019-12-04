package persist

import (
	"context"
	"errors"
	"fmt"
	"github.com/olivere/elastic"
	"goweb/crawler/model"
)

type ElasticService struct {
	Client  *elastic.Client
	Context context.Context
}

func (e *ElasticService) Save(item model.JobProfile) error {
	if e.Client == nil {
		return errors.New("Please init elastic client first")
	}
	res, err := e.Client.
		Index().
		Index("nab-career").
		Type("joblist").
		Id(item.JobNo).
		BodyJson(item).
		Do(e.Context)
	if err != nil {
		return err
	}
	if res.Result == "created" {
		fmt.Println("New Job gets released: Id = ", res.Id)
	}
	return nil
}

func (e *ElasticService) Init() {
	e.Context = context.Background()
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	e.Client = client
	exists, err := client.IndexExists("nab-career").Do(e.Context)
	if !exists {
		_, err := client.CreateIndex("nab-career").Do(e.Context)
		if err != nil {
			panic(err)
		}
	}
}