package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {

	str := "www.topgoer.comwww.topgoer.comwww.topgoer.comwww.topgoer.comwww.topgoer.comwww.topgoer.comwww.topgoer.comwww.topgoer.comwww.topgoer.comwww.topgoer.com"
	fmt.Printf("string : %v\n", str)

	input := []byte(str)
	fmt.Printf("[]byte : %v\n", input)

	// 演示base64编码
	encodeString := base64.StdEncoding.EncodeToString(input)
	fmt.Printf("encode base64 : %v\n", encodeString)

	// 对上面的编码结果进行base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("decode base64 : %v\n", string(decodeBytes))

	fmt.Println()

	// 如果要用在url中，需要使用URLEncoding
	urlencode := base64.URLEncoding.EncodeToString([]byte(input))
	fmt.Printf("urlencode : %v\n", urlencode)
	// URLEncoding
	urldecode, err := base64.URLEncoding.DecodeString(urlencode)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("urldecode : %v\n", string(urldecode))
}
