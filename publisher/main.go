package main

import (
	//"fmt"
	// "log"
	// "time"

	"github.com/nats-io/stan.go"
	//"github.com/nats-io/nats.go"
	"io"
	"nats_server/config"
	"os"
)

func main() {
	// nc, err := nats.Connect(nats.DefaultURL)
	// if err != nil {
	// 	log.Fatalf("can't connect to NATS: %v", err)
	// }
	// defer nc.Close()

	// for {
	// 	nc.Publish("intros", []byte("Hello world"))
	// 	time.Sleep(1 * time.Second)
	// }

	sc, err := stan.Connect("test-nats", "route_user")
	if err != nil {
		panic(err)
	}

	file, err := os.Open(config.FilePath1)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	orderData, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = sc.Publish("New order", orderData)
	if err != nil {
		panic(err)
	}

	defer sc.Close()
}
