package ss

import (
	"net"
	"errors"
	"fmt"
	"io"
)

const BUF_SIZE = 1024

type SecureConn struct {
	Conn *net.TCPConn
}

func (conn *SecureConn) Read(bs []byte) (n int, err error) {
	n, err = conn.Conn.Read(bs)
	if err != nil {
		return
	}
	GlobalConfig.Cipher.decode(bs[:n])
	return
}

func (conn *SecureConn) Write(bs []byte) (int, error) {
	GlobalConfig.Cipher.encode(bs)
	return conn.Conn.Write(bs)
}

func Copy(dst io.Writer, src io.Reader) error {
	buf := make([]byte, BUF_SIZE)
	return CopyBuf(dst, src, buf)
}

func CopyBuf(dst io.Writer, src io.Reader, buf []byte) error {
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if ew != nil {
				return ew
			}
			if nr != nw {
				return io.ErrShortWrite
			}
		}
		if er != nil {
			if er != io.EOF {
				return er
			} else {
				return nil
			}
		}
	}
	buf = nil
	return nil
}

func DialServer() (*SecureConn, error) {
	remoteConn, err := net.DialTCP("tcp", nil, GlobalConfig.ServerAddr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("dail remote %s fail:%s", GlobalConfig.ServerAddr, err))
	}
	return &SecureConn{Conn: remoteConn }, nil
}

func ServerListen() (chan *SecureConn, error) {
	ch := make(chan *SecureConn)
	listener, err := net.ListenTCP("tcp", GlobalConfig.LocalAddr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("listen error:%s", err))
	}
	go func() {
		defer listener.Close()
		for {
			localConn, _ := listener.AcceptTCP()
			ch <- &SecureConn{Conn: localConn}
		}
	}()
	return ch, nil
}
