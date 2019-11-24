package parser

import (
	"goweb/model"

	"github.com/PuerkitoBio/goquery"
)

//ParseJobList func
func ParseJobList(doc *goquery.Document) model.ParseResult {
	result := model.ParseResult{}

	doc.Find(".job-link").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr("href")
		if ok {
			result.Requests = append(
				result.Requests,
				model.Request{
					URL:        "http://careers.nab.com.au" + link,
					ParserFunc: ParseJob,
				})
			//result.Items = append(result.Items, s.Text())
		}
	})
	return result
}
