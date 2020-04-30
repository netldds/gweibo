package models

import (
	"fmt"
	"io"
	"net/url"
	"time"
)

type Reuqester interface {
	GetBody() io.Reader
	GetQuery() string
	GetTimeOut() int
	GetMethmod() string
	GetUrl() string
}

const (
	GET = "GET"
)
const (
	ROOT_URL = "https://weibo.com/p/"
	PID      = "1005056006394101"
)

type ClientRequest struct {
	method  string
	parames []string
	Timeout time.Duration
	Query   url.Values
}

var (
	HomePageReq = ClientRequest{
		method:  GET,
		parames: []string{"1005056006394101", "home"},
		Timeout: time.Second * 10,
		Query: url.Values{
			"pids":        []string{"Pl_Official_MyProfileFeed__20"},
			"is_all":      []string{"1"},
			"ajaxpagelet": []string{"1"},
		},
	}
)

func (s *ClientRequest) GetBody() io.Reader {
	return nil
}
func (s *ClientRequest) GetQuery() string {
	return s.Query.Encode()
}
func (s *ClientRequest) GetTimeOut() float64 {
	return s.Timeout.Seconds()
}
func (s *ClientRequest) GetMethmod() string {
	return s.method
}
func (s *ClientRequest) GetUrl() string {
	param := ""
	for _, v := range s.parames {
		param += v
	}
	return fmt.Sprintf("%v%v?%v", ROOT_URL, param, s.GetQuery())
}
