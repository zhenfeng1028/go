package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	inputs := [][]byte{
		{0x01},
		{0x02},
		{0x7f},
		{0x80, 0x01},
		{0xff, 0x01},
		{0x80, 0x02},
	}
	for _, b := range inputs {
		x, n := binary.Uvarint(b)
		if n != len(b) {
			fmt.Println("Uvarint did not consume all of in")
		}
		fmt.Println(x)
	}
}

// reference
// https://protobuf.dev/programming-guides/encoding/#varints

// 8001 <-> 128
// 10000000 00000001	// Original inputs
//  0000000  0000001	// Drop continuation bits
//  0000001  0000000	// Convert to big-endian
//    00000010000000	// Concatenate
// 2 ^ 7 = 128			// Interpret as an unsigned 64-bit integer

// ff01 <-> 255
// 11111111 00000001	// Original inputs
//  1111111  0000001	// Drop continuation bits
//  0000001  1111111	// Convert to big-endian
//    00000011111111	// Concatenate
// 2 ^ 8 - 1 = 255		// Interpret as an unsigned 64-bit integer
