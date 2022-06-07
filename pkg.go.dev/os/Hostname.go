package main

import (
	"fmt"
	"os"
)

func main() {
	hostname, _ := os.Hostname()
	fmt.Printf("hostname is: %s\n", hostname)
}
