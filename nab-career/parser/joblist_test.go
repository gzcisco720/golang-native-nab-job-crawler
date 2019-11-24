package parser

import (
	"goweb/helper"
	"io/ioutil"
	"testing"
)

func TestParseJobList(t *testing.T) {
	file, err := ioutil.ReadFile("joblist_test.html")
	if err!=nil {
		t.Error(err)
	}

	doc, err := helper.CreateDocByByte(file)
	if err!=nil {
		t.Error(err)
	}

	result := ParseJobList(doc)

	if len(result.Items)!=384 {
		t.Error("Test has not met requirement")
	}
}