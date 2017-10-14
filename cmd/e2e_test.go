package cmd

import (
	"net"
	"io"
	"log"
	"golang.org/x/net/proxy"
	"math/rand"
	"testing"
	"reflect"
	"github.com/gwuhaolin/lightsocks/core"
	"github.com/gwuhaolin/lightsocks/server"
	"github.com/gwuhaolin/lightsocks/local"
)

const (
	MaxPackSize               = 1024 * 1024 * 2 // 2Mb
	EchoServerAddr            = "127.0.0.1:3453"
	LightSocksProxyLocalAddr  = "127.0.0.1:7441"
	LightSocksProxyServerAddr = "127.0.0.1:7442"
)

var (
	lightsocksDialer proxy.Dialer
)

func init() {
	go runEchoServer()
	go runLightsocksProxyServer()
	// 初始化代理socksDialer
	var err error
	lightsocksDialer, err = proxy.SOCKS5("tcp", LightSocksProxyLocalAddr, nil, proxy.Direct)
	if err != nil {
		log.Fatalln(err)
	}
}

// 启动echo server
func runEchoServer() {
	listener, err := net.Listen("tcp", EchoServerAddr)
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go func() {
			defer conn.Close()
			io.Copy(conn, conn)
		}()
	}
}

func runLightsocksProxyServer() {
	password := core.RandPassword()
	localAddr, _ := net.ResolveTCPAddr("tcp", LightSocksProxyLocalAddr)
	serverAddr, _ := net.ResolveTCPAddr("tcp", LightSocksProxyServerAddr)
	serverS := local.New(password, localAddr, serverAddr)
	localS := server.New(password, serverAddr)
	go serverS.Listen(nil)
	localS.Listen(nil)
}

// 发生一次连接测试经过代理后的数据传输的正确性
// packSize 代表这个连接发生数据的大小
func testConnect(packSize int) {
	// 随机生产 MaxPackSize byte的[]byte
	data := make([]byte, packSize)
	_, err := rand.Read(data)

	// 连接
	conn, err := lightsocksDialer.Dial("tcp", EchoServerAddr)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// 写
	go func() {
		conn.Write(data)
	}()

	// 读
	buf := make([]byte, len(data))
	_, err = io.ReadFull(conn, buf)
	if err != nil {
		log.Fatalln(err)
	}
	if !reflect.DeepEqual(data, buf) {
		log.Fatalln("通过 Lightsocks 代理传输得到的数据前后不一致")
	} else {
		log.Println("数据一致性验证通过")
	}
}

// 获取 发送 data 到 echo server 并且收到全部返回 所花费到时间
func BenchmarkLightsocks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		testConnect(rand.Intn(MaxPackSize))
		b.StopTimer()
	}
}
