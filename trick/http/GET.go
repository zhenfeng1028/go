package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	apiUrl := "http://127.0.0.1:9090/"
	// URL param
	params := url.Values{}
	params.Set("name", "lzf")
	params.Set("age", "18")
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("parse url failed, err: %v\n", err)
	}
	u.RawQuery = params.Encode()
	fmt.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("get failed, err: %v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err: %v\n", err)
		return
	}
	fmt.Println(string(b))
}
