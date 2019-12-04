package parser

import (
	"goweb/crawler/helper"
	"io/ioutil"
	"testing"
)

type Profile map[string]string

func TestParseJob(t *testing.T) {
	file, err := ioutil.ReadFile("job_test.html")
	if err!=nil {
		t.Error(err)
	}
	doc, err := helper.CreateDocByByte(file)

	jobProfile := ParseJob(doc)

	t.Log(jobProfile)

	if len(jobProfile.Items) == 0 {
		t.Error("Test has not met requirement")
	}
}