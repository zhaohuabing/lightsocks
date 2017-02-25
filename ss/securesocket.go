package ss

import (
	"net"
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

func (conn *SecureConn) Close() error {
	return conn.Conn.Close()
}

func Dial(config *Config) (*SecureConn, error) {
	remoteConn, err := net.DialTimeout("tcp", config.Remote, config.Timeout)
	if err != nil {
		return nil, err
	}
	return &SecureConn{
		Conn:   remoteConn,
		cipher: config.Cipher,
	}, nil
}

func Listen(config *Config) (chan *SecureConn) {
	ch := make(chan *SecureConn)
	l, err := net.Listen("tcp", config.Local)
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			localConn, _ := l.Accept()
			ch <- &SecureConn{
				Conn:   localConn,
				cipher: config.Cipher,
			}
		}
	}()
	return ch
}
