package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	//"github.com/nats-io/stan.go"
	//"nats_server/order_service"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("can't connect to NATS: %v", err)
	}
	defer nc.Close()

	nc.Subscribe("intros", func(m *nats.Msg) {
		log.Printf("I got a message: %s\n", string(m.Data))
	})

	time.Sleep(1 * time.Hour)

	// sc, err := stan.Connect("world-nats-stage", "endeharh", stan.NatsURL("nats://localhost:4222"))
	// if err != nil {
	// 	panic(err)
	// }

	// sub, err := sc.Subscribe("New order", MsgProcess)

	//time.Sleep(1 * time.Hour)
	defer nc.Close()

}

func MsgProcess(m *nats.Msg) {
	fmt.Println(m)

}
