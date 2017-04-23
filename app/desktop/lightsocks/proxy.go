package lightsocks

import (
	"github.com/gwuhaolin/lightsocks/local"
	"github.com/gwuhaolin/lightsocks/core"
	"github.com/gwuhaolin/lightsocks/app/desktop/util"
	"github.com/gwuhaolin/lightsocks/app/manager/model"
	"net"
)

var (
	lsLocal *local.LsLocal
)

func ListenLightsocks(service *model.Service, succ func(listenAddr net.Addr)) error {
	var err error
	if lsLocal != nil {
		lsLocal.Close()
		lsLocal = nil
	}
	password, err := core.ParsePassword(service.Password)
	if err != nil {
		return err
	}
	localAddr, err := net.ResolveTCPAddr("tcp", util.LOCAL_SOCKS_ADDR)
	if err != nil {
		return err
	}
	serverAddr, err := net.ResolveTCPAddr("tcp", service.Addr)
	if err != nil {
		return err
	}
	lsLocal = local.New(password, localAddr, serverAddr)
	lsLocal.AfterListen = succ
	return lsLocal.Listen()
}
