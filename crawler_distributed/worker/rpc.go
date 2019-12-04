package worker

import "goweb/crawler/engine"

type WorkerService struct {}

func (WorkerService)Process(req Request, result *ParseResult) error {
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