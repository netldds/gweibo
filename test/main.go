package main

import (
	"gweibo"
	"gweibo/services"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	req := gweibo.NewGetTheOnePostRequest()
	var client *gweibo.Client
	if _, ok := os.LookupEnv("debug"); ok {
		client = gweibo.NewClient(time.Second, services.DefaultStore, services.NewProxyAgent())
	} else {
		v := time.Duration(rand.Int63n(10))
		t := time.Minute + time.Second*v
		client = gweibo.NewClient(t, services.DefaultStore, services.NewProxyAgent())
	}
	client.GetTheOnePost(req)
}
