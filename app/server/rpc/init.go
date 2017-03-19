package rpc

import (
	"net/rpc"
	"net"
	"log"
)

const (
	MANAGER_ADDR = ":9110"
)

func ListenRPC() {
	//自动获取操作系统上可用的下一个监听端口
	listenAddr, err := net.ResolveTCPAddr("tcp", ":0")
	listener, err := net.ListenTCP("tcp", listenAddr)
	if err != nil {
		panic(err)
	}

	rpc.Register(new(Handler))

	listenAddr, _ = net.ResolveTCPAddr("tcp", listener.Addr().String())
	log.Println("启动RPC监听")
	CallRegisterServer(listenAddr.Port)

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}
