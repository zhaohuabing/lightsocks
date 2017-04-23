package dao

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/gwuhaolin/lightsocks/app/manager/model"
)

// 更新用户信息
func UpdateUser(user *model.User) error {
	_, err := connPool.Exec("UpdateUser", &user.Password, &user.Token, &user.Balance, &user.Cost, &user.Config, &user.Email)
	return err
}

// 新建一个用户
func CreateUser(email, password string) error {
	cPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = connPool.Exec("CreateUser", email, cPassword, model.RandToken())
	if err != nil {
		return err
	}
	return nil
}

// 获取对应邮箱的用户
func GetUserByEmail(email string) (*model.User, error) {
	user := &model.User{Email: email}
	err := connPool.QueryRow("GetUserByEmail", email).Scan(&user.Password, &user.Token, &user.Balance, &user.Cost, &user.Config)
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

// 获取对应邮箱的用户
func GetUserByToken(token string) (*model.User, error) {
	user := &model.User{Token: token}
	err := connPool.QueryRow("GetUserByToken", token).Scan(&user.Email, &user.Password, &user.Balance, &user.Cost, &user.Config)
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}
}
