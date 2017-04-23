package service

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/gwuhaolin/lightsocks/app/manager/model"
	"github.com/gwuhaolin/lightsocks/app/manager/rpc/caller"
	"github.com/gwuhaolin/lightsocks/app/manager/dao"
)

// 用户申请新的服务
func UserApplyService(user *model.User, device *model.DeviceInfo) (*model.Service, error) {
	server, err := dao.GetNextServer()
	if err != nil {
		return nil, err
	}
	newService, err := caller.CallAddService(server)
	userService := &model.UserService{
		Email:  user.Email,
		Device: device,
		UUID:   device.UUID,
		Addr:   newService.Addr,
	}
	err = dao.CreateUserService(userService)
	if err != nil {
		return nil, err
	}
	return newService, nil
}

// 检验用户对应的账号密码是否匹配
// 如果匹配成功会重置用户的token再返回用户的信息
func UserAuth(email, password string) (*model.User, error) {
	user, err := dao.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		return nil, err

	}
	return user, nil
}
