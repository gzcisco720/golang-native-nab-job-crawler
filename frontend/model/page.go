package model

import (
	"goweb/model"
)

type PagerItem struct {
	URL string
	PageNumber int
}

type JobListPage struct {
	Hits int64
	Start int
	Jobs [] model.JobProfile
	Pagers []PagerItem
}
