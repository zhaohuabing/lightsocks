package main

import (
	"net"
	"github.com/gwuhaolin/lightsocks/ss"
	"io"
	"log"
	"os"
)

var Config *ss.Config

func handleConn(userConn net.Conn) {
	defer userConn.Close()
	server, err := ss.Dial(Config.Server, Config.Cipher)
	if err != nil {
		log.Println(err)
		return
	}
	defer server.Close()
	go io.Copy(server, userConn)
	io.Copy(userConn, server)
}

func Run() {
	listener, err := net.Listen("tcp", Config.Local)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	for {
		userConn, _ := listener.Accept()
		go handleConn(userConn)
	}
}

func main() {
	filePath := os.Args[1]
	var err error
	Config, err = ss.ParseConfig(filePath)
	if err != nil {
		panic(err)
	}
	Run()
}
