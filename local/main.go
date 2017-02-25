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
	Local    string `json:"local"`
	Server   string `json:"server"`
	Password string `json:"password"`
}

var config *LocalConfig

func handleConn(userConn net.Conn) {
	defer userConn.Close()
	server, err := ss.Dial(config.Server, config.Password)
	if err != nil {
		log.Println(err)
		return
	}
	defer server.Close()
	go io.Copy(server, userConn)
	io.Copy(userConn, server)
}

func Run() {
	l, err := net.Listen("tcp", config.Local)
	if err != nil {
		panic(err)
	}
	for {
		userConn, _ := l.Accept()
		go handleConn(userConn)
	}
}

func main() {
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	config = &LocalConfig{}
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		panic(err)
	}
	if len(config.Password) == 0 {
		config.Password = ss.RandPassword()
		log.Println("Use password:", config.Password)
	}
	Run()
}
