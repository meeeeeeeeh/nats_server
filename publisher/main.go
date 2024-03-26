package main

import (
	//"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	// "github.com/nats-io/stan.go"
	"io"
	"nats_server/config"
	"os"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("can't connect to NATS: %v", err)
	}
	defer nc.Close()

	file, err := os.Open(config.FilePath1)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	orderData, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	nc.Publish("intros", []byte(orderData))
	time.Sleep(1 * time.Second)

	// sc, err := stan.Connect("test-cluster", "publisher", stan.NatsURL("nats://localhost:4222"))
	// if err != nil {
	// 	panic(err)
	// }

	// file, err := os.Open(config.FilePath1)
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// orderData, err := io.ReadAll(file)
	// if err != nil {
	// 	panic(err)
	// }

	// err = nc.Publish("New order", []byte(orderData))
	// if err != nil {
	// 	panic(err)
	// }

	// //defer nc.Close()
}
