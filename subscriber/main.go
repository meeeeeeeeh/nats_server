package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"nats_server/order_service"
	"time"
	// "github.com/nats-io/stan.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("can't connect to NATS: %v", err)
	}
	defer nc.Close()

	nc.Subscribe("order", order_service.MsgProcess)
	time.Sleep(1 * time.Hour)

	defer nc.Close()

}
