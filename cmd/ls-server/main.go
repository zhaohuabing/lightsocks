package main

import (
	"log"
	"github.com/gwuhaolin/lightsocks/server"
	"github.com/gwuhaolin/lightsocks/ss"
)

func main() {
	var err error
	Config, err := ss.ParseConfig()
	if err != nil {
		log.Fatalln(err)
	}
	server.Config = Config
	server.Run()
}
