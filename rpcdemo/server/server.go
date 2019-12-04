package main

import (
	rpcdemo2 "goweb/rpcdemo"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	err := rpc.Register(rpcdemo2.DemoService{})
	if err != nil {
		panic(err)
	}
	listener ,err :=net.Listen("tcp",":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}

}
