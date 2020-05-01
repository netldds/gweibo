package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gweibo/common"
	"gweibo/models"
	"gweibo/services"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/html"
)

type WeiClient struct {
	mu         sync.Mutex
	client     http.Client
	requests   models.WeiReuqester
	ProxyAgent services.Socks5Proxy
	Saver      models.Saver
	ElapseTime time.Duration
	ticker     *time.Ticker
}

func NewClient(proxyAddr string, timeInterval time.Duration) models.Clienter {
	s := &WeiClient{
		mu:         sync.Mutex{},
		ProxyAgent: services.NewProxyAgent(proxyAddr),
		ElapseTime: timeInterval,
		requests:   services.HomePageReq,
	}
	if proxyAddr != "" {
		ts := &http.Transport{DialContext: s.ProxyAgent.GetDial()}
		s.client.Transport = ts
	}
	if s.Saver == nil {
		s.Saver = &common.DefaultSaver{}
	}
	return s
}

func (s *WeiClient) Run() {
	s.ticker = time.NewTicker(s.ElapseTime)
	for {
		select {
		case <-s.ticker.C:
			s.Launch()
		}
	}
}
func (s *WeiClient) Launch() {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, v := range s.requests.GetReqs() {
		resp, err := s.client.Do(v)
		if err != nil {
			log.Println(err)
		}
		//1.
		raw, _ := ioutil.ReadAll(resp.Body)
		middleContent := raw[23 : len(raw)-11]
		mjson := make(map[string]interface{})
		err = json.Unmarshal([]byte(middleContent), &mjson)
		if err != nil {
			log.Println(err)
		}
		htmlStr := mjson["html"].(string)
		data := bytes.NewBufferString(htmlStr)
		body, err := html.Parse(data)
		if err != nil {
			fmt.Println(err)
		}
		var f, g func(*html.Node)
		//获取图片地址  //wx4.sinaimg.cn/orj360/006yududgy1g5iabhjai3j310c0rkgob.jpg
		g = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "img" {
				for _, a := range n.Attr {
					if a.Key == "src" {
						//mid,date
						//fmt.Println(a.Val)
						s.Saver.Save(a.Val)
						return
					}
				}
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				g(c)
			}
		}
		f = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "a" {
				for _, a := range n.Attr {
					if a.Key == "node-type" && a.Val == "feed_list_item_date" {
						//mid,date
						//html.Render(os.Stdout, n)
						//fmt.Println(n.Attr[0].Val)
						//g(n.Parent.Parent)
						s.Saver.Save(n.Attr[0].Val)
						break
					}
				}
			}
			//https://godoc.org/golang.org/x/net/html#Node
			//https://html.spec.whatwg.org/multipage/parsing.html#tree-construction
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
		f(body)
		resp.Body.Close()
	}
}
