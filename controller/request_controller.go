package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gweibo/models"
	"gweibo/services"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/html"
)

type Client interface {
	Run()
}
type WeiClient struct {
	mu         sync.Mutex
	client     http.Client
	requests   models.WeiReuqester
	ProxyAgent services.Socks5Proxy
	Saver      services.Store
	ElapseTime time.Duration
	ticker     *time.Ticker
	postInfo   midInfo
}
type midInfo struct {
	Mid    string
	ImgUrl string
	t      time.Time
}

func NewClient(proxyAddr string, timeInterval time.Duration) Client {
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
		s.Saver = &services.StandardOutput{}
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

type RespBody struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		HTML string `json:"html"`
	} `json:"data"`
}

/*
	1.去掉
	<script>parent.FM.view(		data	)</script>
	2.中间data   -->   json
	3.条件 a标签 node-type="feed_list_item_date"  找到name 是条记录id  并且 标签内容是时间   4月15日 15:17
	4.通过 https://weibo.com/p/aj/mblog/getlongtext?mid=4498297848918260  获取json {"code":"100000","msg":"","data":{"html":" \u90a3\u9ad8\u7ea7\u7684\u662f\u600e\u4e48\u4f20\u8f93\u4fe1\u606f\u7684\uff1f\u91cf\u5b50\u7ea0\u7f20\uff1f<br>\u7b54\uff1a\u4f20\u8f93\u8fd9\u4e2a\u6982\u5ff5\u7684\u524d\u63d0\u662f\u5f97\u6709\u8ddd\u79bb\u969c\u788d\uff0c\u5982\u679c\u514b\u670d\u4e86\u8ddd\u79bb\u969c\u788d\uff0c\u90a3\u4f20\u8f93\u7684\u6982\u5ff5\u4e5f\u4e0d\u5b58\u5728\u4e86\uff0c\u7b11\u3002\u8fd1\u524d\u7684\uff0c\u75ab\u60c5\u671f\u5730\u7403\u5168\u4f53\u4eba\u7c7b\u7fa4\u4f53\u5669\u68a6\uff0c\u68a6\u4e2d\u88ab\u793a\u8b66\u88ab\u89e3\u6bd2\u600e\u4e48\u89e3\u91ca\uff0c\u4fe1\u606f\u65f6\u4ee3\uff0c\u8d8a\u6765\u8d8a\u591a\u7684\u795e\u8ff9\u5f81\u5146\u4f1a\u88ab\u6709\u4fe1\u4ef0\u7684\u4eba\u7fa4\u53d1\u73b0\u6c47\u603b\u5e76\u8ba8\u8bba\uff0c\u7b11\u3002 \u200b\u200b\u200b\u200b"}}
	5. unicode解码 不包括 图片
*/
func (s *WeiClient) Launch() {
	s.mu.Lock()
	defer s.mu.Unlock()
	v := s.requests.GetNextRequest()
	resp, err := s.client.Do(v)
	if err != nil {
		log.Println(err)
	}
	raw, _ := ioutil.ReadAll(resp.Body)
	newMid := s.Parse(string(raw))
	if newMid.Mid == s.postInfo.Mid {
		s.requests.Reset()
		return
	}
	s.postInfo = newMid
	if err != nil {
		log.Fatal(err)
	}
	raw = nil
	//get long text
	v = s.requests.GetNextRequest()
	v.URL.RawQuery = fmt.Sprintf("%v%v", v.URL.RawQuery, s.postInfo.Mid)
	resp, err = s.client.Do(v)
	if err != nil {
		log.Println(err)
	}
	raw, _ = ioutil.ReadAll(resp.Body)
	var body RespBody
	err = json.Unmarshal(raw, &body)
	if err != nil {
		log.Println(err)
	}
	s.Saver.SaveContext(s.postInfo.t, []byte(body.Data.HTML), s.postInfo.ImgUrl)
	resp.Body.Close()
}
func (s *WeiClient) Parse(raw string) midInfo {
	middleContent := raw[23 : len(raw)-11]
	mjson := make(map[string]interface{})
	err := json.Unmarshal([]byte(middleContent), &mjson)
	if err != nil {
		log.Fatal(err)
	}
	htmlStr := mjson["html"].(string)
	data := bytes.NewBufferString(htmlStr)
	body, err := html.Parse(data)
	if err != nil {
		log.Fatal(err)
	}
	return FindIds(body)
}

func FindIds(n *html.Node) midInfo {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "node-type" && a.Val == "feed_list_item_date" {
				//mid,date
				//html.Render(os.Stdout, n)
				//fmt.Println(n.Attr[0].Val)
				imgUrl := FindImg(n.Parent.Parent)
				//get long text
				//s.SaveContext(time.Now(), []byte(n.Attr[0].Val), imgUrl)
				//s.Save(n.Attr[0].Val)
				var t time.Time
				var err error
				for _, v := range n.Attr {
					if v.Key == "title" {
						t, err = time.Parse("2006-01-02 15:4", v.Val)
						if err != nil {
							log.Println(err)
						}
					}
				}
				return midInfo{
					Mid:    n.Attr[0].Val,
					ImgUrl: imgUrl,
					t:      t,
				}
			}
		}
	}
	//https://godoc.org/golang.org/x/net/html#Node
	//https://html.spec.whatwg.org/multipage/parsing.html#tree-construction
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		v := FindIds(c)
		if v.Mid != "" {
			return v
		}
	}
	return midInfo{}
}

//find img url
//获取图片地址  //wx4.sinaimg.cn/orj360/006yududgy1g5iabhjai3j310c0rkgob.jpg
func FindImg(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "img" {
		for _, a := range n.Attr {
			if a.Key == "src" {
				//mid,date
				//fmt.Println(a.Val)
				//s.Save(a.Val)
				return a.Val
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		FindImg(c)
	}
	return ""
}
