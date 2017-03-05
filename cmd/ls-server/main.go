package main

import (
	"log"
	"github.com/gwuhaolin/lightsocks/server"
	"github.com/gwuhaolin/lightsocks/cmd"
	"github.com/gwuhaolin/lightsocks/ss"
)

func main() {
	var err error
	config := cmd.ReadConfig()
	ssConfig, err := config.ToSsConfig()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(config)
	ss.GlobalConfig = ssConfig
	server.Run()
}
