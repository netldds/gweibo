package main

import (
	"gweibo/controller"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	client := controller.NewClient("", 1*time.Second)
	client.Run()
}
