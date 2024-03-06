package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		i := bytes.IndexByte(data, ',')
		if i == -1 {
			if !atEOF {
				return 0, nil, nil
			}
			// If we have reached the end, return the last token.
			return 0, data, bufio.ErrFinalToken
		}
		// If the token is "STOP", stop the scanning and ignore the rest.
		if string(data[:i]) == "STOP" {
			return i + 1, nil, bufio.ErrFinalToken
		}
		// Otherwise, return the token before the comma.
		return i + 1, data[:i], nil
	}
	const input = "1,2,STOP,4,"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(onComma)
	for scanner.Scan() {
		fmt.Printf("Got a token %q\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
