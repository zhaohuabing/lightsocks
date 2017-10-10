package cmd

import (
	"fmt"
	"os"
	"encoding/json"
	"log"
	"github.com/gwuhaolin/lightsocks/core"
)

type Config struct {
	ListenAddr string `json:"listen"`
	RemoteAddr string `json:"remote"`
	Password   string `json:"password"`
}

func (config *Config) String() string {
	return fmt.Sprintf(`
=== Use Config ===
Listen Address
	%s
Remote Address
	%s
Password
	%s
	`, config.ListenAddr, config.RemoteAddr, config.Password)
}

func ReadConfig() *Config {
	config := &Config{
		ListenAddr: ":7474",
		RemoteAddr: ":7474",
		Password:   core.RandPassword().String(),
	}

	if len(os.Args) == 2 {
		filePath := os.Args[1]
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatalf("open file %s error:%s", filePath, err)
		}
		defer file.Close()

		//parse & set Cipher
		err = json.NewDecoder(file).Decode(config)
		if err != nil {
			log.Fatalf("invalid json config file:\n%s", file)
		}
	}
	return config
}
