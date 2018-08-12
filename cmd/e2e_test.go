package cmd

import (
	"github.com/gwuhaolin/lightsocks"
	"golang.org/x/net/proxy"
	"io"
	"log"
	"math/rand"
	"net"
	"reflect"
	"sync"
	"testing"
	"time"
)

const (
	MaxPackSize               = 1024 * 1024 * 5 // 5Mb
	EchoServerAddr            = "127.0.0.1:3453"
	LightSocksProxyLocalAddr  = "127.0.0.1:8448"
	LightSocksProxyServerAddr = "127.0.0.1:8449"
)

var (
	lightsocksDialer proxy.Dialer
)

func init() {
	log.SetFlags(log.Lshortfile)
	go runEchoServer()
	go runLightsocksProxyServer()
	// 初始化代理socksDialer
	var err error
	// 等它们启动好
	time.Sleep(time.Second)
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
		log.Println("echoServer connect Accept")
		go func() {
			defer func() {
				conn.Close()
				log.Println("echoServer connect Close")
			}()
			io.Copy(conn, conn)
		}()
	}
}

func runLightsocksProxyServer() {
	password := lightsocks.RandPassword()
	serverS, err := lightsocks.NewLsLocal(password, LightSocksProxyLocalAddr, LightSocksProxyServerAddr)
	if err != nil {
		log.Fatalln(err)
	}
	localS, err := lightsocks.NewLsServer(password, LightSocksProxyServerAddr)
	if err != nil {
		log.Fatalln(err)
	}
	go func() {
		log.Fatalln(serverS.Listen(func(listenAddr net.Addr) {
			log.Println(listenAddr)
		}))
	}()
	log.Fatalln(localS.Listen(func(listenAddr net.Addr) {
		log.Println(listenAddr)
	}))
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

func TestLightsocks(t *testing.T) {
	testConnect(rand.Intn(MaxPackSize))
}

// 获取并发发送 data 到 echo server 并且收到全部返回 所花费到时间
func benchmarkLightsocks(concurrenceCount int) {
	wg := sync.WaitGroup{}
	wg.Add(concurrenceCount)
	for i := 0; i < concurrenceCount; i++ {
		go func() {
			testConnect(rand.Intn(MaxPackSize))
			wg.Done()
		}()
	}
	wg.Wait()
}

// 获取 发送 data 到 echo server 并且收到全部返回 所花费到时间
func BenchmarkLightsocks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		benchmarkLightsocks(10)
		b.StopTimer()
	}
}
