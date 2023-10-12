package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func main() {

	nc, err := nats.Connect("nats://localhost:4333")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	if err := nc.Publish("foo", []byte("the first message!")); err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second)
	sb, err := nc.Subscribe("foo", func(msg *nats.Msg) {
		fmt.Printf("receive message -> %s \n", msg.Data)
	})

	fmt.Println(sb.Type())

	time.Sleep(time.Second)

}
