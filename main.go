package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func download(w http.ResponseWriter, req *http.Request) {
	var arg string = req.URL.String()
	// values := req.URL.Query()
	// arg = values.Get("url")
	arg = arg[14:]
	fmt.Println("arg", arg)
	// fmt.Println("url", req.URL)
	// resp, err := http.Get(arg)
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	req1, err := http.NewRequest("GET", arg, nil)
	// ...
	req.Header.Add("Referer", `https://douyin.com"`)
	resp, err := client.Do(req1)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	// fmt.Println("resp.Header", resp.Header)
	for name, headers := range resp.Header {
		// fmt.Printf("aaaaaa%v: %v\n", name, headers)
		for _, h := range headers {
			// fmt.Printf("bbbbb%v: %v\n", name, h)
			w.Header().Set(name, h)
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	fmt.Println(resp.StatusCode)
	w.Write(body)
	w.WriteHeader(resp.StatusCode)
	// if resp.StatusCode == 200 {
	// 	fmt.Println("ok")
	// }
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}
func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
func main() {
	http.HandleFunc("/download", download)
	http.ListenAndServe(":8090", nil)
}
