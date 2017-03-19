package rpc

import (
	"github.com/gwuhaolin/lightsocks/app/server/util"
	"github.com/gwuhaolin/lightsocks/app/manager/model"
	"log"
	"net/rpc"
	"fmt"
	"github.com/gwuhaolin/lightsocks/app/server/service"
)

func call(serviceMethod string, args interface{}, reply interface{}) error {
	client, err := rpc.Dial("tcp", MANAGER_ADDR)
	if err != nil {
		return err
	}
	defer client.Close()
	return client.Call(serviceMethod, args, reply)
}

func CallRegisterServer(rpcListenPost int) {
	ip := util.HostPublicIP()
	server := &model.Server{
		Ip:   ip,
		Port: rpcListenPost,
	}
	log.Println(fmt.Sprintf("向 manager=%s 注册 IP=%s PORT=%d", MANAGER_ADDR, ip, rpcListenPost))
	var services []*model.Service
	err := call("Handler.RegisterServer", server, services)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("成功注册服务器", server)
	service.Recover(services)
}
