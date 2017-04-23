package model

import (
	"strings"
	"strconv"
)

type Service struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
}

func (service *Service) Port() int {
	arr := strings.Split(service.Addr, ":")
	if len(arr) == 2 {
		port, _ := strconv.Atoi(arr[1])
		return port
	} else {
		return 0
	}
}

func (service *Service) Ip() string {
	arr := strings.Split(service.Addr, ":")
	if len(arr) == 2 {
		return arr[0]
	} else {
		return ""
	}
}
