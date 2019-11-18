package lightsocks

import (
	"crypto/rand"
	"reflect"
	"testing"
)

const (
	MB = 1024 * 1024
)

// 测试 cipher 加密解密
func TestCipher(t *testing.T) {
	password := "38K1EaMp+f+H+w51EB4Mn2peMPjKsxfkdl8k9vRi2feB3GY926hI4DzGS7aNwAOEBBnyXbIlTMOaBwn+wSrjfGzlu7B3NE8i2EA+9Zdx082bAlqIy5OPFDi/li69SVil2oAWMZS+olYtW0o51R1Btz9HCNK4/J0Kbv1r3lPUOqzEO6sAKNAYGrxlVaBhJ5DMpAYjgzI3jGfz6Vcbe+uYp/Er+q2clX0vpoLsbxNtH9eO6OYcZENoeaF+isUVNlSvqbGS4WAByJmRUYvOtO94JjPJ567RubpELMfPnnBZDTXwcmlF4kLdBWMhhRJORnOJC4btIA/q1npSTX+qUO5cdA=="
	t.Log(password)
	p, _ := parsePassword(password)
	encodeCipher := newCipher(p)
	decodeCipher := newCipher(p)
	// 原数据
	org := make([]byte, 10000)
	for i := 0; i < 10000; i++ {
		org[i] = byte(1)
	}
	/*for i := passwordLength; i < passwordLength*2; i++ {

		org[i] = byte(i)
	}*/
	// 复制一份原数据到 tmp
	tmp := make([]byte, 10000)
	copy(tmp, org)
	t.Log(tmp)
	// 加密 tmp
	encodeCipher.encode(tmp)
	t.Log(tmp)
	// 解密 tmp
	decodeCipher.decode(tmp)
	t.Log(tmp)
	if !reflect.DeepEqual(org, tmp) {
		t.Error("解码编码数据后无法还原数据，数据不对应")
	}
}

func BenchmarkEncode(b *testing.B) {
	password := RandPassword()
	p, _ := parsePassword(password)
	cipher := newCipher(p)
	bs := make([]byte, MB)
	b.ResetTimer()
	rand.Read(bs)
	cipher.encode(bs)
}

func BenchmarkDecode(b *testing.B) {
	password := RandPassword()
	p, _ := parsePassword(password)
	cipher := newCipher(p)
	bs := make([]byte, MB)
	b.ResetTimer()
	rand.Read(bs)
	cipher.decode(bs)
}
