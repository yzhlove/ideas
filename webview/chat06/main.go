package main

import (
	webview "github.com/webview/webview_go"
)

func main() {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Basic Example")
	w.SetSize(1920, 1080, webview.HintNone)

	//u := "file:///Users/yostar/index/Language/LanguageCommon.zh_CN.csv.html"
	w.Navigate("file:////Users/yostar/index/Language/LanguageCommon.zh_CN.csv.html")
	w.Run()
}
