package server

import (
	"goweb/crawler/engine"
)

type RpcWorkerService struct {}

func (r *RpcWorkerService)Process(req Request, result *ParseResult) error {
	dRequest, err := DeserialiseRequest(req)
	if err != nil {
		return err
	}
	wResult, err := engine.Worker(dRequest)
	if err != nil {
		return err
	}
	*result = SerialiseResult(wResult)
	return nil
}