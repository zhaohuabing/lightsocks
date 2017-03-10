package cmd

import (
	"net"
	"io"
	"log"
	"reflect"
	"golang.org/x/net/proxy"
	"math/rand"
	"testing"
)

const (
	PACK_SIZE        = 1024 * 1024 * 2 //2Mb
	ECHO_SERVER_ADDR = "127.0.0.1:3453"
	SOCKS_PROXY_ADDR = "127.0.0.1:8010"
)

var (
	socksDialer proxy.Dialer
)

func init() {
	go runEchoServer()
	//初始化代理socksDialer
	var err error
	socksDialer, err = proxy.SOCKS5("tcp", SOCKS_PROXY_ADDR, nil, proxy.Direct)
	if err != nil {
		log.Fatalln(err)
	}
}

//启动echo server
func runEchoServer() {
	listener, err := net.Listen("tcp", ECHO_SERVER_ADDR)
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("listener.Accept", err)
			continue
		}
		log.Println("EchoServer", "listener.Accept")
		go func() {
			defer conn.Close()
			io.Copy(conn, conn)
			log.Println("EchoServer", "conn.Close")
		}()
	}
}

//获取 发送 data 到 echo server 并且收到全部返回 所花费到时间
func BenchmarkAll(b *testing.B) {
	//随机生产 PACK_SIZE byte的[]byte
	data := make([]byte, PACK_SIZE)
	n, err := rand.Read(data)
	if err != nil || n != PACK_SIZE {
		log.Fatalln("rand.Read", err)
	}
	buf := make([]byte, len(data))
	b.StartTimer()
	conn, err := socksDialer.Dial("tcp", ECHO_SERVER_ADDR)
	if err != nil {
		log.Fatalln(err)
	}
	go func() {
		conn.Write(data)
	}()
	_, err = io.ReadFull(conn, buf)
	conn.Close()
	b.StopTimer()
	if err != nil {
		log.Fatalln("io.ReadFull", err)
	}
	if !reflect.DeepEqual(data, buf) {
		log.Fatalln("response data length not equal to org")
	}
}
