package main

import (
	webview "github.com/webview/webview_go"
	"io"
	"os"
)

var View webview.WebView

func main() {
	View = webview.New(true)
	defer View.Destroy()
	View.SetTitle("Basic Example")
	View.SetSize(1024, 768, webview.HintNone)

	file, err := os.Open("/Users/yostar/Desktop/xxx.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	//View.Navigate("data:text/html;base64," + base64.StdEncoding.EncodeToString(bytes))
	View.SetHtml(string(bytes))
	View.Run()
}
