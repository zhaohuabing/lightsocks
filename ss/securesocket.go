package ss

import (
	"net"
	"errors"
	"fmt"
)

type SecureConn struct {
	Conn   net.Conn
	cipher *Cipher
}

func (conn *SecureConn) Read(bs []byte) (n int, err error) {
	n, err = conn.Conn.Read(bs)
	if err != nil {
		return
	}
	conn.cipher.decode(bs[:n])
	return
}

func (conn *SecureConn) Write(bs []byte) (int, error) {
	conn.cipher.encode(bs)
	return conn.Conn.Write(bs)
}

func Dial(config *Config) (*SecureConn, error) {
	remoteConn, err := net.DialTimeout("tcp", config.Remote, config.Timeout)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("dail remote %s fail:%s", config.Remote, err))
	}
	return &SecureConn{
		Conn:   remoteConn,
		cipher: config.Cipher,
	}, nil
}

func Listen(config *Config) (chan *SecureConn, error) {
	ch := make(chan *SecureConn)
	listener, err := net.Listen("tcp", config.Local)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("listen error:%s", err))
	}
	go func() {
		defer listener.Close()
		for {
			localConn, _ := listener.Accept()
			ch <- &SecureConn{
				Conn:   localConn,
				cipher: config.Cipher,
			}
		}
	}()
	return ch, nil
}
