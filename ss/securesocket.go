package ss

import (
	"net"
)

type SecureConn struct {
	Conn   net.Conn
	cipher *Cipher
}

func (conn *SecureConn) Read(b []byte) (n int, err error) {
	n, err = conn.Conn.Read(b)
	if err != nil {
		return
	}
	for i := 0; i < n; i++ {
		b[i] = conn.cipher.decode(b[i])
	}
	return
}

func (conn *SecureConn) Write(b []byte) (int, error) {
	for i, v := range b {
		b[i] = conn.cipher.encode(v)
	}
	return conn.Conn.Write(b)
}

func (conn *SecureConn) Close() error {
	return conn.Conn.Close()
}

func Dial(remoteAddr string, cipher *Cipher) (*SecureConn, error) {
	remoteConn, err := net.Dial("tcp", remoteAddr)
	if err != nil {
		return nil, err
	}
	return &SecureConn{
		Conn:   remoteConn,
		cipher: cipher,
	}, nil
}

func Listen(laddr string, cipher *Cipher) (chan *SecureConn) {
	ch := make(chan *SecureConn)
	l, err := net.Listen("tcp", laddr)
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			localConn, _ := l.Accept()
			ch <- &SecureConn{
				Conn:   localConn,
				cipher: cipher,
			}
		}
	}()
	return ch
}
