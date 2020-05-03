package main

import (
	"gweibo/controller"
	"log"
	"math/rand"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	v := time.Duration(rand.Int63n(10))
	t := time.Minute + time.Second*v
	client := controller.NewClient("", t)
	client.Run()
}
