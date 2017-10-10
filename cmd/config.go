package cmd

import (
	"fmt"
	"os"
	"encoding/json"
	"log"
	"github.com/gwuhaolin/lightsocks/core"
	"github.com/mitchellh/go-homedir"
	"path"
	"io/ioutil"
)

type Config struct {
	ListenAddr string `json:"listen"`
	RemoteAddr string `json:"remote"`
	Password   string `json:"password"`
}

// 配置文件路径
var configPath string

func init() {
	home, _ := homedir.Dir()
	configPath = path.Join(home, ".lightsocksrc")
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

// 保存配置到配置文件
func (config *Config) SaveConfig() {
	configJson, _ := json.MarshalIndent(config, "", "	")
	err := ioutil.WriteFile(configPath, configJson, 0644)
	if err != nil {
		fmt.Errorf("save config error: %s", err)
	}
	log.Printf("save config successful to %s\n", configPath)
}

func ReadConfig() *Config {
	config := &Config{
		ListenAddr: ":7474",
		RemoteAddr: ":7474",
		Password:   core.RandPassword().String(),
	}

	if _, err := os.Stat(configPath); !os.IsNotExist(err) {
		// 如果配置文件存在，就采用配置文件中的配置
		log.Printf("use config form %s\n", configPath)
		file, err := os.Open(configPath)
		if err != nil {
			log.Fatalf("open config file %s error:%s", configPath, err)
		}
		defer file.Close()

		// parse & set Cipher
		err = json.NewDecoder(file).Decode(config)
		if err != nil {
			log.Fatalf("invalid config file:\n%s", file)
		}
	}
	config.SaveConfig()
	return config
}
