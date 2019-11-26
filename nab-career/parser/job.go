package parser

import (
	"github.com/PuerkitoBio/goquery"
	"goweb/helper"
	"goweb/model"
	"strings"
	"time"
)


func ParseJob(doc *goquery.Document) model.ParseResult {
	var str string
	var jobTitle string
	if len(doc.Find(".jobcon").Nodes) !=0 {
		str = doc.Find(".jobcon p").First().Text()
		jobTitle = doc.Find(".jobcon h3").First().Text()
	} else if len(doc.Find("#job-info").Nodes) !=0 {
		str = doc.Find("#job-info").Text()
		jobTitle = doc.Find("#job-info").Parent().Find("h1 strong").Text()
	} else {
		str = doc.Find("#job-content p").First().Text()
		jobTitle = doc.Find("#job-content").Find("h2").Text()
	}

	date, hasDate := doc.Find(".open-date time").Attr("datetime")
	rows := strings.Split(str,"\n")

	rows = helper.StringSliceFilter(rows, func(str string) bool {
		formattedString := strings.Trim(str, "\t \n")
		return len(formattedString)!= 0 && formattedString != "Apply now"
	})

	var profile = model.JobProfile{}
	profile.Title = jobTitle
	for _, row := range rows {
		title := helper.CamelStyle(strings.Split(row,":")[0])
		switch title {
			case "JobNo":
				jobId := strings.TrimSpace(strings.Split(row,":")[1])
				if strings.Index(jobId,"-") >= 0 {
					jobId = strings.TrimSpace(strings.Split(row,"-")[1])
				}
				profile.JobNo = jobId
			case "Location":
				profile.Location = strings.Split(row,":")[1]
			case "WorkType":
				profile.WorkType = strings.Split(row,":")[1]
			case "BusinessUnit":
				profile.BusinessUnit = strings.Split(row,":")[1]
			default:
		}
		if hasDate {
			t, _ := time.Parse(
				time.RFC3339,
				date)
			profile.Date = t.Format("2006-8-2 15:04")
		}
	}

	return model.ParseResult{
		Items: []model.JobProfile{profile},
	}
}