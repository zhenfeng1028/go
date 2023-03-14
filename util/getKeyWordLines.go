package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const filepath = "gpu_list.txt"

func main() {
	lines, err := getKeyWordLines(filepath)
	if err != nil {
		fmt.Printf("getKeyWordLines err: %v\n", err)
	}
	for _, line := range lines {
		fmt.Println(line)
	}
}

func getKeyWordLines(filepath string) ([]string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)

	lines := make([]string, 0)
	// https://golang.org/pkg/bufio/#Scanner.Scan
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "Tesla T4") {
			lines = append(lines, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
