package model

import "testing"

func TestService_Ip(t *testing.T) {
	s := Service{Addr: "1.2.3.5:134"}
	if s.Port() != 134 {
		t.Error(s.Port())
	}
	if s.Ip() != "1.2.3.5" {
		t.Error(s.Ip())
	}
}
