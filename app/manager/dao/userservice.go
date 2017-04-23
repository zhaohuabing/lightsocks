package dao

import "github.com/gwuhaolin/lightsocks/app/manager/model"

func CreateUserService(userService *model.UserService) error {
	_, err := connPool.Exec("UserService.Create", userService.Email, userService.Addr, userService.UUID, userService.Device)
	return err
}
