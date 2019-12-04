package parser

import (
	"goweb/crawler/fetcher"
	"goweb/crawler/model"
	model2 "goweb/crawler_distributed/model"

	"github.com/PuerkitoBio/goquery"
)

//ParseJobList func
func ParseJobList(doc *goquery.Document) model.ParseResult {
	result := model.ParseResult{}
	jobLinks, err := GetJobLinks()
	if err!=nil {
		panic(err)
	}
	for _, link := range jobLinks{
		result.Requests = append(
			result.Requests,
			model.Request{
				URL:        link,
				Parser: 	model2.FuncParserFactory(ParseJob,"ParseJob"),
			})
	}
	return result
}

func GetJobLinks()([]string, error) {
	doc, err := fetcher.Fetch("http://careers.nab.com.au/aw/en/listing/?page=1&page-items=9999")
	if err != nil {
		return nil, err
	}
	linksWithoutDuplicate := make(map[string]int, 0)
	doc.Find(".job-link").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr("href")
		_, exist := linksWithoutDuplicate[link]
		if ok && !exist {
			linksWithoutDuplicate[link] = i
		}
	})
	var jobList = make([]string, 0)

	for key := range linksWithoutDuplicate{
		jobList = append(jobList, "http://careers.nab.com.au"+key)
	}

	return jobList, nil
}