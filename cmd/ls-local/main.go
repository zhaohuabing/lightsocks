package main

import (
	"log"
	"github.com/gwuhaolin/lightsocks/local"
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
	lsLocal := &local.LsLocal{SecureSocket: secureSocket}
	lsLocal.Listen()
}
