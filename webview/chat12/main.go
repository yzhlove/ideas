package main

import (
	webview "github.com/webview/webview_go"
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

type IncrementResult struct {
	Count uint `json:"count"`
}

func main() {
	testWebView()
}

func testWebView() error {
	var count uint = 1
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Bind Example")
	w.SetSize(800, 600, webview.HintNone)

	// A binding that increments a value and immediately returns the new value.
	w.Bind("increment", func() IncrementResult {
		count *= factor
		return IncrementResult{Count: count}
	})

	w.SetHtml(html)
	w.Run()
	return nil
}

var factor uint = 100
