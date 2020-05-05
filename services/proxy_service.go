package services

import (
	"context"
	"gweibo/common"
	"log"
	"net"

	"golang.org/x/net/proxy"
)

type Agent struct {
	Addr string
}

var defaultAddr = "localhost:1080"

func NewProxyAgent() common.Socks5Proxy {
	return &Agent{}
}
func (s *Agent) SetAddr(addr string) {
	s.Addr = addr
}
func (s *Agent) GetDial() func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
	return func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
		if s.Addr == "" {
			s.Addr = defaultAddr
		}
		p, err := proxy.SOCKS5(network, addr, nil, proxy.Direct)
		if err != nil {
			log.Println(err)
		}
		return p.Dial(network, addr)
	}
}
