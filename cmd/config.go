package cmd

import (
	"time"
	"fmt"
	"github.com/gwuhaolin/lightsocks/ss"
	"os"
	"encoding/json"
	"log"
	"net"
)

type Config struct {
	Local    string `json:"local"`
	Server   string `json:"remote"`
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
	`, config.Local, config.Server, config.Password, config.Timeout)
}

func (config *Config) ToSsConfig() (*ss.Config, error) {
	password, err := ss.ParsePassword(config.Password)
	if err != nil {
		return nil, err
	}

	localAddr, err := net.ResolveTCPAddr("tcp", config.Local)
	if err != nil {
		return nil, err
	}

	serverAddr, err := net.ResolveTCPAddr("tcp", config.Server)
	if err != nil {
		return nil, err
	}

	return ss.NewConfig(config.Timeout, password, localAddr, serverAddr), nil
}

func ReadConfig() *Config {
	if len(os.Args) != 2 {
		log.Fatalln(`require param json config file path, call like this:
		ls-exec ./path/to/json/config/file/path
		`)
	}
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalln("open file %s error:%s", filePath, err)
	}
	defer file.Close()

	config := &Config{
		Local:    ":8010",
		Server:   ":8010",
		Password: ss.RandPassword().String(),
		Timeout:  10 * time.Second,
	}

	//parse & set Cipher
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		log.Fatalln(fmt.Sprintf("invalid json config file:\n%s", file))
	}
	return config
}
