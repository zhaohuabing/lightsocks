package util

import (
	"net/http"
	"io/ioutil"
	"github.com/gwuhaolin/lightsocks/app/manager/model"
	"log"
	"github.com/gwuhaolin/lightsocks/app/server/service"
)

/**
获取当前计算机的公网IP
 */
func HostPublicIP() string {
	var apiList = []string{"https://api.ipify.org", "https://www.trackip.net/ip", "https://ident.me"}
	var res *http.Response
	var err error
	var ip []byte
	for _, api := range apiList {
		res, err = http.Get(api)
		if err == nil {
			ip, err = ioutil.ReadAll(res.Body)
			if err == nil {
				res.Body.Close()
				break
			}
		}
	}
	if len(ip) == 0 {
		log.Fatalln("你的服务器不能连接到外网")
	}
	return string(ip)
}

/**
获取当前服务器的状态
 */
func HostStatus() *model.ServerStatus {
	status := &model.ServerStatus{
		ServiceCount: len(service.LsServerMap),
	}
	return status
}
