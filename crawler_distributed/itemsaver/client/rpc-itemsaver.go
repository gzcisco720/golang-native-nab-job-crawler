package client

import (
	"goweb/crawler/model"
	rpc_support "goweb/crawler_distributed/rpc-support"
)

func RpcItemSaver(host string) (chan interface{}, error) {
	client, err := rpc_support.NewRpcClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan interface{})
	go func() {
		for {
			item := <-out
			if i, ok := item.(model.JobProfile); ok{
				result := " "
				client.Call("RpcItemService.Save", i , &result)
			}
		}
	}()
	return out, nil
}