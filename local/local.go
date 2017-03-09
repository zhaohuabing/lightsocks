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
		userConn, _ := listener.AcceptTCP()
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
	server.SetLinger(0)
	server.SetDeadline(time.Now().Add(core.GlobalConfig.Timeout))
	defer server.Close()
	//进行转发
	go core.EncodeCopy(server, userConn)
	core.DecodeCopy(userConn, server)
}
