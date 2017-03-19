package cmd

import (
	"fmt"
	"os"
	"encoding/json"
	"log"
	"github.com/gwuhaolin/lightsocks/core"
)

type Config struct {
	Local    string `json:"local"`
	Server   string `json:"remote"`
	Password string `json:"password"`
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
	`, config.Local, config.Server, config.Password)
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
		Password: core.RandPassword().String(),
	}
	//parse & set Cipher
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		log.Fatalln(fmt.Sprintf("invalid json config file:\n%s", file))
	}
	return config
}
