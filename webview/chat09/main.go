package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	var a A
	flag.IntVar(&a.X, "x", 0, "this is version")
	flag.Parse()
	fmt.Println("a --> ", a)

	if len(os.Args) == 0 {
		panic("os.Args is nil")
	}
	fmt.Println("1 -> ", os.Args[0])
	fmt.Println(filepath.Abs(os.Args[0]))
	root, err := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(root, err)
}

type A struct {
	B
}

type B struct {
	X int
}
