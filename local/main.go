package main

import (
	"net"
	"github.com/gwuhaolin/lightsocks/ss"
	"io"
	"log"
	"os"
	"encoding/json"
)

type LocalConfig struct {
	cipher   *ss.Cipher `json:"-"`
	Local    string `json:"local"`
	Server   string `json:"server"`
	Password string `json:"password"`
}

var Config *LocalConfig

func handleConn(userConn net.Conn) {
	defer userConn.Close()
	server, err := ss.Dial(Config.Server, Config.cipher)
	if err != nil {
		log.Println(err)
		return
	}
	defer server.Close()
	go io.Copy(server, userConn)
	io.Copy(userConn, server)
}

func Run() {
	listener, err := net.Listen("tcp", Config.Local)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	for {
		userConn, _ := listener.Accept()
		go handleConn(userConn)
	}
}

func main() {
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	Config = &LocalConfig{}
	err = json.NewDecoder(file).Decode(Config)
	if err != nil {
		panic(err)
	}
	if len(Config.Password) == 0 {
		Config.Password = ss.RandPassword()
		log.Println("Use password:", Config.Password)
	}
	cipher, err := ss.NewCipher(Config.Password)
	if err != nil {
		panic(err)
	}
	Config.cipher = cipher
	Run()
}
