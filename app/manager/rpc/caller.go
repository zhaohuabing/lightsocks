package rpc

import (
	"net/rpc"
	"github.com/gwuhaolin/lightsocks/app/manager/model"
)

func call(server *model.Server, serviceMethod string, args interface{}, reply interface{}) error {
	client, err := rpc.Dial("tcp", server.Addr())
	if err != nil {
		return err
	}
	defer client.Close()
	return client.Call(serviceMethod, args, reply)
}

func CallHeartbeat(server *model.Server) (*model.ServerStatus, error) {
	var status *model.ServerStatus
	err := call(server, "Handler.Heartbeat", nil, status)
	if err != nil {
		return nil, err
	}
	return status, nil
}

func CallAddService(server *model.Server) (*model.Service, error) {
	var service *model.Service
	err := call(server, "Handler.AddService", nil, service)
	if err != nil {
		return nil, err
	}
	return service, nil
}

func CallUpdatePassword(server *model.Server, config *model.Service) error {
	return call(server, "Handler.UpdatePassword", config, nil)
}

func CallRemove(server *model.Server, port int) error {
	return call(server, "Handler.Remove", port, nil)
}
