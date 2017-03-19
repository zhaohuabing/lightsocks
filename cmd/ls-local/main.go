package main

import (
	"log"
	"github.com/gwuhaolin/lightsocks/local"
	"github.com/gwuhaolin/lightsocks/cmd"
	"github.com/gwuhaolin/lightsocks/core"
	"net"
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
	serverAddr, err := net.ResolveTCPAddr("tcp", config.Server)
	if err != nil {
		log.Fatalln(err)
	}
	lsLocal := local.New(password, localAddr, serverAddr)
	lsLocal.AfterListen = func(listenAddr net.Addr) {
		log.Println("lightsocks listen on " + listenAddr.String() + config.String())
	}
	log.Fatalln(lsLocal.Listen())
}
