package ss

import (
	"os"
	"encoding/json"
	"time"
	"fmt"
	"strings"
)

type Config struct {
	Cipher   *Cipher `json:"-"`
	Local    string `json:"local"`
	Remote   string `json:"remote"`
	Password string `json:"password"`
	Timeout  time.Duration `json:"timeout"`
}

func (config *Config) String() string {
	return fmt.Sprintf(`
=== Use Config ===
Local
	%s
Remote
	%s
Password
	%s
Timeout
	%s
	`, config.Local, config.Remote, config.Password, config.Timeout)
}

func ParseConfig(filePath string) (config *Config, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()
	config = &Config{
		Local:    ":8010",
		Remote:   ":8010",
		Password: RandPassword(),
		Timeout:  10 * time.Second,
	}
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		return
	}
	fmt.Println(config)
	cipher, err := NewCipher(strings.TrimSpace(config.Password))
	if err != nil {
		return
	}
	config.Cipher = cipher
	return
}
