package local

import (
	"net"
	"log"
	"time"
	"github.com/gwuhaolin/lightsocks/core"
)

type LsLocal struct {
	*core.SecureSocket
}

func (local *LsLocal) Listen() {
	listener, err := net.ListenTCP("tcp", local.LocalAddr)
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
		go local.handleConn(userConn)
	}
}

func (local *LsLocal) handleConn(userConn *net.TCPConn) {
	defer userConn.Close()
	server, err := local.DialServer()
	if err != nil {
		log.Println(err)
		return
	}
	defer server.Close()
	server.SetLinger(0)
	server.SetDeadline(time.Now().Add(core.TIMEOUT))
	//进行转发
	go local.EncodeCopy(server, userConn)
	local.DecodeCopy(userConn, server)
}
