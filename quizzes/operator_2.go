package main

func main() {
	count := 0
	for i := range [256]struct{}{} {
		if n := byte(i); n == -n {
			count++
		}
	}
	println(count)
}

// 10:  0000 1010
//-10: 11111 0110 = 128+64+32+16+4+2 = 246

/*
	Key points:

	In Go, for a byte (a.k.a. uint8) non-constant non-zero value x, -x overflows the range of type byte and is wrapped as (-x + 256) % 256.

	For a variable x of type byte, x == -x happens only when x is 0 or 128.
*/
