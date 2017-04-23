package model

import "testing"

const (
	ip   = "12.23.43.56"
	port = 2132
)

func TestServer_Addr(t *testing.T) {
	server := Server{
		Ip:   ip,
		Port: port,
	}
	addr := server.Addr()
	if addr != "12.23.43.56:2132" {
		t.Error(addr)
	}
}
