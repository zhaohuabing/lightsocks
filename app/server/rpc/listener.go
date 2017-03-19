package rpc

import (
	"github.com/gwuhaolin/lightsocks/app/server/util"
	"github.com/gwuhaolin/lightsocks/app/manager/model"
	"github.com/gwuhaolin/lightsocks/app/server/service"
	"github.com/gwuhaolin/lightsocks/core"
	"sync"
)

type Handler struct {
}

func (handle *Handler) Heartbeat(_ interface{}, replay *model.ServerStatus) error {
	replay = util.HostStatus()
	return nil
}

func (handle *Handler) AddService(_ interface{}, replay *model.Service) error {
	var wait sync.WaitGroup
	wait.Add(1)
	var err error
	go func() {
		err = service.Create(func(port int, password *core.Password) {
			replay = &model.Service{
				Port:     port,
				Password: password,
			}
			wait.Done()
		})
		wait.Done()
	}()
	wait.Wait()
	return err
}

func (handle *Handler) UpdatePassword(config *model.Service, _ interface{}) error {
	return service.UpdatePassword(config)
}

func (handle *Handler) Remove(port int, _ interface{}) error {
	return service.Remove(port)
}
