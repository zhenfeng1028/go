package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File contents: %s\n", content)
}
