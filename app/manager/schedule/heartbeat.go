package schedule

import (
	"github.com/gwuhaolin/lightsocks/app/manager/model"
	"github.com/gwuhaolin/lightsocks/app/manager/rpc"
	"log"
)

func CheckAllServersAlive() {
	log.Println("schedule", "CheckAllServersAlive")
	servers, _ := model.GetAllServers()
	for _, server := range servers {
		go func() {
			status, err := rpc.CallHeartbeat(server)
			if err != nil {
				server.UpdateAlive(false)
				server.UpdateStatus(nil)
			} else {
				server.UpdateAlive(true)
				server.UpdateStatus(status)
			}
		}()
	}
}
