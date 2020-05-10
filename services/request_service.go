package services

import (
	"fmt"
	"gweibo/common"
	"io"
	"net/url"
	"path"
	"time"
)

const (
	HOST   = "weibo.com"
	PID    = "1005056006394101"
	SCHEME = "https"
)

//ROOT https://weibo.com/p/
var (
	HomePageReq = &WeiRequests{
		Req: []WeiRequest{
			{
				Method:  common.GET,
				Params:  []string{"p", PID, "home"},
				Timeout: time.Second * 10,
				Query: url.Values{
					//"stat_date":   []string{"201907"},
					"pids":        []string{"Pl_Official_MyProfileFeed__20"},
					"ajaxpagelet": []string{"1"},
				},
			},
			{
				Method:  common.GET,
				Params:  []string{"p", "aj", "mblog", "getlongtext"},
				Timeout: time.Second * 10,
				Query: url.Values{
					"mid": []string{""},
				},
			},
		},
	}
)

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

func (s *WeiRequests) NextRequest() {
	s.seq++
	if s.seq >= len(s.Req) {
		s.seq = 0
	}
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
func (s *WeiRequests) GetMethod() string {
	return s.Req[s.seq].Method
}
func (s *WeiRequests) GetRoot() string {
	return fmt.Sprintf("%v://%v", SCHEME, HOST)
}
func (s *WeiRequests) GetPath() string {
	param := ""
	for _, v := range s.Req[s.seq].Params {
		param = path.Join(param, v)
	}
	return param
}
func (s *WeiRequests) Reset() {
	s.seq = 0
}
