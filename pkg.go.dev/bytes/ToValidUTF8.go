package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%s\n", bytes.ToValidUTF8([]byte("abc"), []byte("\uFFFD")))
	fmt.Printf("%s\n", bytes.ToValidUTF8([]byte("a\xffb\xC0\xAFc\xff"), []byte("")))
	fmt.Printf("%s\n", bytes.ToValidUTF8([]byte("\xed\xa0\x80"), []byte("abc")))
}
