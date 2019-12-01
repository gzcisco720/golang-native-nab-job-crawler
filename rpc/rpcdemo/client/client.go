package main

import (
	"fmt"
	"goweb/rpc/rpcdemo"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err!=nil {
		panic("connect error")
	}
	client := jsonrpc.NewClient(conn)

	var result float64
	err = client.Call("DemoService.Div", rpcdemo.Args{2,3}, &result)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
