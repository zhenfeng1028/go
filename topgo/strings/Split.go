package main

import (
	"fmt"
	"strings"
)

func main() {
	topic1 := "jtcs_2730_01"

	s := strings.Split(topic1, "_")
	for index, v := range s {
		fmt.Printf("index: %d, value: %s\n", index, v)
	}
}
