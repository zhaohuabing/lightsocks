package model

import (
	"fmt"
	"time"
)

type ServerStatus struct {
	ServiceCount int `json:"serviceCount"`
}

type Server struct {
	Ip       string `json:"ip"`
	Port     int `json:"port"`
	Alive    bool `json:"alive"`
	Status   *ServerStatus `json:"status"`
	UpdateAt time.Time `json:"updateAt"`
}

func (server *Server) Addr() string {
	return fmt.Sprintf("%s:%d", server.Ip, server.Port)
}
