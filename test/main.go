package main

import (
	"gweibo"
	"gweibo/services"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	//v := time.Duration(rand.Int63n(10))
	//t := time.Minute + time.Second*v
	req := gweibo.NewGetTheOnePostRequest()
	client := gweibo.NewClient(time.Second, services.DefaultStore, services.NewProxyAgent())
	client.GetTheOnePost(req)
}
