package ss

import (
	"time"
	"net"
)

type Config struct {
	Timeout    time.Duration
	Cipher     *Cipher
	LocalAddr  *net.TCPAddr
	ServerAddr *net.TCPAddr
}

var GlobalConfig *Config

func NewConfig(timeout time.Duration, password *Password, localAddr *net.TCPAddr, serverAddr *net.TCPAddr) *Config {
	cipher := NewCipher(password)
	return &Config{
		Timeout:    timeout,
		Cipher:     cipher,
		LocalAddr:  localAddr,
		ServerAddr: serverAddr,
	}
}
