package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	for _, i := range []int64{1, 2, 3, 4} {
		b := w.AvailableBuffer()
		b = strconv.AppendInt(b, i, 10)
		b = append(b, ' ')
		w.Write(b)
	}
	w.Flush()
}
