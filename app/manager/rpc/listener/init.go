package listener

import (
	"net"
	"log"
	"net/rpc"
)

const (
	//中心管理服务器RPC监听地址
	MANAGER_RPC_LISTEN_ADDR = ":1501"
)

func ListenRPC() {
	addr, err := net.ResolveTCPAddr("tcp", MANAGER_RPC_LISTEN_ADDR)
	if err != nil {
		log.Fatalln(err)
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	rpc.Register(new(Handler))
	log.Println("ListenRPC", "成功启动RPC监听", addr.String())
	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}
