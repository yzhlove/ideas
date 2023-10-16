package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

// queue 模型，只能有一个人收到数据

func main() {

	go func() {
		for i := 0; i < 10; i++ {
			go subscribe("hello", "group1", i+1)
		}

		for i := 0; i < 10; i++ {
			go subscribe("hello", "group2", i+10+1)
		}

	}()

	nc, err := nats.Connect("nats://127.0.0.1:4333")
	if err != nil {
		log.Fatal(err)
	}

	defer nc.Close()

	time.Sleep(time.Millisecond * 100)

	nc.Publish("hello", []byte(" hello hello hello !!!"))

	time.Sleep(time.Second)

}

func subscribe(sub, queue string, id int) {
	nc, err := nats.Connect("nats://127.0.0.1:4333")
	if err != nil {
		log.Fatal(err)
	}

	defer nc.Close()

	s := make(chan struct{})

	nc.QueueSubscribe(sub, queue, func(msg *nats.Msg) {
		defer close(s)
		fmt.Println("id = ", id, " msg ==> ", string(msg.Data))
	})

	<-s

}
