package main

import (
	"flag"
	"fmt"
	webview "github.com/webview/webview_go"
	"os"
	"os/exec"
	"time"
)

const html = `<button id="increment">Tap me</button>
<div>You tapped <span id="count">0</span> time(s).</div>
<script>
  const [incrementElement, countElement] =
    document.querySelectorAll("#increment, #count");
  document.addEventListener("DOMContentLoaded", () => {
    incrementElement.addEventListener("click", () => {
      window.increment().then(result => {
        countElement.textContent = result.count;
      });
    });
  });
</script>`

func main() {
	t := flag.Bool("m", true, "open single module. ")
	flag.Parse()
	if *t {
		path := "/Users/yostar/Develop/Go/GoPath/src/rain.com/ideas/webview/chat18/" + appName
		cmd := exec.Command(path, "--m=false")
		data, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error => ", err)
		} else {
			fmt.Println("Data => ", string(data))
		}
	} else {
		openChrome()
	}
}

const appName = "captcha"

type IncrementResult struct {
	Count string `json:"count"`
}

func openChrome() {

	d := make(chan struct{})
	go func() {
		<-d
		fmt.Println("sleep before")
		time.Sleep(time.Second * 3)
		fmt.Println("sleep after")
		os.Exit(0)
	}()
	var count = 1
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Bind Example")
	w.SetSize(800, 600, webview.HintNone)

	w.SetHtml(html)
	w.Bind("increment", func() IncrementResult {
		count++
		if count == 10 {
			d <- struct{}{}
		}
		fmt.Println("count print ==> ", count)
		return IncrementResult{Count: fmt.Sprintf("count=>%d", count)}
	})
	w.Run()
}
