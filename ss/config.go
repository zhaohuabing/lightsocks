package ss

import (
	"os"
	"encoding/json"
	"log"
)

type Config struct {
	Cipher   *Cipher `json:"-"`
	Local    string `json:"local"`
	Server   string `json:"server"`
	Password string `json:"password"`
}

func ParseConfig(filePath string) (config *Config, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()
	config = &Config{}
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		return
	}
	if len(config.Password) == 0 {
		config.Password = RandPassword()
		log.Println("Use password:", config.Password)
	}
	cipher, err := NewCipher(config.Password)
	if err != nil {
		return
	}
	config.Cipher = cipher
	return
}
