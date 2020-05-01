package services

import (
	"context"
	"log"
	"net"

	"golang.org/x/net/proxy"
)

type Socks5Proxy interface {
	GetDial() func(ctx context.Context, network, addr string) (conn net.Conn, e error)
}
type Agent struct {
	Addr string
}

var defaultAddr = "localhost:1080"

func NewProxyAgent(proxyAddr string) Socks5Proxy {
	return &Agent{Addr: proxyAddr}
}
func (s *Agent) Dial(ctx context.Context, network string, addr string) {

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
