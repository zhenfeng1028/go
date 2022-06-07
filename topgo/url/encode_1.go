package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlStr string = "运维之路"
	escapeUrl := url.QueryEscape(urlStr)
	fmt.Println("编码:", escapeUrl)
	enEscapeUrl, _ := url.QueryUnescape(escapeUrl)
	fmt.Println("解码:", enEscapeUrl)
}
