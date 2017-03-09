package main

import (
	"log"
	"github.com/gwuhaolin/lightsocks/server"
	"github.com/gwuhaolin/lightsocks/cmd"
	"github.com/gwuhaolin/lightsocks/core"
	"time"
	_ "net/http/pprof"
	"net/http"
)

func main() {
	go func() {
		http.ListenAndServe("localhost:6061", nil)
	}()

	var err error
	defaultConfig := &cmd.Config{
		Local:    ":8011",
		Password: core.RandPassword().String(),
		Timeout:  10 * time.Second,
	}
	cmd.ReadConfig(defaultConfig)
	ssConfig, err := defaultConfig.ToSsConfig()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(defaultConfig)
	core.GlobalConfig = ssConfig
	server.Run()
}
