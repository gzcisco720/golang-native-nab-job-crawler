package view

import (
	model2 "goweb/crawler/frontend/model"
	"html/template"
	"io"
)

type RenderService struct {
	template *template.Template
}

func (r RenderService) InitService(filename string) RenderService{
	return RenderService{
		template: template.Must(template.ParseFiles(filename)),
	}
}

func (r RenderService) Render(w io.Writer, data model2.JobListPage) error {
	return r.template.Execute(w, data)
}

