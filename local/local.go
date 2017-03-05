package local

import (
	"net"
	"log"
	"github.com/gwuhaolin/lightsocks/ss"
	"time"
)

func Run() {
	listener, err := net.ListenTCP("tcp", ss.GlobalConfig.LocalAddr)
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	defer func() {
		log.Println(recover())
	}()
	for {
		userConn, _ := listener.AcceptTCP()
		userConn.SetLinger(0)
		go handleConn(userConn)
	}
}

func handleConn(userConn *net.TCPConn) {
	defer userConn.Close()
	server, err := ss.DialServer()
	server.SetLinger(0)
	server.SetDeadline(time.Now().Add(ss.GlobalConfig.Timeout))
	if err != nil {
		log.Println(err)
		return
	}
	defer server.Close()
	//进行转发
	go ss.EncodeCopy(server, userConn)
	ss.DecodeCopy(userConn, server)
}
