package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(regexp.Match(`\w+`, []byte("hello")))
	fmt.Println(regexp.MatchString(`\d+`, "hello"))
}
