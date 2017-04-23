package caller

import (
	"net/rpc"
	"github.com/gwuhaolin/lightsocks/app/manager/model"
	"log"
)

// 向对应到的服务器发送RPC请求
func call(server *model.Server, serviceMethod string, args interface{}, reply interface{}) error {
	client, err := rpc.Dial("tcp", server.Addr())
	if err != nil {
		return err
	}
	defer client.Close()
	return client.Call(serviceMethod, args, reply)
}

//向发送服务器获取状态请求
func CallGetStatus(server *model.Server) (*model.ServerStatus, error) {
	status := &model.ServerStatus{}
	log.Println("CallGetStatus", "向发送服务器获取状态请求", server.Addr())
	err := call(server, "Handler.GetStatus", &struct{}{}, status)
	if err != nil {
		return nil, err
	}
	log.Println("CallGetStatus", "收到来自服务器的状态", server.Addr(), status)
	return status, nil
}

//让服务器新开一个服务
func CallAddService(server *model.Server) (*model.Service, error) {
	service := &model.Service{}
	log.Println("CallAddService", "让服务器新开一个服务", server.Addr())
	err := call(server, "Handler.AddService", &struct{}{}, service)
	if err != nil {
		return nil, err
	}
	return service, nil
}

//让服务器销毁一个服务
func CallRemoveService(server *model.Server, port int) error {
	log.Println("CallAddService", "让服务器销毁一个服务", server.Addr())
	return call(server, "Handler.RemoveService", port, &struct{}{})
}
