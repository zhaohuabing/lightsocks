package model

import (
	"fmt"
	"github.com/jackc/pgx"
	"time"
	"log"
)

type ServerStatus struct {
	ServiceCount int
}

type Server struct {
	Ip       string
	Port     int
	Alive    bool
	Status   *ServerStatus
	UpdateAt time.Time
}

func (server *Server) Addr() string {
	return fmt.Sprintf("%s:%d", server.Ip, server.Port)
}

func (server *Server) Create() (error) {
	return DB.QueryRow("Server.Create", &server.Ip, server.Port).Scan(&server.UpdateAt)
}

func (server *Server) UpdateStatus(status *ServerStatus) (pgx.CommandTag, error) {
	return DB.Exec("Server.UpdateStatus", status, &server.Ip)
}

func (server *Server) UpdateAlive(alive bool) (pgx.CommandTag, error) {
	return DB.Exec("Server.UpdateAlive", alive, &server.Ip, )
}

func GetAllServers() ([]*Server, error) {
	rows, err := DB.Query("Server.GetAllServers")
	if err != nil {
		return nil, err
	}
	servers := []*Server{}
	for rows.Next() {
		server := &Server{}
		if err := rows.Scan(&server.Ip, &server.Port, &server.Alive, &server.Status, &server.UpdateAt); err == nil {
			servers = append(servers, server)
		} else {
			log.Println(err)
		}
	}
	return servers, nil
}
