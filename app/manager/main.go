package main

import (
	"github.com/gwuhaolin/lightsocks/app/manager/rpc"
	"github.com/gwuhaolin/lightsocks/app/manager/model"
	"github.com/gwuhaolin/lightsocks/app/manager/schedule"
)

func main() {
	model.InitDB()
	go schedule.Start()
	rpc.ListenRPC()
}
