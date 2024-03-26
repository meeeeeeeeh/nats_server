package main

import (
	"fmt"
	// "log"
	"time"

	//	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	//"nats_server/order_service"
)

func main() {
	// nc, err := nats.Connect(nats.DefaultURL)
	// if err != nil {
	// 	log.Fatalf("can't connect to NATS: %v", err)
	// }
	// defer nc.Close()

	// nc.Subscribe("intros", func(m *nats.Msg) {
	// 	log.Printf("I got a message: %s\n", string(m.Data))
	// })
	// //log.Println("Message recieved")

	// time.Sleep(1 * time.Hour)

	sc, err := stan.Connect("world-nats-stage", "endeharh", stan.NatsURL("nats://localhost:1234"))
	if err != nil {
		panic(err)
	}

	sub, err := sc.Subscribe("New order", MsgProcess)

	time.Sleep(1 * time.Hour)
	defer sub.Close()

}

func MsgProcess(m *stan.Msg) {
	fmt.Println(m)

}
