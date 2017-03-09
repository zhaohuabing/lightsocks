package core

import (
	"testing"
	"reflect"
)

func TestRandPassword(t *testing.T) {
	t.Log(RandPassword())
}

func TestNewCipher(t *testing.T) {
	password := RandPassword()
	t.Log(password)
	cipher := NewCipher(password)
	org := make([]byte, PASSWORD_LENGTH)
	for i := 0; i < PASSWORD_LENGTH; i++ {
		org[i] = byte(i)
	}
	tmp := make([]byte, PASSWORD_LENGTH)
	copy(tmp, org)
	t.Log(tmp)
	cipher.encode(tmp)
	t.Log(tmp)
	cipher.decode(tmp)
	t.Log(tmp)
	if !reflect.DeepEqual(org, tmp) {
		t.Error("encode decode error")
	}
}
