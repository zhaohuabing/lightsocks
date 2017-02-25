package ss

import "testing"

func TestRandPassword(t *testing.T) {
	t.Log(RandPassword())
}

func TestNewCipher(t *testing.T) {
	password := RandPassword()
	cipher, err := NewCipher(password)
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < PASSWORD_LENGTH; i++ {
		org := byte(i)
		e := cipher.encode(org)
		d := cipher.decode(e)
		if d != org {
			t.Error("Decode Encode error:", org, e, d)
		}
	}
}
