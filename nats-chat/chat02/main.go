package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func main() {

	go func() {
		ncc, err := nats.Connect("127.0.0.1:4333")
		try(err)
		defer ncc.Close()

		for {
			ncc.Subscribe("test", func(msg *nats.Msg) {
				fmt.Println("data => ", string(msg.Data))
			})
			time.Sleep(time.Millisecond * 500)
		}

	}()

	nc, err := nats.Connect("127.0.0.1:4333")
	try(err)
	defer nc.Close()

	try(nc.Publish("test", []byte("hello world!!!")))

	time.Sleep(time.Second)

}

func try(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
