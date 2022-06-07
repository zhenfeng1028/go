package main

import (
	"fmt"
	"strings"
)

func main() {
	var a []string = []string{"aaa", "bbb", "ccc"}
	b := strings.Join(a, ",")
	fmt.Println(b)
}
