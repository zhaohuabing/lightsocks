package service

import (
	"net"
	"github.com/gwuhaolin/lightsocks/server"
	"github.com/gwuhaolin/lightsocks/core"
	"github.com/gwuhaolin/lightsocks/app/manager/model"
	"log"
	"fmt"
	"errors"
)

var (
	ERR_NO_SERVICE  = errors.New("没有服务监听在对应的端口")
	LsServerMap    map[int]*server.LsServer
)

//更新服务端配置
func UpdatePassword(config *model.Service) error {
	lsServer := LsServerMap[config.Port]
	if lsServer == nil {
		return ERR_NO_SERVICE
	}
	lsServer.Update(core.NewCipher(config.Password))
	return nil
}

//添加一个服务器并且让它运行监听
//自动寻找一个可用的端口监听
func Create(onSucc func(port int, password *core.Password)) error {
	localAddr := &net.TCPAddr{}
	password := core.RandPassword()
	lsServer := server.New(password, localAddr)
	lsServer.AfterListen = func(listenAddr net.Addr) {
		tcpAddr, _ := net.ResolveTCPAddr("tcp", listenAddr.String())
		port := tcpAddr.Port
		LsServerMap[port] = lsServer
		if onSucc != nil {
			onSucc(port, password)
		}
	}
	return lsServer.Listen()
}

//根据服务器
func Recover(services []*model.Service) {
	for _, config := range services {
		go func() {
			localAddr := &net.TCPAddr{Port: config.Port}
			lsServer := server.New(config.Password, localAddr)
			lsServer.AfterListen = func(listenAddr net.Addr) {
				LsServerMap[config.Port] = lsServer
				log.Println(fmt.Sprintf("恢复服务,监听地址=%s", listenAddr))
			}
			lsServer.Listen()
		}()
	}
}

//删除一个正在运行的服务端并且释放资源
func Remove(port int) error {
	lsServer := LsServerMap[port]
	if lsServer == nil {
		return ERR_NO_SERVICE
	}
	lsServer.Close()
	delete(LsServerMap, port)
	return nil
}
