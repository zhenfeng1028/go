package main

import (
	"fmt"
	"regexp"
)

var (
	p = regexp.MustCompile(`^[a-z]+\[\d+\]$`)
)

func main() {
	fmt.Println(p.MatchString("larry[12]"))
	fmt.Println(p.MatchString("jacky[12]"))
	fmt.Println(p.MatchString("linda[a12]"))
}
