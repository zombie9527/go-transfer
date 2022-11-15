package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func download(w http.ResponseWriter, req *http.Request) {
	var arg string = req.URL.String()
	arg = arg[6:]
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	req1, err := http.NewRequest("GET", arg, nil)
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

	w.Write(body)
	w.WriteHeader(resp.StatusCode)

}

func main() {
	http.HandleFunc("/", download)
	http.ListenAndServe(":8090", nil)
}
