package model

import "github.com/PuerkitoBio/goquery"

type Request struct {
	URL        string
	ParserFunc func(*goquery.Document) ParseResult
}