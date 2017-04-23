package dao

import (
	"github.com/gwuhaolin/lightsocks/app/manager/model"
)

func FindAllServicesByIp(ip string) ([]*model.Service, error) {
	rows, err := connPool.Query("FindAllServicesByIp", ip)
	if err != nil {
		return nil, err
	}
	services := []*model.Service{}
	for rows.Next() {
		service := &model.Service{}
		if err := rows.Scan(&service.Addr, &service.Password); err == nil {
			services = append(services, service)
		}
	}
	return services, nil
}

func CreateService(service *model.Service) error {
	_, err := connPool.Exec("CreateService", &service.Addr, &service.Password, service.Ip())
	return err
}
