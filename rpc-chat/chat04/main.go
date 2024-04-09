package main

import (
	"fmt"
	pb "github.com/golang/protobuf/proto"
	"github.com/yzhlove/ideas/rpc-chat/chat04/proto"
	"log"
)

func main() {
	test2()
}

func test1() {
	msg1 := &proto.Msg1{
		Name:    "yuzihan",
		Age:     26,
		Sex:     true,
		Desc:    "what are you doing?",
		Project: "mimimi",
		Job:     3,
	}

	fmt.Println("msg1 => ", msg1)

	data, err := pb.Marshal(msg1)
	if err != nil {
		log.Fatal(err)
	}

	msg2 := &proto.Msg2{}

	if err := pb.Unmarshal(data, msg2); err != nil {
		log.Fatal(err)
	}

	fmt.Println("msg2 => ", msg2)

}

func test2() {

	msg1 := &proto.Msg1{
		Name:    "yuzihan",
		Age:     26,
		Sex:     true,
		Desc:    "",
		Project: "",
		Job:     3,
	}

	msg2 := &proto.Msg2{
		Name: "yuzihan",
		Age:  26,
		Sex:  true,
		Job:  3,
	}

	data1, err := pb.Marshal(msg1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("data1 len => ", len(data1))

	data2, err := pb.Marshal(msg2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("data2 len => ", len(data2))

}
