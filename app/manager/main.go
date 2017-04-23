package main

import (
	"github.com/gwuhaolin/lightsocks/app/manager/schedule"
	"github.com/gwuhaolin/lightsocks/app/manager/http"
	"github.com/gwuhaolin/lightsocks/app/manager/dao"
	"github.com/gwuhaolin/lightsocks/app/manager/rpc/listener"
)

func main() {
	dao.InitDB()
	go schedule.Start()
	go listener.ListenRPC()
	http.ListenHTTP()
}
