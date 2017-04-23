package model

import (
	"crypto/rand"
	"encoding/base64"
)

type UserConfig struct {
}

type User struct {
	// 邮件
	Email string `json:"email"`
	// 加密后的密码
	Password []byte `json:"-"`
	// 鉴权token 172个字符
	Token string `json:"-"`
	// 用户余额 单位分
	Balance float64 `json:"balance"`
	// 用户每日花费 单位分
	Cost float64 `json:"cost"`
	// 用户配置
	Config UserConfig `json:"config"`
}

// 生成token
func RandToken() string {
	bs := make([]byte, 128)
	rand.Read(bs)
	return base64.StdEncoding.EncodeToString(bs)
}
