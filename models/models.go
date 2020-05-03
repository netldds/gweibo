package models

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"time"
)

const (
	GET = "GET"
)
const (
	ROOT_URL = "https://weibo.com/p/"
	PID      = "1005056006394101"
)

type WeiReuqester interface {
	GetNextRequest() *http.Request
	Reset()
}
type WeiRequests struct {
	Req []WeiRequest
	seq int
}
type WeiRequest struct {
	Method  string
	Params  []string
	Timeout time.Duration
	Query   url.Values
}

func (s *WeiRequests) GetBody() io.Reader {
	return nil
}
func (s *WeiRequests) GetQuery() string {
	return s.Req[s.seq].Query.Encode()
}
func (s *WeiRequests) GetTimeOut() float64 {
	return s.Req[s.seq].Timeout.Seconds()
}
func (s *WeiRequests) GetMethmod() string {
	return s.Req[s.seq].Method
}
func (s *WeiRequests) GetUrl() string {
	param := ""
	for _, v := range s.Req[s.seq].Params {
		param = path.Join(param, v)
	}
	return fmt.Sprintf("%v%v?%v", ROOT_URL, param, s.GetQuery())
}
func (s *WeiRequests) GetNextRequest() *http.Request {
	req, err := http.NewRequest(s.GetMethmod(), s.GetUrl(), nil)
	if err != nil {
		log.Println(err)
	}
	req.AddCookie(&http.Cookie{Name: "Path", Value: "/"})
	req.AddCookie(&http.Cookie{Name: "YF-Page-G0", Value: "08eca8e8b3cf854de2e10f8127216863|1588529579|1588529547"})
	req.AddCookie(&http.Cookie{Name: "SUB", Value: "_2A25zq3ERDeRhGedM6VYU-S_LyjuIHXVQweXZrDV8PUNbmtANLU6ikW9NWcOv5pZmSc0mCA9j5wGZ-u28YzWbtq0j"})
	s.seq++
	if s.seq >= len(s.Req) {
		s.seq = 0
	}
	return req
}
func (s *WeiRequests) Reset() {
	s.seq = 0
}
