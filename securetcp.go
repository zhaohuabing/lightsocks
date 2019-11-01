package lightsocks

import (
	"io"
	"log"
	"net"
)

const (
	bufSize = 1024
)

// 加密传输的 TCP Socket
type SecureTCPConn struct {
	io.ReadWriteCloser
	EncodeCipher *cipher
	DecodeCipher *cipher
}

// 从输入流里读取加密过的数据，解密后把原数据放到bs里
func (secureSocket *SecureTCPConn) DecodeRead(bs []byte) (n int, err error) {
	n, err = secureSocket.Read(bs)
	if err != nil {
		return
	}
	secureSocket.DecodeCipher.decode(bs[:n])
	return
}

// 把放在bs里的数据加密后立即全部写入输出流
func (secureSocket *SecureTCPConn) EncodeWrite(bs []byte) (int, error) {
	secureSocket.EncodeCipher.encode(bs)
	return secureSocket.Write(bs)
}

// 从src中源源不断的读取原数据加密后写入到dst，直到src中没有数据可以再读取
func (secureSocket *SecureTCPConn) EncodeCopy(dst io.ReadWriteCloser) error {
	buf := make([]byte, bufSize)
	for {
		readCount, errRead := secureSocket.Read(buf)
		if errRead != nil {
			if errRead != io.EOF {
				return errRead
			} else {
				return nil
			}
		}
		if readCount > 0 {
			writeCount, errWrite := (&SecureTCPConn{
				ReadWriteCloser: dst,
				EncodeCipher:          secureSocket.EncodeCipher,
				DecodeCipher:          secureSocket.DecodeCipher,
			}).EncodeWrite(buf[0:readCount])
			if errWrite != nil {
				return errWrite
			}
			if readCount != writeCount {
				return io.ErrShortWrite
			}
		}
	}
}

// 从src中源源不断的读取加密后的数据解密后写入到dst，直到src中没有数据可以再读取
func (secureSocket *SecureTCPConn) DecodeCopy(dst io.Writer) error {
	buf := make([]byte, bufSize)
	for {
		readCount, errRead := secureSocket.DecodeRead(buf)
		if errRead != nil {
			if errRead != io.EOF {
				return errRead
			} else {
				return nil
			}
		}
		if readCount > 0 {
			writeCount, errWrite := dst.Write(buf[0:readCount])
			if errWrite != nil {
				return errWrite
			}
			if readCount != writeCount {
				return io.ErrShortWrite
			}
		}
	}
}

// see net.DialTCP
func DialTCPSecure(raddr *net.TCPAddr, password *password) (*SecureTCPConn, error) {
	remoteConn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		return nil, err
	}
	return &SecureTCPConn{
		ReadWriteCloser: remoteConn,
		EncodeCipher:     newCipher(password),
		DecodeCipher:     newCipher(password),
	}, nil
}

// see net.ListenTCP
func ListenSecureTCP(laddr *net.TCPAddr, password *password, handleConn func(localConn *SecureTCPConn), didListen func(listenAddr net.Addr)) error {
	log.Println("Try listen on : ", laddr)
	listener, err := net.ListenTCP("tcp", laddr)
	log.Println("Start listen on : ", laddr)
	if err != nil {
		return err
	}

	defer listener.Close()

	if didListen != nil {
		didListen(listener.Addr())
	}

	for {
		localConn, err := listener.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("Accept client connection : ", localConn.RemoteAddr())
		// localConn被关闭时直接清除所有数据 不管没有发送的数据
		localConn.SetLinger(0)
		go handleConn(&SecureTCPConn{
			ReadWriteCloser: localConn,
			EncodeCipher:     newCipher(password),
			DecodeCipher:     newCipher(password),
		})
	}
}
