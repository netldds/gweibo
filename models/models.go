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

type Clienter interface {
	Run()
}
type Saver interface {
	Save(ct string)
}
type WeiReuqester interface {
	GetReqs() []*http.Request
}
type WeiRequest struct {
	Method  string
	Params  []string
	Timeout time.Duration
	Query   url.Values
}

func (s *WeiRequest) GetBody() io.Reader {
	return nil
}
func (s *WeiRequest) GetQuery() string {
	return s.Query.Encode()
}
func (s *WeiRequest) GetTimeOut() float64 {
	return s.Timeout.Seconds()
}
func (s *WeiRequest) GetMethmod() string {
	return s.Method
}
func (s *WeiRequest) GetUrl() string {
	param := ""
	for _, v := range s.Params {
		param = path.Join(param, v)
	}
	return fmt.Sprintf("%v%v?%v", ROOT_URL, param, s.GetQuery())
}
func (s *WeiRequest) GetReqs() []*http.Request {
	req, err := http.NewRequest(s.GetMethmod(), s.GetUrl(), nil)
	if err != nil {
		log.Println(err)
	}
	e := &http.Cookie{Name: "SUB", Value: "_2AkMp8PS_f8NxqwJRmf0Qym3hZYtxywzEieKfrAVkJRMxHRl-yj9kqlMOtRB6AnDaUG_s1jiCP8TQ5l46n-oHZaanbsDs"}
	req.AddCookie(e)
	return []*http.Request{req}
}
