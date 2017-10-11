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
	configPath = path.Join(home, ".lightsocks.json")
}

// 保存配置到配置文件
func (config *Config) SaveConfig() {
	configJson, _ := json.MarshalIndent(config, "", "	")
	err := ioutil.WriteFile(configPath, configJson, 0644)
	if err != nil {
		fmt.Errorf("保存配置到文件 %s 出错: %s", configPath, err)
	}
	log.Printf("保存配置到文件 %s 成功\n", configPath)
}

func ReadConfig() *Config {
	config := &Config{
		ListenAddr: ":7448",
		RemoteAddr: ":7448",
		Password:   core.RandPassword().String(),
	}

	// 如果配置文件存在，就采用配置文件中的配置
	if _, err := os.Stat(configPath); !os.IsNotExist(err) {
		log.Printf("从文件 %s 中读取配置\n", configPath)
		file, err := os.Open(configPath)
		if err != nil {
			log.Fatalf("打开文件 %s 出错:%s", configPath, err)
		}
		defer file.Close()

		// parse & set Cipher
		err = json.NewDecoder(file).Decode(config)
		if err != nil {
			log.Fatalf("格式不合法的 JSON 配置文件:\n%s", file)
		}
	}
	config.SaveConfig()
	return config
}
