package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	webview "github.com/webview/webview_go"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var file *os.File

func init() {
	path := "/Users/yostar/Develop/Go/GoPath/src/rain.com/ideas/webview/chat16/captha.log"
	var err error
	file, err = os.OpenFile(path, os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func closeFile() {
	if file != nil {
		file.Close()
	}
}

func wLog(dat string) {
	if file != nil {
		file.WriteString(dat + "\n")
		file.Sync()
	}
}

func main() {

	defer closeFile()

	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Bind Example")
	w.SetSize(800, 600, webview.HintNone)

	data, err := os.ReadFile("/Users/yostar/Develop/Go/GoPath/src/rain.com/ideas/webview/chat16/index.html")
	if err != nil {
		panic(err)
	}
	w.SetHtml(string(data))

	w.Bind("getCaptchaDataApi", func() RespData {
		res := getReportCaptchaDataApi()
		return RespData{Data: res}
	})

	w.Bind("checkCaptchaDataApi", func(data string) RespData {
		wLog(fmt.Sprintf("resp-> checkCaptchaDataApi params --> %s ", data))
		res := checkReportCaptchaDataApi(data)
		wLog(fmt.Sprintf("resp-> checkCaptchaDataApi result --> %s ", res))
		return RespData{Data: res}
	})
	w.Run()
}

type RespData struct {
	Data string `json:"data"`
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

	wLog(fmt.Sprintf("[getReportCaptchaDataApi] values ==> %s ", values.Encode()))

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
		xt.Result["code"] = 1
	}

	xxx, _ := json.Marshal(xt.Result)
	wLog(fmt.Sprintf("[getReportCaptchaDataApi] values ==> %s ", string(xxx)))
	return string(xxx)
}

func checkReportCaptchaDataApi(data string) string {

	wLog(fmt.Sprintf("[checkReportCaptchaDataApi] data ==> %s ", data))

	type Req struct {
		Key  string `json:"key"`
		Dots string `json:"dots"`
	}

	var req = &Req{}
	if len(data) != 0 {
		if err := json.Unmarshal([]byte(data), req); err != nil {
			panic(err)
		}
	}

	wLog(fmt.Sprintf("[checkReportCaptchaDataApi] req ==> %s , %s ", req.Key, req.Dots))

	smap := make(map[string]string)
	smap[httpTimestamp] = fmt.Sprintf("%d", time.Now().Unix())
	smap[httpNonce] = GenerateCharacter(32)
	smap["key"] = req.Key
	smap["dots"] = req.Dots
	secret, err := GenerateSign(smap)
	if err != nil {
		panic(err)
	}
	smap[httpSign] = secret
	smap["idp"] = uuid.New().String()

	values := url.Values{}
	for k, v := range smap {
		values.Set(k, v)
	}

	address := fmt.Sprintf("%s?%s",
		"http://10.155.121.231:8008/api/user/captcha/verify",
		values.Encode())

	wLog(fmt.Sprintf("[checkReportCaptchaDataApi] url => %s", address))

	c := &http.Client{Timeout: time.Second * 5}

	dat := url.Values{}
	dat.Set("key", req.Key)
	dat.Set("dots", req.Dots)

	resp, err := c.Post(address, "application/x-www-form-urlencoded", strings.NewReader(dat.Encode()))
	if err != nil {
		panic(err)
	}

	ddd, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	wLog(fmt.Sprintf("[checkReportCaptchaDataApi] resp => %s", string(ddd)))

	type T struct {
		Success bool           `json:"Success"`
		Result  map[string]any `json:"Result"`
	}

	var xt = &T{}
	if err := json.Unmarshal(ddd, xt); err != nil {
		panic(err)
	}

	if xt.Success {
		xt.Result["code"] = 1
	}
	xxx, _ := json.Marshal(xt.Result)
	return string(xxx)

}
