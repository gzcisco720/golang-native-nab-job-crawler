package model

import (
	"github.com/PuerkitoBio/goquery"
	"goweb/crawler/model"
)

type SerialisedParser struct {
	Name string
	Args interface{}
}

// Example for how to create a Parser
type NilParser struct {}

func (n NilParser) Parse(_ *goquery.Document) model.ParseResult {
	return model.ParseResult{}
}

func (n NilParser) Serialise() (name string, args interface{}) {
	return "NilParser", nil
}

type FuncParser struct {
	name string
	parser model.ParserFunc
}

func (f *FuncParser) Parse(doc *goquery.Document) model.ParseResult {
	return f.parser(doc)
}

func (f *FuncParser) Serialise() (name string, args interface{}) {
	return f.name, nil
}

func FuncParserFactory(f model.ParserFunc, name string) *FuncParser {
	return &FuncParser{
		name:   name,
		parser: f,
	}
}