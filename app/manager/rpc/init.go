package rpc

import (
	"net"
	"log"
	"net/rpc"
)

const (
	LISTEN_ADDR = ":9110"
)

func ListenRPC() {
	addr, err := net.ResolveTCPAddr("tcp", LISTEN_ADDR)
	if err != nil {
		log.Fatalln(err)
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	rpc.Register(new(Handler))

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}
