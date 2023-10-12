package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func main() {

	go subscribe()
	publish()

	time.Sleep(time.Second * 5)
}

func tryCatch(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func publish() {
	nc, err := nats.Connect("127.0.0.1:4333")
	tryCatch(err)
	defer nc.Close()

	tryCatch(nc.Publish("test1", []byte("this is first message!")))

}

func subscribe() {

	nc, err := nats.Connect("127.0.0.1:4333")
	tryCatch(err)
	defer nc.Close()

	for {
		sb, err := nc.Subscribe("test1", func(msg *nats.Msg) {
			fmt.Println("receive message: ", string(msg.Data))
		})
		tryCatch(err)
		sb = sb
		time.Sleep(time.Second)
	}
}
