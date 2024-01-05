package main

import (
	"flag"
	"fmt"
	webview "github.com/webview/webview_go"
	"io"
	"os"
)

func main() {

	x := flag.Bool("web", false, "open web view")
	y := flag.Bool("version", false, "show version")
	flag.Parse()

	if *x {
		fmt.Println("start open view")

		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}

		View := webview.New(false)
		defer View.Destroy()
		View.SetTitle("Basic Example")
		View.SetSize(1280, 720, webview.HintNone)
		View.SetHtml(string(bytes))
		View.Run()
	}

	if *y {
		fmt.Println("version is 0.0.0")
	}
}
