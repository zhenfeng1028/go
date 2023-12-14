package main

import (
	"bytes"
	"os"
	"strconv"
)

func main() {
	var buf bytes.Buffer
	for i := 0; i < 4; i++ {
		b := buf.AvailableBuffer()
		b = strconv.AppendInt(b, int64(i), 10)
		b = append(b, ' ')
		buf.Write(b)
	}
	os.Stdout.Write(buf.Bytes())
}
