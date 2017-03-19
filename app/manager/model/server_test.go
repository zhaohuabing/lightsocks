package model

import "testing"

func init() {
	InitDB()
}

func TestServer_Addr(t *testing.T) {
	server := Server{
		Ip:   "12.23.43.56",
		Port: 2132,
	}
	addr := server.Addr()
	if addr != "12.23.43.56:2132" {
		t.Error(addr)
	}
}

func TestGetAllServers(t *testing.T) {
	servers, err := GetAllServers()
	if err != nil {
		t.Error(err)
	}
	for _, server := range servers {
		t.Log(server)
	}
}

func TestServer_Create(t *testing.T) {
	server := Server{
		Ip:   "12.43.65.43",
		Port: 1234,
	}
	err := server.Create()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(server)
	}
}
