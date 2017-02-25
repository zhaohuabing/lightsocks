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
		b[i] = conn.cipher.Decode(b[i])
	}
	return
}

func (conn *SecureConn) Write(b []byte) (int, error) {
	for i, v := range b {
		b[i] = conn.cipher.Encode(v)
	}
	return conn.Conn.Write(b)
}

func (conn *SecureConn) Close() error {
	return conn.Conn.Close()
}

func Dial(serverAddr string, password string) (*SecureConn, error) {
	server, err := net.Dial("tcp", serverAddr)
	if err != nil {
		return nil, err
	}
	cipher, err := NewCipher(password)
	return &SecureConn{
		Conn:   server,
		cipher: cipher,
	}, nil
}

func Listen(laddr string, password string) (chan *SecureConn) {
	ch := make(chan *SecureConn, 10)
	l, err := net.Listen("tcp", laddr)
	if err != nil {
		panic(err)
	}
	cipher, err := NewCipher(password)
	go func() {
		for {
			client, _ := l.Accept()
			ch <- &SecureConn{
				Conn:   client,
				cipher: cipher,
			}
		}
	}()
	return ch
}
