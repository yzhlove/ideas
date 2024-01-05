package main

import (
	"flag"
	"fmt"
)

func main() {

	var a A
	flag.IntVar(&a.X, "x", 0, "this is version")
	flag.Parse()

	fmt.Println("a --> ", a)
}

type A struct {
	B
}

type B struct {
	X int
}
