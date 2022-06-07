package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpResp struct {
	Id   string
	Resp *http.Response
	Err  error
}

func AsyncGet(urls map[string]string) []*HttpResp {
	ch := make(chan *HttpResp)
	responses := []*HttpResp{}

	for track_id, url := range urls {

		go func(i, u string) {
			resp, err := http.Get(u)
			ch <- &HttpResp{i, resp, err}
		}(track_id, url)
	}

loop:
	for {
		select {
		case r := <-ch:
			responses = append(responses, r)
			if len(responses) == len(urls) {
				break loop
			}
		case <-time.After(1 * time.Millisecond):
			fmt.Printf(".")
		}
	}
	return responses
}

func main() {
	urls := make(map[string]string)
	urls["1"] = "http://100.100.151.241:8765/v1/charge_unit"
	urls["2"] = "http://100.100.151.241:8765/v1/charge_unit"

	httpResps := AsyncGet(urls)
	for _, httpResp := range httpResps {
		defer httpResp.Resp.Body.Close()

		data, err := ioutil.ReadAll(httpResp.Resp.Body)
		if err != nil {
			fmt.Println("err:", err)
		}
		if httpResp.Resp.StatusCode/100 != 2 {
			fmt.Println("err:", err)
		}

		fmt.Println(string(data))
	}
}
