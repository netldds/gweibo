package gweibo

import (
	"gweibo/common"
	"gweibo/controller"
	"gweibo/services"
	"log"
	"net/http"
	"sync"
	"time"
)

type Client struct {
	common.GCleint
}

func init() {
	common.LoadConf()
}
func NewClient(elapseTime time.Duration, saver common.Store, proxy common.Socks5Proxy) *Client {
	//todo proxy setup
	return &Client{
		GCleint: common.GCleint{
			HttpClient: &http.Client{},
			Ticker:     time.NewTicker(elapseTime),
			ElapseTime: elapseTime,
			Saver:      saver,
			ProxyAgent: proxy,
		},
	}
}

//TODO web service

func (s *Client) GetTheOnePost(request common.RequestController) {
	s.Ticker = time.NewTicker(s.ElapseTime)
	for {
		select {
		case <-s.Ticker.C:
			err := request.Send(&s.GCleint)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func NewGetTheOnePostRequest() *controller.GetTheOnePostRequest {
	return &controller.GetTheOnePostRequest{
		Mu:             sync.Mutex{},
		RequestService: services.HomePageReq,
		LastInfo:       controller.MidInfo{},
	}
}
