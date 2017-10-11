package main

import (
	"log"
	"net"
	"github.com/gwuhaolin/lightsocks/server"
	"github.com/gwuhaolin/lightsocks/cmd"
	"github.com/gwuhaolin/lightsocks/core"
	"fmt"
)

var version = "master"

func main() {
	var err error
	config := cmd.ReadConfig()
	password, err := core.ParsePassword(config.Password)
	if err != nil {
		log.Fatalln(err)
	}
	localAddr, err := net.ResolveTCPAddr("tcp", config.ListenAddr)
	if err != nil {
		log.Fatalln(err)
	}
	lsServer := server.New(password, localAddr)
	lsServer.AfterListen = func(listenAddr net.Addr) {
		log.Printf("lightsocks-server:%s 启动成功 监听在 %s\n", version, listenAddr.String())
		log.Println("使用配置：", fmt.Sprintf(`
本地监听地址 listen：
%s
密码 password：
%s
	`, config.ListenAddr, config.Password))
	}
	log.Fatalln(lsServer.Listen())
}
