package model

import (
	model2 "goweb/crawler/model"
)

type PagerItem struct {
	URL string
	PageNumber int
}

type JobListPage struct {
	Hits int64
	Start int
	Jobs [] model2.JobProfile
	Pagers []PagerItem
}
