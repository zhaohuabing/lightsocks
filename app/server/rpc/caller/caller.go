package caller

import (
	"github.com/gwuhaolin/lightsocks/app/server/util"
	"github.com/gwuhaolin/lightsocks/app/manager/model"
	"log"
	"net/rpc"
	"fmt"
	"github.com/gwuhaolin/lightsocks/app/manager/rpc/listener"
)

// 发送RPC请求到中心管理服务器
func call(serviceMethod string, args interface{}, reply interface{}) error {
	client, err := rpc.Dial("tcp", listener.MANAGER_RPC_LISTEN_ADDR)
	if err != nil {
		return err
	}
	defer client.Close()
	return client.Call(serviceMethod, args, reply)
}

//向中心管理服务器注册服务器
func CallRegisterServer(rpcListenPost int) error {
	ip := util.HostPublicIP()
	server := &model.Server{
		Ip:   ip,
		Port: rpcListenPost,
	}
	log.Println("CallRegisterServer", fmt.Sprintf("向 manager=%s 注册我的 IP=%s PORT=%d", listener.MANAGER_RPC_LISTEN_ADDR, ip, rpcListenPost))
	err := call("Handler.RegisterServer", server, &struct{}{})
	if err != nil {
		return err
	}
	log.Println("CallRegisterServer", "成功注册服务器", server)
	return nil
}
