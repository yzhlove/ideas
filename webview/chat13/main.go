package main

import (
	"fmt"
	webview "github.com/webview/webview_go"
	"os"
)

func main() {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Bind Example")
	w.SetSize(800, 600, webview.HintNone)

	data, err := os.ReadFile("/Users/yostar/Develop/Go/GoPath/src/rain.com/ideas/webview/chat13/test.html")
	if err != nil {
		panic(err)
	}

	data = data
	w.SetHtml(string(data))

	w.Bind("increment", func(dat string) A {

		fmt.Println("str ==> ", dat)
		return A{Count: fmt.Sprintf("yzh:%s", dat)}
	})
	//w.Navigate(`file:///Users/yostar/Develop/Go/GoPath/src/rain.com/ideas/webview/chat13/test.html`)
	w.Run()
}

type A struct {
	Count string `json:"count"`
}
