package controller

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"goweb/frontend/model"
	"goweb/frontend/view"
	Types "goweb/model"
	"math"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type JobListController struct {
	Renderer view.RenderService
	Client *elastic.Client
}

func (j JobListController) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	if q=="" {
		q="*"
	}
	from, err := strconv.Atoi(req.FormValue("from"))
	if err!=nil {
		from = 0
	}

	page, err := j.GetJobList(q, from)
	if err!=nil {
		panic(err)
	}

	err = j.Renderer.Render(w, page)
	if err!=nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}

func (j JobListController) GetJobList(q string, from int) (model.JobListPage, error) {
	var result model.JobListPage
	pageSize := 50
	resp,err :=j.Client.Search("nab-career").
				Type("joblist").
				Query(elastic.NewQueryStringQuery(q)).
				From(from).
				Size(pageSize).
				Do(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		return result, err
	}
	result.Hits = resp.TotalHits()
	result.Start = from
	for _, v := range resp.Each(reflect.TypeOf(Types.JobProfile{})){
		item := v.(Types.JobProfile)
		result.Jobs = append(result.Jobs, item)
	}

	pageNumber := int(math.Ceil(float64(result.Hits)/float64(pageSize)))

	for i := 0; i < pageNumber; i++ {
		result.Pagers = append(result.Pagers, model.PagerItem{
			URL:        "/search?q=*&from="+strconv.Itoa(i*pageSize),
			PageNumber: i+1,
		})
	}

	return result,nil
}
