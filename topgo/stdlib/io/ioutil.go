package main

import (
	"fmt"
	"os"
)

func write() {
	err := os.WriteFile("./yyy.txt", []byte("www.baidu.com"), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func read() {
	content, err := os.ReadFile("./yyy.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}

func main() {
	write()
	read()
}
