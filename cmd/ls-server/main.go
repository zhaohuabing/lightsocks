package main

import (
	"log"
	"net"
	"github.com/gwuhaolin/lightsocks/server"
	"github.com/gwuhaolin/lightsocks/cmd"
	"github.com/gwuhaolin/lightsocks/core"
)

var version = "master"

func main() {
	var err error
	config := cmd.ReadConfig()
	password, err := core.ParsePassword(config.Password)
	if err != nil {
		log.Fatalln(err)
	}
	localAddr, err := net.ResolveTCPAddr("tcp", config.ListenAddr)
	if err != nil {
		log.Fatalln(err)
	}
	lsServer := server.New(password, localAddr)
	lsServer.AfterListen = func(listenAddr net.Addr) {
		log.Printf("lightsocks:%s listen on %s %s", version, listenAddr.String(), config.String())
	}
	log.Fatalln(lsServer.Listen())
}
