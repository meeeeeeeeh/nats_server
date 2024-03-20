package main

import (
	//"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("can't connect to NATS: %v", err)
	}
	defer nc.Close()

	for {
		nc.Publish("intros", []byte("Hello world"))
		time.Sleep(1 * time.Second)
	}

}
