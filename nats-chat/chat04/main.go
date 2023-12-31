package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"sync"
	"time"
)

func main() {

	// 多个客户端同时获取

	go subscribe()

	nc, err := nats.Connect("nats://127.0.0.1:4333")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	time.Sleep(time.Millisecond * 10)

	nc.Publish("hello", []byte("what are you doing ?"))

	time.Sleep(time.Second)
}

func subscribe() {

	nc, err := nats.Connect("nats://127.0.0.1:4333")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	var wg sync.WaitGroup
	wg.Add(1)

	nc.Subscribe("hello", func(msg *nats.Msg) {
		defer wg.Done()

		fmt.Println("1111")

		fmt.Println("msg ---> ", string(msg.Data))
	})

	fmt.Println("2222")

	wg.Wait()

	fmt.Println("3333")
}
