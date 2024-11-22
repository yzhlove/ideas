package main

import (
	webview "github.com/webview/webview_go"
	"os"
)

func main() {

	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Bind Example")
	w.SetSize(352, 400, webview.HintNone)

	data, err := os.ReadFile("/Users/yostar/Develop/Go/GoPath/src/rain.com/ideas/webview/chat14/index.html")
	if err != nil {
		panic(err)
	}

	w.SetHtml(string(data))

	//w.Navigate(`file:///Users/yostar/Develop/Go/GoPath/src/rain.com/ideas/webview/chat13/test.html`)
	w.Run()

}
