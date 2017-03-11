package core

import (
	"net"
	"errors"
	"fmt"
	"io"
	"time"
)

const BUF_SIZE = 1024

type SecureSocket struct {
	Timeout    time.Duration
	Cipher     *Cipher
	LocalAddr  *net.TCPAddr
	ServerAddr *net.TCPAddr
}

func NewSecureSocket(timeout time.Duration, encodePassword *Password, localAddr, serverAddr *net.TCPAddr) *SecureSocket {
	return &SecureSocket{
		Timeout:    timeout,
		Cipher:     NewCipher(encodePassword),
		LocalAddr:  localAddr,
		ServerAddr: serverAddr,
	}
}

func (secureSocket *SecureSocket) DecodeRead(conn *net.TCPConn, bs []byte) (n int, err error) {
	n, err = conn.Read(bs)
	if err != nil {
		return
	}
	secureSocket.Cipher.decode(bs[:n])
	return
}

func (secureSocket *SecureSocket) EncodeWrite(conn *net.TCPConn, bs []byte) (int, error) {
	secureSocket.Cipher.encode(bs)
	return conn.Write(bs)
}

func (secureSocket *SecureSocket) EncodeCopy(dst *net.TCPConn, src *net.TCPConn) error {
	buf := make([]byte, BUF_SIZE)
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := secureSocket.EncodeWrite(dst, buf[0:nr])
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

func (secureSocket *SecureSocket) DecodeCopy(dst *net.TCPConn, src *net.TCPConn) error {
	buf := make([]byte, BUF_SIZE)
	for {
		nr, er := secureSocket.DecodeRead(src, buf)
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

func (secureSocket *SecureSocket) DialServer() (*net.TCPConn, error) {
	remoteConn, err := net.DialTCP("tcp", nil, secureSocket.ServerAddr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("dail remote %s fail:%s", secureSocket.ServerAddr, err))
	}
	return remoteConn, nil
}
