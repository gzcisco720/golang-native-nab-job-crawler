package client

import (
	"goweb/crawler/engine"
	"goweb/crawler/model"
	"goweb/crawler_distributed/worker/server"
	"net/rpc"
)

func CreateWorker(clientChan chan *rpc.Client) engine.Processor {
	return func(request model.Request) (result model.ParseResult, e error) {
		sRequest := server.SerialiseRequest(request)
		var sResult server.ParseResult
		c := <- clientChan
		err := c.Call("RpcWorkerService.Process", sRequest, &sResult)
		if err != nil {
			return  model.ParseResult{}, nil
		}
		return server.DeserialiseResult(sResult), nil
	}
}
