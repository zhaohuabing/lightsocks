package listener

import (
	"net/rpc"
	"net"
	"log"
	"github.com/gwuhaolin/lightsocks/app/server/rpc/caller"
)

func ListenRPC() {
	//自动获取操作系统上可用的下一个监听端口
	listenAddr, err := net.ResolveTCPAddr("tcp", ":0")
	tcpListener, err := net.ListenTCP("tcp", listenAddr)
	if err != nil {
		log.Fatalln(err)
	}
	rpc.Register(new(Handler))
	// 获取当前监听地址
	listenAddr, _ = net.ResolveTCPAddr("tcp", tcpListener.Addr().String())
	log.Println("ListenRPC", "成功启动RPC监听", listenAddr)

	// 向管理服务器注册
	go func() {
		err = caller.CallRegisterServer(listenAddr.Port)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	for {
		conn, err := tcpListener.AcceptTCP()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}
