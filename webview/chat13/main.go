package main

import (
	webview "github.com/webview/webview_go"
	"os"
)

func main() {
	var count int32 = 1
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Bind Example")
	w.SetSize(800, 600, webview.HintNone)

	data, err := os.ReadFile("/Users/yostar/Develop/Go/GoPath/src/rain.com/ideas/webview/chat13/test.html")
	if err != nil {
		panic(err)
	}

	data = data
	//w.SetHtml(string(data))

	w.Bind("increment", func() A {
		count += 15
		return A{Count: count}
	})
	w.Navigate(`file:///Users/yostar/Develop/Go/GoPath/src/rain.com/ideas/webview/chat13/test.html`)
	w.Run()
}

type A struct {
	Count int32 `json:"count"`
}
