package listener

import (
	"github.com/gwuhaolin/lightsocks/app/manager/model"
	"errors"
	"fmt"
	"github.com/gwuhaolin/lightsocks/app/manager/dao"
	"github.com/gwuhaolin/lightsocks/app/manager/rpc/caller"
	"log"
)

type Handler struct {
}

/**
翻墙服务器注册到中心
 */
func (handle *Handler) RegisterServer(server *model.Server, _ *struct{}) error {
	log.Println("RegisterServer", "收到来自服务器的注册请求", server.Addr())
	status, err := caller.CallGetStatus(server)
	if err != nil {
		return errors.New(fmt.Sprintf("manager无法连接到你提供的地址 IP=%s PORT=%d 错误:%s", server.Ip, server.Port, err))
	}
	server.Status = status
	server.Alive = true
	err = dao.SaveServer(server)
	if err != nil {
		return err
	}
	return nil
}
