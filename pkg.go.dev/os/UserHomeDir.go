package main

import (
	"fmt"
	"os"
)

func main() {
	home, _ := os.UserHomeDir()
	fmt.Printf("user home directory: %s\n", home)
}
