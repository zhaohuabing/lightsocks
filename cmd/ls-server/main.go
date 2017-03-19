package main

import (
	"log"
	"net"
	"github.com/gwuhaolin/lightsocks/server"
	"github.com/gwuhaolin/lightsocks/cmd"
	"github.com/gwuhaolin/lightsocks/core"
)

func main() {
	var err error
	config := cmd.ReadConfig()
	password, err := core.ParsePassword(config.Password)
	if err != nil {
		log.Fatalln(err)
	}
	localAddr, err := net.ResolveTCPAddr("tcp", config.Local)
	if err != nil {
		log.Fatalln(err)
	}
	lsServer := server.New(password, localAddr)
	lsServer.AfterListen = func(listenAddr net.Addr) {
		log.Println("lightsocks listen on " + listenAddr.String() + config.String())
	}
	log.Fatalln(lsServer.Listen())
}
