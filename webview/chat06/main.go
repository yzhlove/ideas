package main

import (
	"fmt"
	webview "github.com/webview/webview_go"
	"time"
)

var View webview.WebView

func main() {
	View = webview.New(false)
	defer View.Destroy()
	View.SetTitle("Basic Example")
	View.SetSize(1, 1, webview.HintNone)
	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("dsjfasdb")
		View.SetSize(1280, 720, webview.HintNone)
		View.Navigate("file:///Users/yostar/Desktop/rm.txt")
	}()
	View.SetHtml("what are you doing?")
	View.Run()
}
