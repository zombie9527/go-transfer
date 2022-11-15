package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func download(w http.ResponseWriter, req *http.Request) {
	var arg string = req.URL.String()
	arg = arg[7:]

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	req1, err := http.NewRequest("GET", arg, nil)
	req1.Header.Add("Referer", `https://www.douyin.com/"`)
	resp, err := client.Do(req1)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	for name, headers := range resp.Header {

		for _, h := range headers {
			w.Header().Set(name, h)
		}
	}
	body, err := ioutil.ReadAll(resp.Body)

	w.WriteHeader(resp.StatusCode)
	w.Write(body)

}

func main() {
	http.HandleFunc("/d", download)
	http.ListenAndServe(":8090", nil)
}
