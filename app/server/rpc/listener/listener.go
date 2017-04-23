package listener

import (
	"github.com/gwuhaolin/lightsocks/app/server/util"
	"github.com/gwuhaolin/lightsocks/app/manager/model"
	"github.com/gwuhaolin/lightsocks/app/server/service"
	"github.com/gwuhaolin/lightsocks/core"
	"sync"
	"fmt"
	"log"
)

type Handler struct {
}

//告诉管理者我的状态
func (handle *Handler) GetStatus(_ *struct{}, replay *model.ServerStatus) error {
	replay = util.HostStatus()
	log.Println("GetStatus", "告诉管理者我的状态", *replay)
	return nil
}

func (handle *Handler) AddService(_ *struct{}, replay *model.Service) error {
	var wait sync.WaitGroup
	wait.Add(1)
	var err error
	go func() {
		err = service.CreateService(func(port int, password *core.Password) {
			ip := util.HostPublicIP()
			replay = &model.Service{
				Addr:     fmt.Sprintf("%s:%d", ip, port),
				Password: password.String(),
			}
			wait.Done()
		})
		wait.Done()
	}()
	wait.Wait()
	return err
}

func (handle *Handler) RemoveService(port int, _ *struct{}) error {
	return service.RemoveService(port)
}
