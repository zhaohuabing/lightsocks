package main

import (
	"log"
	"github.com/gwuhaolin/lightsocks/local"
	"github.com/gwuhaolin/lightsocks/cmd"
	"github.com/gwuhaolin/lightsocks/core"
	"net"
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
	serverAddr, err := net.ResolveTCPAddr("tcp", config.RemoteAddr)
	if err != nil {
		log.Fatalln(err)
	}
	lsLocal := local.New(password, localAddr, serverAddr)
	lsLocal.AfterListen = func(listenAddr net.Addr) {
		log.Printf("lightsocks:%s listen on "+listenAddr.String()+config.String(), version)
	}
	log.Fatalln(lsLocal.Listen())
}
