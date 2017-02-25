package ss

import (
	"os"
	"encoding/json"
	"time"
	"fmt"
	"strings"
	"errors"
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

func ParseConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("open file %s error:%s", filePath, err))
	}
	defer file.Close()
	config := &Config{
		Local:    ":8010",
		Remote:   ":8010",
		Password: RandPassword(),
		Timeout:  10 * time.Second,
	}
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("invalid json config file:\n%s", file))
	}
	fmt.Println(config)
	cipher, err := NewCipher(strings.TrimSpace(config.Password))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("invalid password:%s", config.Password))
	}
	config.Cipher = cipher
	return config, nil
}
