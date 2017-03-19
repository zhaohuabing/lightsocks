package rpc

import (
	"github.com/gwuhaolin/lightsocks/app/manager/model"
	"errors"
	"fmt"
)

var (
	DB_ERROR = errors.New("执行数据库操作发生错误")
)

type Handler struct {
}

func (handle *Handler) RegisterServer(server *model.Server, reply []*model.Service) error {
	status, err := CallHeartbeat(server)
	if err != nil {
		return errors.New(fmt.Sprintf("manager无法连接到你提供的地址 IP=%s PORT=%d 错误:%s", server.Ip, server.Port, err))
	}
	err = server.Create()
	if err != nil {
		return DB_ERROR
	}
	_, err = server.UpdateStatus(status)
	if err != nil {
		return DB_ERROR
	}
	services, err := model.AllServicesByIp(server.Ip)
	if err != nil {
		return DB_ERROR
	}
	reply = services
	return nil
}
