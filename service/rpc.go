package service

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRpc(host string, service interface{}) error {
	err := rpc.Register(service)
	if err != nil {
		return err
	}
	listener ,err :=net.Listen("tcp",host)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
	//still needed, otherwise compile error
	return nil
}

func NewRpcClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err!=nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), nil
}
