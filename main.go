package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

func main() {
	// If connecting to the default port, the URL can be simplified
// to just the hostname/IP.
// That is, the connect below is equivalent to:
// nats.Connect("nats://demo.nats.io:4222")
nc, err := nats.Connect("demo.nats.io")
if err != nil {
    log.Fatal(err)
}
defer nc.Close()

// Do something with the connection nc = Nats.connect("nats://demo.nats.io:4222");






}