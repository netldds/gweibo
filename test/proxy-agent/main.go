package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	"golang.org/x/net/proxy"
)

const (
	PROXY_ADDR = "127.0.0.1:9050"
	URL        = "http://skunksworkedp2cg.onion/sites.html"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime)
	client := http.Client{}
	ts := &http.Transport{}
	client.Transport = ts
	p, err := proxy.SOCKS5("tcp", "localhost:1080", nil, proxy.Direct)
	if err != nil {
		log.Println(err)
	}
	//Dial(network string, addr string) (c net.Conn, err error)
	//func(ctx context.Context, network string, addr string) (net.Conn, error)
	ts.DialContext = func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
		return p.Dial(network, addr)
	}
	req, err := http.NewRequest("GET", "https://baidu.com", nil)
	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(resp)
		log.Println(err)
	}
	io.Copy(os.Stdout, resp.Body)
	defer resp.Body.Close()
}
func A() {
	// create a socks5 dialer
	dialer, err := proxy.SOCKS5("tcp", PROXY_ADDR, nil, proxy.Direct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
		os.Exit(1)
	}
	// setup a http client
	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}
	// set our socks5 as the dialer
	httpTransport.Dial = dialer.Dial
	// create a request
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't create request:", err)
		os.Exit(2)
	}
	// use the http client to fetch the page
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't GET page:", err)
		os.Exit(3)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading body:", err)
		os.Exit(4)
	}
	fmt.Println(string(b))
}
