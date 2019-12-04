package model

import "github.com/PuerkitoBio/goquery"

type ParserFunc func(*goquery.Document) ParseResult

type Parser interface {
	Parse(*goquery.Document) ParseResult
	Serialise() (name string, args interface{})
}