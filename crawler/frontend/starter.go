package main

import (
	"fmt"
	controller2 "goweb/crawler/frontend/controller"
	view2 "goweb/crawler/frontend/view"
	"goweb/service"
	"net/http"
	"path/filepath"
)

func main() {
	elastic := service.ElasticService{}
	elastic.Init()
	templatePath, _ := filepath.Abs("frontend/view/joblist.html")
	filePath, _ := filepath.Abs("frontend/view")
	handler := controller2.JobListController{
		Renderer: view2.RenderService{}.InitService(templatePath),
		Client:   elastic.Client ,
	}

	http.Handle("/", http.FileServer(http.Dir(filePath)))

	http.Handle("/search", handler)
	fmt.Println("Server is listening on port: 8888")
	err := http.ListenAndServe(":8888", nil)
	if err!=nil {
		panic(err)
	}
}