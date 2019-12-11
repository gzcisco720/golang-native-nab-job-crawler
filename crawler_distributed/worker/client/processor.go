package client

import (
	"goweb/crawler/engine"
	"goweb/crawler/model"
	"goweb/crawler_distributed/rpc-support"
	"goweb/crawler_distributed/worker/server"
)

func CreateWorker() (engine.Processor, error) {
	client, err := rpc_support.NewRpcClient(":8082")
	if err != nil {
		return nil, err
	}
	return func(request model.Request) (result model.ParseResult, e error) {
		sRequest := server.SerialiseRequest(request)
		var sResult server.ParseResult
		err := client.Call("RpcWorkerService.Process", sRequest, &sResult)
		if err != nil {
			return  model.ParseResult{}, nil
		}
		return server.DeserialiseResult(sResult), nil
	}, nil
}
