package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

func main() {
	buf := new(bytes.Buffer)
	var pi float64 = math.Pi
	err := binary.Write(buf, binary.LittleEndian, pi)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("% x", buf.Bytes())
}

// Big-endian. A computer architecture that stores multiple-byte numerical values with the most significant byte (MSB) values first.
// Little-endian. A computer architecture that stores multiple-byte numerical values with the least significant byte (LSB) values first.
