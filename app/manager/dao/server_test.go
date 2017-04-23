package dao

import (
	"testing"
	"github.com/gwuhaolin/lightsocks/app/manager/model"
)

const (
	ip   = "12.23.43.56"
	port = 2132
)

func init() {
	InitDB()
}

func TestGetAllServers(t *testing.T) {
	servers, err := FindAllServers()
	if err != nil {
		t.Error(err)
	}
	t.Log(servers)
}

func TestServer_Save(t *testing.T) {
	server := &model.Server{
		Ip:   ip,
		Port: port,
	}
	err := SaveServer(server)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(server)
	}
}
