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

func (cipher *Cipher) encode(byte byte) byte {
	return cipher.encodePassword[byte]
}

func (cipher *Cipher) decode(byte byte) byte {
	return cipher.decodePassword[byte]
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

func RandPassword() string {
	ints := rand.Perm(PASSWORD_LENGTH)
	password := Password{}
	for i, v := range ints {
		password[i] = byte(v)
	}
	return base64.StdEncoding.EncodeToString(password[:])
}