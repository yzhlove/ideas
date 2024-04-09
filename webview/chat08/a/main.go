package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {

	go openWeb()

	fmt.Println("------1 ")

	time.Sleep(time.Second * 10)
	fmt.Println("------2 ")

	time.Sleep(time.Second * 5)
	fmt.Println("------3 ")

}

func openWeb() {
	file, err := os.Open("/Users/yostar/Develop/Go/GoPath/src/rain.com/ideas/webview/chat05/main.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	cmd := exec.Command("/Users/yostar/Develop/Go/GoPath/src/rain.com/ideas/webview/chat08/b/b", "--web=true")
	cmd.Stdin = file

	res, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	fmt.Println("output=> ", string(res))

}
