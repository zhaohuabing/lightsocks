package main

import (
	"log"
	"github.com/gwuhaolin/lightsocks/server"
	"github.com/gwuhaolin/lightsocks/cmd"
)

func main() {
	var err error
	config := cmd.ReadConfig()
	secureSocket, err := config.ToSecureSocket()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(config)
	lsServer := &server.LsServer{SecureSocket: secureSocket}
	lsServer.Listen()
}
