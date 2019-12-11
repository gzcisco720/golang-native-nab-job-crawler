package main

import (
	"goweb/crawler/model"
	"goweb/crawler_distributed/itemsaver/server"
	"goweb/crawler_distributed/rpc-support"
	"testing"
	"time"
)

func TestRpcServer(t *testing.T)  {
	const host = ":8081"

	go rpc_support.ServeRpc(host, &server.RpcItemService{})

	time.Sleep(time.Second)

	client, err := rpc_support.NewRpcClient(host)
	if err != nil {
		t.Error(err)
	}

	item := model.JobProfile{
		JobNo:        "123",
		Title:        "test",
		Location:     "home",
		WorkType:     "full-time",
		BusinessUnit: "self",
		Date:         "0000-00-00",
	}
	result := ""
	err = client.Call("RpcItemService.Save", item, &result)
	if err!=nil {
		t.Error(err)
	}
	t.Log(result)
}
