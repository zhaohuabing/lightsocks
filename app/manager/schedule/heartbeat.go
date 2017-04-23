package schedule

import (
	"log"
	"github.com/gwuhaolin/lightsocks/app/manager/dao"
	"github.com/gwuhaolin/lightsocks/app/manager/rpc/caller"
)

func CheckAllServersAlive() {
	log.Println("schedule", "CheckAllServersAlive")
	servers, _ := dao.FindAllServers()
	for _, server := range servers {
		go func() {
			status, err := caller.CallGetStatus(server)
			if err != nil {
				server.Alive = false
				server.Status = nil
				dao.SaveServer(server)
			} else {
				server.Alive = true
				server.Status = status
				dao.SaveServer(server)
			}
		}()
	}
}
