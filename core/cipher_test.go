package core

import (
	"testing"
	"reflect"
	"crypto/rand"
)

const (
	MB = 1024 * 1024
)

func TestRandPassword(t *testing.T) {
	password := RandPassword()
	t.Log(password)
}

func TestNewCipher(t *testing.T) {
	password := RandPassword()
	t.Log(password)
	cipher := NewCipher(password)
	org := make([]byte, PasswordLength)
	for i := 0; i < PasswordLength; i++ {
		org[i] = byte(i)
	}
	tmp := make([]byte, PasswordLength)
	copy(tmp, org)
	t.Log(tmp)
	cipher.encode(tmp)
	t.Log(tmp)
	cipher.decode(tmp)
	t.Log(tmp)
	if !reflect.DeepEqual(org, tmp) {
		t.Error("解码编码数据后无法还原数据，数据不对应")
	}
}

func BenchmarkEncode(b *testing.B) {
	password := RandPassword()
	cipher := NewCipher(password)
	bs := make([]byte, MB)
	b.ResetTimer()
	rand.Read(bs)
	cipher.encode(bs)
}

func BenchmarkDecode(b *testing.B) {
	password := RandPassword()
	cipher := NewCipher(password)
	bs := make([]byte, MB)
	b.ResetTimer()
	rand.Read(bs)
	cipher.decode(bs)
}
