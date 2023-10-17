package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

// 请求 - 回复 模型

func main() {

	go subscribe1()

	nc, err := nats.Connect("nats://127.0.0.1:4333")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	time.Sleep(time.Millisecond * 100)

	ret, err := nc.Request("hello", []byte("你瞅啥？"), time.Second)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ret.Data ==> ", string(ret.Data))

	time.Sleep(time.Second * 2)

}

func subscribe1() {
	nc, err := nats.Connect("nats://127.0.0.1:4333")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	stop := make(chan struct{})

	nc.Subscribe("hello", func(msg *nats.Msg) {
		fmt.Println("msg --> ", string(msg.Data))
		if err := nc.Publish(msg.Reply, []byte("瞅你咋滴！")); err != nil {
			log.Fatal(err)
		}
		//msg.Respond([]byte("瞅你咋滴！")) 同上
	})

	<-stop
}
