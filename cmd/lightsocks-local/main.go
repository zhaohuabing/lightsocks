package main

import (
	"fmt"
	"github.com/zhaohuabing/lightsocks"
	"github.com/zhaohuabing/lightsocks/cmd"
	"log"
	"net"
	"os"
)

const (
	DefaultListenAddr = ":7448"
)

func main() {
	log.SetFlags(log.Lshortfile)

	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("usage: light-client listenAddr serverAddr Password, for example ./light-client :8080 10.75.8.83:12345 o475dVEctO+BuOXDRsQ...  ")
		os.Exit(1)
	}

	// 默认配置
	config := &cmd.Config{
		ListenAddr: args[0],
		RemoteAddr: args[1],
		Password:   args[2],
	}

	if len(args) == 4 {

	}

	// 启动 local 端并监听
	lsLocal, err := lightsocks.NewLsLocal(config.Password, config.ListenAddr, config.RemoteAddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println()
	log.Fatalln(lsLocal.Listen(func(listenAddr net.Addr) {
		fmt.Println(fmt.Sprintf(`
lightsocks-local 启动成功，配置如下：
本地监听地址：%s
远程服务地址：%s
密码：%s...`, listenAddr, config.RemoteAddr, config.Password[0:20]))
	}))
}
