package main

import (
	"fmt"
	"net/url"
)

func main() {
	params := url.Values{}
	params.Add("name", "@Rajeev")
	params.Add("phone", "+919999999999")
	fmt.Println(params.Encode())
}
