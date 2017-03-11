package main

import (
	"net"
	"time"
	"github.com/gwuhaolin/lightsocks/server"
	"github.com/gwuhaolin/lightsocks/core"
)

type ServerConfig struct {
	Timeout  time.Duration `json:"timeout"`
	Password string `json:"password"`
}

var lsServerMap map[string]*server.LsServer

//添加一个服务器并且让它运行监听
//更新服务端配置
func Save(id string, config *ServerConfig) error {
	lsServer := lsServerMap[id]
	password, err := core.ParsePassword(config.Password)
	if err != nil {
		return err
	}
	if lsServer == nil {
		localAddr := &net.TCPAddr{}
		lsServer = &server.LsServer{
			SecureSocket: core.NewSecureSocket(time.Second, password, localAddr, nil),
		}
		lsServerMap[id] = lsServer
		lsServer.Listen()
	} else {
		lsServer.Update(config.Timeout, core.NewCipher(password))
	}
	return nil
}

//删除一个正在运行的服务端并且释放资源
func Remove(id string) {
	lsServer := lsServerMap[id]
	if lsServer != nil {
		lsServer.Close()
		delete(lsServerMap, id)
	}
}
