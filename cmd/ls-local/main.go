package main

import (
	"log"
	"github.com/gwuhaolin/lightsocks/local"
	"github.com/gwuhaolin/lightsocks/cmd"
	"github.com/gwuhaolin/lightsocks/core"
	_ "net/http/pprof"
	"net/http"
)

func main() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	var err error
	config := cmd.ReadConfig()
	ssConfig, err := config.ToSsConfig()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(config)
	core.GlobalConfig = ssConfig
	local.Run()
}
