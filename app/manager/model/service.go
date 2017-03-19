package model

import "github.com/gwuhaolin/lightsocks/core"

type Service struct {
	Ip       string
	Port     int
	Password *core.Password
}

func AllServicesByIp(ip string) ([]*Service, error) {
	rows, err := DB.Query("Service.AllServicesByIp", ip)
	if err != nil {
		return nil, err
	}
	services := []*Service{}
	for rows.Next() {
		service := &Service{}
		if err := rows.Scan(&service.Port, &service.Password); err == nil {
			services = append(services, service)
		}
	}
	return services, nil
}
