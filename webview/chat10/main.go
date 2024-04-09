//go:build windows
// +build windows

package main

import "github.com/jchv/go-webview2"

func main() {
	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Debug:     false,
		AutoFocus: true,
		WindowOptions: webview2.WindowOptions{
			Title:  "显示差异文件",
			Width:  1920,
			Height: 1080,
			IconId: 2, // icon resource id
			Center: true,
		},
	})
	defer w.Destroy()
	w.SetSize(1920, 1080, webview2.HintNone)
	w.Navigate("https://en.m.wikipedia.org/wiki/Main_Page")
	w.Run()
}
