package ss

import (
	"encoding/base64"
	"math/rand"
	"time"
)

const PASSWORD_LENGTH = 256

type Password [PASSWORD_LENGTH]byte

func init() {
	rand.Seed(time.Now().Unix())
}

type Cipher struct {
	encodePassword *Password
	decodePassword *Password
}

func (cipher *Cipher) encode(bs []byte) {
	for i, v := range bs {
		bs[i] = cipher.encodePassword[v]
	}
}

func (cipher *Cipher) decode(bs []byte) {
	for i, v := range bs {
		bs[i] = cipher.decodePassword[v]
	}
}

func NewCipher(passwordStr string) (*Cipher, error) {
	bs, err := base64.StdEncoding.DecodeString(passwordStr)
	if err != nil {
		return nil, err
	}
	encodePassword := &Password{}
	decodePassword := &Password{}
	for i, v := range bs {
		encodePassword[i] = v
		decodePassword[v] = byte(i)
	}
	cipher := &Cipher{
		encodePassword: encodePassword,
		decodePassword: decodePassword,
	}
	return cipher, nil
}

//产生 256个byte随机组合的 密码，最后使用base64编码为字符串
//不会出现如何一个byte位出现重复
func RandPassword() string {
	ints := rand.Perm(PASSWORD_LENGTH)
	password := Password{}
	sameCount := 0
	for i, v := range ints {
		password[i] = byte(v)
		if i == v {
			sameCount++
		}
	}
	//不会出现如何一个byte位出现重复
	if sameCount > 0 {
		return RandPassword()
	} else {
		return base64.StdEncoding.EncodeToString(password[:])
	}
}
