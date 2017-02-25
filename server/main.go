package main

import (
	"github.com/gwuhaolin/lightsocks/ss"
	"net"
	"io"
	"encoding/binary"
	"bufio"
	"os"
	"log"
)

var Config *ss.Config

// https://www.ietf.org/rfc/rfc1928.txt
func handleConn(localConn *ss.SecureConn) {
	defer localConn.Conn.Close()
	reader := bufio.NewReader(localConn)
	/**
	The localConn connects to the server, and sends a ver
   	identifier/method selection message:
		   +----+----------+----------+
                   |VER | NMETHODS | METHODS  |
                   +----+----------+----------+
                   | 1  |    1     | 1 to 255 |
                   +----+----------+----------+
	The VER field is set to X'05' for this ver of the protocol.  The
   	NMETHODS field contains the number of method identifier octets that
   	appear in the METHODS field.
	 */
	// 第一个字段VER代表Socks的版本，Socks5默认为0x05，其固定长度为1个字节
	ver, err := reader.ReadByte()
	// 只支持版本5
	if err != nil || ver != 0x05 {
		return
	}
	// 第二个字段METHODS表示第三个字段METHODS的长度，它的长度也是1个字节
	nMethods, err := reader.ReadByte()
	if err != nil {
		return
	}
	// 第三个METHODS表示客户端支持的验证方式，可以有多种，他的尝试是1-255个字节。
	_, err = reader.Discard(int(nMethods))
	if err != nil {
		return
	}
	/**
	The server selects from one of the methods given in METHODS, and
   	sends a METHOD selection message:

                         +----+--------+
                         |VER | METHOD |
                         +----+--------+
                         | 1  |   1    |
                         +----+--------+
	 */
	// 不需要验证，直接验证通过
	localConn.Write([]byte{0x05, 0x00})
	/**
	+----+-----+-------+------+----------+----------+
        |VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
        +----+-----+-------+------+----------+----------+
        | 1  |  1  | X'00' |  1   | Variable |    2     |
        +----+-----+-------+------+----------+----------+
	 */
	// VER代表Socks协议的版本，Socks5默认为0x05，其值长度为1个字节
	_, err = reader.Discard(1)
	if err != nil {
		return
	}
	// CMD代表客户端请求的类型，值长度也是1个字节，有三种类型
	cmd, err := reader.ReadByte()
	switch cmd {
	case 0x01:
	//	CONNECT X'01'
	case 0x02:
	//	BIND X'02'
	case 0x03:
	//	UDP ASSOCIATE X'03'
	default:
		return
	}
	// RSV保留字，值长度为1个字节
	_, err = reader.Discard(1)
	if err != nil {
		return
	}
	// aType 代表请求的远程服务器地址类型，值长度1个字节，有三种类型
	aType, err := reader.ReadByte()
	if err != nil {
		return
	}
	var dIP []byte
	switch aType {
	case 0x01:
		//	IP V4 address: X'01'
		ipv4 := make([]byte, net.IPv4len)
		_, err = reader.Read(ipv4)
		if err != nil {
			return
		}
		dIP = ipv4
	case 0x03:
		//	DOMAINNAME: X'03'
		domainLen, err := reader.ReadByte()
		if err != nil {
			return
		}
		domain := make([]byte, domainLen)
		_, err = reader.Read(domain)
		if err != nil {
			return
		}
		ipAddr, err := net.ResolveIPAddr("ip", string(domain))
		if err != nil {
			localConn.Write([]byte{0x05, 0x04, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}) //响应客户端Host unreachable
			return
		}
		dIP = ipAddr.IP
	case 0x04:
		//	IP V6 address: X'04'
		ipv6 := make([]byte, net.IPv6len)
		_, err = reader.Read(ipv6)
		if err != nil {
			return
		}
		dIP = ipv6
	default:
		return
	}
	dPort := make([]byte, 2)
	_, err = reader.Read(dPort)
	if err != nil {
		return
	}
	server, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   dIP,
		Port: int(binary.BigEndian.Uint16(dPort)),
	})
	/**
	 +----+-----+-------+------+----------+----------+
        |VER | REP |  RSV  | ATYP | BND.ADDR | BND.PORT |
        +----+-----+-------+------+----------+----------+
        | 1  |  1  | X'00' |  1   | Variable |    2     |
        +----+-----+-------+------+----------+----------+
	 */
	if err != nil {
		localConn.Write([]byte{0x05, 0x03, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}) //响应客户端Network unreachable
		return
	} else {
		localConn.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}) //响应客户端连接成功
		defer server.Close()
	}
	//进行转发
	go io.Copy(server, localConn)
	io.Copy(localConn, server)
}

func Run() {
	defer func() {
		log.Println(recover())
	}()
	for localConn := range ss.Listen(Config) {
		go handleConn(localConn)
	}
}

func main() {
	filePath := os.Args[1]
	var err error
	Config, err = ss.ParseConfig(filePath)
	if err != nil {
		panic(err)
	}
	Run()
}
