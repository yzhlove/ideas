package main

import (
	"encoding/json"
	"fmt"
	webview "github.com/webview/webview_go"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {

	//fmt.Println("report ==> ", getReportCaptchaDataApi())

	//fmt.Println("local ==> ", getCaptchaDataApi())

	testVVV()

}

func testVVV() {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Bind Example")
	w.SetSize(800, 600, webview.HintNone)

	data, err := os.ReadFile("/Users/yostar/Develop/Go/GoPath/src/rain.com/ideas/webview/chat14/index.html")
	if err != nil {
		panic(err)
	}

	w.SetHtml(string(data))
	w.Bind("getCaptchaDataApi", func() CaptchaData {
		//res := getCaptchaDataApi()
		res := getReportCaptchaDataApi()
		return CaptchaData{Data: res}
	})

	w.Bind("checkCaptDataApi", func(data string) CaptchaData {
		res := checkCaptchaDataApi(data)
		return CaptchaData{Data: res}
	})

	w.Run()
}

type CaptchaData struct {
	Data string `json:"data"`
}

func getCaptchaDataApi() string {

	type Resp struct {
		Msg string `json:"message"`
	}

	client := http.Client{Timeout: time.Second * 5}
	resp, err := client.Get("http://10.155.120.139:9002/api/go-captcha-data/click-basic")
	if err == nil {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		return string(body)
	}
	res := Resp{Msg: fmt.Sprint(err)}
	data, _ := json.Marshal(res)
	return string(data)
}

func getReportCaptchaDataApi() string {

	smap := make(map[string]string)
	smap[httpTimestamp] = fmt.Sprintf("%d", time.Now().Unix())
	smap[httpNonce] = GenerateCharacter(32)
	secret, err := GenerateSign(smap)
	if err != nil {
		panic(err)
	}
	smap[httpSign] = secret

	values := url.Values{}
	for k, v := range smap {
		values.Set(k, v)
	}

	fmt.Println("values.string ==> ", values.Encode())

	address := fmt.Sprintf("%s?%s",
		"http://10.155.121.231:8008/api/user/captcha/get",
		values.Encode())

	c := &http.Client{Timeout: time.Second * 5}
	resp, err := c.Get(address)
	if err != nil {
		panic(err)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	type T struct {
		Success bool           `json:"Success"`
		Result  map[string]any `json:"Result"`
	}

	var xt = &T{}
	if err := json.Unmarshal(dat, xt); err != nil {
		panic(err)
	}

	if xt.Result != nil {
		xt.Result["code"] = 0
	}

	xxx, _ := json.Marshal(xt.Result)

	//fmt.Println("xx ==> ", xt.Success, string(xt.Result))
	return string(xxx)
}

func checkCaptchaDataApi(data string) string {

	type Req struct {
		Key  string `json:"key"`
		Dots string `json:"dots"`
	}

	type Resp struct {
		Msg string `json:"message"`
	}

	var req = &Req{}
	if len(data) != 0 {
		if err := json.Unmarshal([]byte(data), req); err != nil {
			panic(err)
		}
	}

	fmt.Println("key ==> ", req.Key)
	fmt.Println("dots ==> ", req.Dots)

	values := url.Values{}
	values.Set("dots", req.Dots)
	values.Set("key", req.Key)

	r, err := http.NewRequest(
		http.MethodPost,
		"http://10.155.120.139:9002/api/go-captcha-check-data/click-basic",
		strings.NewReader(values.Encode()))

	if err != nil {
		panic(err)
	}

	r.Header.Set("Content-type", "application/x-www-form-urlencoded")
	client := http.Client{Timeout: time.Second * 5}
	ret, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	dat, err := io.ReadAll(ret.Body)
	if err != nil {
		panic(err)
	}
	defer ret.Body.Close()
	return string(dat)
}
