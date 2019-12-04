package helper

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
)

func CreateDocByByte(data []byte) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err!=nil {
		return nil, err
	}
	return doc, nil
}