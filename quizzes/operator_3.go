package main

const N = 1

var (
	n      = N
	a byte = 128 << N >> N // equivalent to var a = byte(int(128) << N >> N)
	b byte = 128 << n >> n // equivalent to var a = byte(128) << N >> N
)

func main() {
	println(a, b)
}

/*
	Key points:

	If the operands in an operator expression are both/all constants, then the expression is evaluated at compile time.
	In the above program, 128 << N >> N is such an expression. In this expression, 128 is deduced as an untyped int value.

	In a bit-shift expression, if the left operand is an untyped constant and the right operand is not constant,
	then the type of the left operand will be deduced as the final assumed type. In the above program, 128 << n >> n is a such expression.
	In this expression, the type of 128 is deduced as the assumed type, byte. 128 << n overflows the value range of byte, so it is truncated to 0.
*/
