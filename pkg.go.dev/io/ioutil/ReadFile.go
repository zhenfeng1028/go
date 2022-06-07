package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File contents: %s\n", content)
}
