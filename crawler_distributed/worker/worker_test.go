package main

import (
	"goweb/crawler_distributed/model"
	rpc_support "goweb/crawler_distributed/rpc-support"
	"goweb/crawler_distributed/worker/server"
	"testing"
	"time"
)

func TestWorkerRpcService(t *testing.T) {
	const host = ":8082"

	go rpc_support.ServeRpc(host, &server.RpcWorkerService{})
	time.Sleep(time.Second)

	client, err := rpc_support.NewRpcClient(host)
	if err != nil {
		t.Error(err)
	}
	req := server.Request{
		URL:    "http://careers.nab.com.au/aw/en/listing/?page=1&page-items=9999",
		Parser: model.SerialisedParser{
			Name: "ParseJobList",
			Args: nil,
		},
	}
	var result server.ParseResult
	err = client.Call("RpcWorkerService.Process", req, &result)
	if err!=nil {
		t.Error(err)
	}
	t.Log(result)
}