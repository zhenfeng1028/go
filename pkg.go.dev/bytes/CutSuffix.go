package main

import (
	"bytes"
	"fmt"
)

func main() {
	show := func(s, sep string) {
		before, found := bytes.CutSuffix([]byte(s), []byte(sep))
		fmt.Printf("CutSuffix(%q, %q) = %q, %v\n", s, sep, before, found)
	}
	show("Gopher", "Go")
	show("Gopher", "er")
}
