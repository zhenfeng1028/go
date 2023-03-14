package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		str = "lizhenfeng is an employee of shanma"
		subStr = "lizhenfeng"
	)
	isTrue := strings.Contains(str, subStr)
	if isTrue {
		fmt.Printf("\"%s\" is a sub string of \"%s\"\n", subStr, str)
	} else {
		fmt.Printf("\"%s\" is not a sub string of \"%s\"\n", subStr, str)
	}
}
