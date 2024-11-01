package main

import (
	"fmt"
	webview "github.com/webview/webview_go"
	"os"
	"time"
)

func main() {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Bind Example")
	w.SetSize(800, 600, webview.HintNone)

	data, err := os.ReadFile("/Users/yostar/Develop/Go/GoPath/src/rain.com/ideas/webview/chat17/index.html")
	if err != nil {
		panic(err)
	}

	w.SetHtml(string(data))

	w.Bind("increment1", func() Resp {
		return Resp{Count: fmt.Sprintf("increment1:%s", fmt.Sprint(time.Now().Unix()))}
	})
	w.Bind("increment2", func() Resp {
		return Resp{Count: fmt.Sprintf("increment2:%s", fmt.Sprint(time.Now().Unix()))}
	})
	w.Run()
}

type Resp struct {
	Count string `json:"count"`
}
