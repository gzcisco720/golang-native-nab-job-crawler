package worker

import (
	"github.com/pkg/errors"
	model2 "goweb/crawler/model"
	"goweb/crawler/nab-career/parser"
	"goweb/crawler_distributed/model"
	"log"
)

type Request struct {
	URL string
	Parser model.SerialisedParser
}

type ParseResult struct {
	Requests []Request
	Items    []model2.JobProfile
}

func SerialiseRequest(r model2.Request) Request {
	name, args := r.Parser.Serialise()
	return Request{
		URL:    r.URL,
		Parser: model.SerialisedParser{
			Name: name,
			Args: args,
		},
	}	
}

func SerialiseResult(r model2.ParseResult) ParseResult {
	result := ParseResult{
		Items:    r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerialiseRequest(req))
	}
	return result
}

func DeserialiseRequest(r Request) (model2.Request,error) {
	p, err := deserialiseParser(r.Parser)
	if err != nil {
		return model2.Request{}, err
	}
	return model2.Request{
		URL:    r.URL,
		Parser: p,
	}, nil
}
func DeserialiseResult(r ParseResult) model2.ParseResult {
	result := model2.ParseResult{
		Items:    r.Items,
	}
	for _, req := range r.Requests {
		r, err := DeserialiseRequest(req)
		if err != nil {
			log.Println("Error happened when deserialise request")
			continue
		}
		result.Requests = append(result.Requests, r)
	}
	return result
}
func deserialiseParser(p model.SerialisedParser) (model2.Parser, error) {
	switch p.Name {
	case "ParseJobList":
		return model.FuncParserFactory(parser.ParseJobList, "ParseJobList"), nil
	case "ParseJob":
		return model.FuncParserFactory(parser.ParseJob, "ParseJob"), nil
	case "NilParser":
		return model.NilParser{}, nil
	default:
		return nil, errors.New("Cannot find Parser")
	}
}
