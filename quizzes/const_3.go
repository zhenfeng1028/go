package main

const X = '\x61' // 'a'
const Y = 0x62
const A = Y - X // 1
const B int64 = 1

var n = 32

func main() {
	if A == B {
		println(A<<n>>n, B<<n>>n)
	}
}

/*
	Key points:

	Go spec says: If the untyped operands of a binary operation (other than a shift) are of different kinds,
	the result is of the operand's kind that appears later in this list: integer, rune, floating-point, complex.
	For example, an untyped integer constant divided by an untyped complex constant yields an untyped complex constant.

	Untyped X is a rune constant (in other words, its default type is rune, a.k.a int32, a 32-bit integer type).
	Untyped Y is an int constant (in other words, its default type is int, a 64-bit integer type on 64-bit OSes).

	At run time, the expression A << n overflows, so it is evaluated as 0; on the other hand, the expression B << n doesn't overflow.
*/
