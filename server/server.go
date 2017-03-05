package server

import (
	"net"
	"encoding/binary"
	"log"
	"github.com/gwuhaolin/lightsocks/ss"
)

func Run() {
	listener, err := net.ListenTCP("tcp", ss.GlobalConfig.LocalAddr)
	if err != nil {
		log.Fatalln("listen error:%s", err)
	}
	defer func() {
		log.Println(recover())
	}()
	for {
		localConn, _ := listener.AcceptTCP()
		go handleConn(localConn)
	}
}

// socks5实现
// https://www.ietf.org/rfc/rfc1928.txt
// http://www.jianshu.com/p/172810a70fad
func handleConn(localConn *net.TCPConn) {
	defer localConn.Close()
	buf := make([]byte, ss.BUF_SIZE)
	/**
	The localConn connects to the dstServer, and sends a ver
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
	_, err := ss.DecodeRead(localConn, buf)
	// 只支持版本5
	if err != nil || buf[0] != 0x05 {
		return
	}
	/**
	The dstServer selects from one of the methods given in METHODS, and
   	sends a METHOD selection message:

                         +----+--------+
                         |VER | METHOD |
                         +----+--------+
                         | 1  |   1    |
                         +----+--------+
	 */
	// 不需要验证，直接验证通过
	ss.EncodeWrite(localConn, []byte{0x05, 0x00})
	/**
	+----+-----+-------+------+----------+----------+
        |VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
        +----+-----+-------+------+----------+----------+
        | 1  |  1  | X'00' |  1   | Variable |    2     |
        +----+-----+-------+------+----------+----------+
	 */
	// VER代表Socks协议的版本，Socks5默认为0x05，其值长度为1个字节
	n, err := ss.DecodeRead(localConn, buf)
	// 最短域名= 3 bytes
	// 9 = 1+1+1+1+3+2
	if err != nil || n < 9 {
		return
	}
	var dIP []byte
	// aType 代表请求的远程服务器地址类型，值长度1个字节，有三种类型
	switch buf[3] {
	case 0x01:
		//	IP V4 address: X'01'
		dIP = buf[4:4+net.IPv4len]
	case 0x03:
		//	DOMAINNAME: X'03'
		ipAddr, err := net.ResolveIPAddr("ip", string(buf[5:n-2]))
		if err != nil {
			return
		}
		dIP = ipAddr.IP
	case 0x04:
		//	IP V6 address: X'04'
		dIP = buf[4:4+net.IPv6len]
	default:
		return
	}
	dPort := buf[n-2:]
	dstAddr := &net.TCPAddr{
		IP:   dIP,
		Port: int(binary.BigEndian.Uint16(dPort)),
	}
	dstServer, err := net.DialTCP("tcp", nil, dstAddr)
	/**
	 +----+-----+-------+------+----------+----------+
        |VER | REP |  RSV  | ATYP | BND.ADDR | BND.PORT |
        +----+-----+-------+------+----------+----------+
        | 1  |  1  | X'00' |  1   | Variable |    2     |
        +----+-----+-------+------+----------+----------+
	 */
	if err != nil {
		return
	} else {
		ss.EncodeWrite(localConn, []byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}) //响应客户端连接成功
		defer dstServer.Close()
	}
	//进行转发
	go ss.DecodeCopyBuf(dstServer, localConn, buf)
	ss.EncodeCopy(localConn, dstServer)
}
