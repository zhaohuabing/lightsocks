package local

import (
	"net"
	"io"
	"log"
	"github.com/gwuhaolin/lightsocks/ss"
)

var Config *ss.Config

func handleConn(userConn net.Conn) {
	defer userConn.Close()
	server, err := ss.Dial(Config)
	if err != nil {
		log.Println(err)
		return
	}
	defer server.Conn.Close()
	go io.Copy(server, userConn)
	io.Copy(userConn, server)
}

func Run() {
	listener, err := net.Listen("tcp", Config.Local)
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	defer func() {
		log.Println(recover())
	}()
	for {
		userConn, _ := listener.Accept()
		go handleConn(userConn)
	}
}