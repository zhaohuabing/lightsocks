package local

import (
	"net"
	"log"
	"time"
	"github.com/gwuhaolin/lightsocks/core"
)

func Run() {
	listener, err := net.ListenTCP("tcp", core.GlobalConfig.LocalAddr)
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	for {
		userConn, err := listener.AcceptTCP()
		if err != nil {
			continue
		}
		//userConn被关闭时直接清除所有数据 不管没有发送的数据
		userConn.SetLinger(0)
		go handleConn(userConn)
	}
}

func handleConn(userConn *net.TCPConn) {
	defer userConn.Close()
	server, err := core.DialServer()
	if err != nil {
		log.Println(err)
		return
	}
	defer server.Close()
	server.SetLinger(0)
	server.SetDeadline(time.Now().Add(core.GlobalConfig.Timeout))
	//进行转发
	go core.EncodeCopy(server, userConn)
	core.DecodeCopy(userConn, server)
}
