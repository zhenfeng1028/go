package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	src := []byte("Hello")
	encodedStr := hex.EncodeToString(src)

	fmt.Printf("%s\n", encodedStr)
}

// Hello
// 0100 1000 0110 0101 0110 1100 0110 1100 0110 1111
