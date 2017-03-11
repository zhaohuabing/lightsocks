package core

import (
	"net"
	"errors"
	"fmt"
	"io"
)

const BUF_SIZE = 1024

func DecodeRead(conn *net.TCPConn, bs []byte) (n int, err error) {
	n, err = conn.Read(bs)
	if err != nil {
		return
	}
	GlobalConfig.Cipher.decode(bs[:n])
	return
}

func EncodeWrite(conn *net.TCPConn, bs []byte) (int, error) {
	GlobalConfig.Cipher.encode(bs)
	return conn.Write(bs)
}

func EncodeCopy(dst *net.TCPConn, src *net.TCPConn) error {
	buf := make([]byte, BUF_SIZE)
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := EncodeWrite(dst, buf[0:nr])
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
}

func DecodeCopy(dst *net.TCPConn, src *net.TCPConn) error {
	buf := make([]byte, BUF_SIZE)
	for {
		nr, er := DecodeRead(src, buf)
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
}

func DialServer() (*net.TCPConn, error) {
	remoteConn, err := net.DialTCP("tcp", nil, GlobalConfig.ServerAddr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("dail remote %s fail:%s", GlobalConfig.ServerAddr, err))
	}
	return remoteConn, nil
}
