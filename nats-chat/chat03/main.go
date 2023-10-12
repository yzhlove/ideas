package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"sync"
	"time"
)

func main() {

	// 同步订阅消息

	go func() {
		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()

				nc, err := nats.Connect("nats://127.0.0.1:4333")
				catch(err)
				defer nc.Close()

				sub, err := nc.SubscribeSync("chat")
				catch(err)

				msg, err := sub.NextMsg(time.Second)
				catch(err)

				fmt.Println("i:", i+1, " msg.data ---> ", string(msg.Data))

			}(i)
		}
		wg.Wait()
	}()

	time.Sleep(time.Millisecond * 100)

	nc, err := nats.Connect("nats://127.0.0.1:4333")
	catch(err)
	defer nc.Close()

	catch(nc.Publish("chat", []byte("what are you doing ???")))

	time.Sleep(time.Second)

}

func catch(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
