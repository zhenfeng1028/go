package main

import (
	"fmt"
	"io/ioutil"
)

func wr() {
	err := ioutil.WriteFile("./yyy.txt", []byte("www.baidu.com"), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func re() {
	content, err := ioutil.ReadFile("./yyy.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}

func main() {
	wr()
	re()
}
