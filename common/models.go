package common

import (
	"context"
	"net"
	"net/http"
	"time"
)

const (
	GET = "GET"
)

//TODO
/*
暴露请求构造中的数据，可以设置
cookie的获取和管理，先支持热更新，放配置文件
*/
type ClinterHandler interface {
	SetCookie(ck string)
}

func (s *GCleint) SetCookie(ck string) {
	//todo
}

type Response interface {
	ParseHttpResponse(body []byte) error
}
type GCleint struct {
	HttpClient *http.Client
	Ticker     *time.Ticker
	ElapseTime time.Duration
	Saver      Store
	ProxyAgent Socks5Proxy
}
type Socks5Proxy interface {
	SetAddr(addr string)
	GetDial() func(ctx context.Context, network, addr string) (conn net.Conn, e error)
}
type Store interface {
	SaveContext(timestamp time.Time, context []byte, img string)
}
type RequestService interface {
	//GetQuery() string
	Reset()
	GetRoot() string
	GetMethod() string
	GetPath() string
	NextRequest()
}
type RequestController interface {
	Send(client *GCleint) error
	RequestService
}
type ServiceConfig struct {
	Cookie map[string]interface{} `yaml:"cookie"`
	Proxy  string                 `yaml:"proxy"`
}

func (s *ServiceConfig) Isillegal() bool {
	return true
}
