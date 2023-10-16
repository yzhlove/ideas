package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func main() {

	go func() {
		for i := 0; i < 10; i++ {
			go subscribe()
		}
	}()

	nc, err := nats.Connect("nats://127.0.0.1:4333")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	time.Sleep(time.Millisecond * 100)

	nc.Publish("hello", []byte("what are you doing ..."))

	time.Sleep(time.Second * 2)
}

func subscribe() {

	nc, err := nats.Connect("nats://127.0.0.1:4333")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	ch := make(chan struct{})

	nc.Subscribe("hello", func(msg *nats.Msg) {
		defer func() {
			ch <- struct{}{}
		}()
		fmt.Println("---------> ", string(msg.Data))
	})
	<-ch
}
