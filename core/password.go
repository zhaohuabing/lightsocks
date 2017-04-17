package core

import (
	"math/rand"
	"errors"
	"strings"
	"time"
	"encoding/base64"
)

const PASSWORD_LENGTH = 256

var ERR_INVALID_PASSWORD = errors.New("invlid password")

type Password [PASSWORD_LENGTH]byte

func init() {
	rand.Seed(time.Now().Unix())
}

//采用base64编码把密码转换为字符串
func (password *Password) String() string {
	return base64.StdEncoding.EncodeToString(password[:])
}

//解析采用base64编码的字符串获取密码
func ParsePassword(passwordString string) (*Password, error) {
	bs, err := base64.StdEncoding.DecodeString(strings.TrimSpace(passwordString))
	if err != nil || len(bs) != PASSWORD_LENGTH {
		return nil, ERR_INVALID_PASSWORD
	}
	password := Password{}
	copy(password[:], bs)
	bs = nil
	return &password, nil
}

//产生 256个byte随机组合的 密码，最后使用base64编码为字符串
//不会出现如何一个byte位出现重复
func RandPassword() *Password {
	ints := rand.Perm(PASSWORD_LENGTH)
	password := &Password{}
	sameCount := 0
	for i, v := range ints {
		password[i] = byte(v)
		if i == v {
			sameCount++
		}
	}
	//不会出现如何一个byte位出现重复
	if sameCount > 0 {
		password = nil
		return RandPassword()
	} else {
		return password
	}
}
