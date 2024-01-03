package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	buf := make([]byte, binary.MaxVarintLen64)

	for _, x := range []uint64{1, 2, 127, 128, 255, 256} {
		n := binary.PutUvarint(buf, x)
		fmt.Printf("%x\n", buf[:n])
	}
}

// 128 <-> 8001
// 10000000 00000001	// Original inputs
//  0000000  0000001	// Drop continuation bits
//  0000001  0000000	// Convert to big-endian
//    00000010000000	// Concatenate

// 255 <-> 8001
// 11111111 00000001	// Original inputs
//  1111111  0000001	// Drop continuation bits
//  0000001  1111111	// Convert to big-endian
//    00000011111111	// Concatenate
