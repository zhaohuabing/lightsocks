package dao

import (
	"log"
	"github.com/gwuhaolin/lightsocks/app/manager/model"
	"time"
)

func SaveServer(server *model.Server) error {
	server.UpdateAt = time.Now()
	_, err := connPool.Exec("SaveServer", &server.Ip, &server.Alive, &server.Port, &server.Status, &server.UpdateAt)
	return err
}

func FindAllServers() ([]*model.Server, error) {
	rows, err := connPool.Query("FindAllServers")
	if err != nil {
		return nil, err
	}
	servers := []*model.Server{}
	for rows.Next() {
		server := &model.Server{}
		if err := rows.Scan(&server.Ip, &server.Port, &server.Alive, &server.Status, &server.UpdateAt); err == nil {
			servers = append(servers, server)
		} else {
			log.Println(err)
		}
	}
	return servers, nil
}

// 获取下一个用来开启新服务的服务器
// 选举算法
func GetNextServer() (*model.Server, error) {
	server := &model.Server{}
	err := connPool.QueryRow("GetNextServer").Scan(&server.Ip, &server.Port, &server.Alive, &server.Status, &server.UpdateAt)
	if err != nil {
		return nil, err
	}
	return server, nil
}
