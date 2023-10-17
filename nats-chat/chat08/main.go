package main

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"log"
	"time"
)

func main() {

	nc, err := nats.Connect("nats://127.0.0.1:4333")
	if err != nil {
		log.Fatal(err)
	}

	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	stream, err := js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     "orders",
		Subjects: []string{"orders.*"},
	})

	for i := 0; i < 100; i++ {
		_, err := js.Publish(ctx, "orders.new", []byte(fmt.Sprintf("hello world: %d", i+1)))
		if err != nil {
			log.Fatal(err)
		}
	}

	cons, err := stream.CreateConsumer(ctx, jetstream.ConsumerConfig{
		Durable:   "cons",
		AckPolicy: jetstream.AckExplicitPolicy,
	})
	if err != nil {
		log.Fatal(err)
	}

	var counter int
	msgs, err := cons.Fetch(10)
	if err != nil {
		log.Fatal(err)
	}

	for msg := range msgs.Messages() {
		fmt.Println("receive message 1==> ", string(msg.Data()))
		counter++
		msg.Ack()
	}

	fmt.Println("counter number ==> ", counter)
	if msgs.Error() != nil {
		log.Fatal(msgs.Error())
	}

	ret, err := cons.Consume(func(msg jetstream.Msg) {
		fmt.Println("receive message 3==> ", string(msg.Data()))
		counter++
		msg.Ack()

		fmt.Println("callback counter number ==> ", counter)
	})

	if err != nil {
		log.Fatal(err)
	}
	defer ret.Stop()

	iter, err := cons.Messages()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 20; i++ {
		msg, err := iter.Next()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("receive message 2==> ", string(msg.Data()))
		counter++
		msg.Ack()
	}
	iter.Stop()

	fmt.Println("counter number ==> ", counter)
	time.Sleep(time.Second * 5)
}
