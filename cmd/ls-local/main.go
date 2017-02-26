package main

import (
	"log"
	"github.com/gwuhaolin/lightsocks/ss"
	"github.com/gwuhaolin/lightsocks/local"
)

func main() {
	var err error
	Config, err := ss.ParseConfig()
	if err != nil {
		log.Fatalln(err)
	}
	local.Config = Config
	local.Run()
}
