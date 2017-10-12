package main

import (
	"log"
	"net"
	"fmt"
	"github.com/gwuhaolin/lightsocks/server"
	"github.com/gwuhaolin/lightsocks/cmd"
	"github.com/gwuhaolin/lightsocks/core"
)

var version = "master"

func main() {
	var err error
	config := cmd.ReadConfig()
	password, err := core.ParsePassword(config.Password)
	if err != nil {
		log.Fatalln(err)
	}
	listenAddr, err := net.ResolveTCPAddr("tcp", config.ListenAddr)
	if err != nil {
		log.Fatalln(err)
	}

	// 启动 server 端并监听
	lsServer := server.New(password, listenAddr)
	lsServer.AfterListen = func(listenAddr net.Addr) {
		log.Printf("lightsocks-server:%s 启动成功 监听在 %s\n", version, listenAddr.String())
		log.Println("使用配置：", fmt.Sprintf(`
本地监听地址 listen：
%s
密码 password：
%s
	`, listenAddr, password))
	}
	log.Fatalln(lsServer.Listen())
}
