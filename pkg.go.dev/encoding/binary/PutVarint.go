package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	buf := make([]byte, binary.MaxVarintLen64)

	for _, x := range []int64{-65, -64, -2, -1, 0, 1, 2, 63, 64} {
		n := binary.PutVarint(buf, x)
		fmt.Printf("%d\t%x\n", x, buf[:n])
	}
}

// reference
// https://protobuf.dev/programming-guides/encoding/#signed-ints

// sintN uses the "ZigZag" encoding instead of two's complement to encode negative integers.
// Positive integers p are encoded as 2 * p (the even numbers), while negative integers n are encoded as 2 * |n| - 1 (the odd numbers).
// The encoding thus "zig-zags" between positive and negative numbers. For example:

// Signed Original	Encoded As
// 0				0
// -1				1
// 1				2
// -2				3
// …				…
// 0x7fffffff		0xfffffffe
// -0x80000000		0xffffffff

// In other words, each value n is encoded using

// (n << 1) ^ (n >> 31)

// for sint32s, or

// (n << 1) ^ (n >> 63)

// for the 64-bit version.
