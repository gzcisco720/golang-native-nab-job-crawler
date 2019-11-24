package parser

import (
	"github.com/PuerkitoBio/goquery"
	"goweb/helper"
	"goweb/model"
	"strings"
)


func ParseJob(doc *goquery.Document) model.ParseResult {
	var str string

	if len(doc.Find(".jobcon").Nodes) !=0 {
		str = doc.Find(".jobcon p").First().Text()
	} else if len(doc.Find("#job-info").Nodes) !=0 {
		str = doc.Find("#job-info").Text()
	} else {
		str = doc.Find("#job-content p").First().Text()
	}

	date, hasDate := doc.Find(".open-date time").Attr("datetime")
	rows := strings.Split(str,"\n")

	rows = helper.StringSliceFilter(rows, func(str string) bool {
		formattedString := strings.Trim(str, "\t \n")
		return len(formattedString)!= 0 && formattedString != "Apply now"
	})

	var profile = model.JobProfile{}
	for _, row := range rows {
		title := helper.CamelStyle(strings.Split(row,":")[0])
		switch title {
			case "JobNo":
				profile.JobNo = strings.TrimSpace(strings.Split(row,":")[1])
			case "Location":
				profile.Location = strings.Split(row,":")[1]
			case "WorkType":
				profile.WorkType = strings.Split(row,":")[1]
			case "BusinessUnit":
				profile.BusinessUnit = strings.Split(row,":")[1]
			default:
		}
		if hasDate {
			//t, _:= time.Parse(time.RFC3339, date)
			profile.Date = date
		}
	}

	return model.ParseResult{
		Items: []model.JobProfile{profile},
	}
}