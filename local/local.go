package local

import (
	"net"
	"log"
	"github.com/gwuhaolin/lightsocks/ss"
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
		go handleConn(userConn)
	}
}

func handleConn(userConn *net.TCPConn) {
	defer userConn.Close()
	server, err := ss.DialServer()
	if err != nil {
		log.Println(err)
		return
	}
	defer server.Close()
	//进行转发
	go ss.EncodeCopy(server, userConn)
	ss.DecodeCopy(userConn, server)
}
