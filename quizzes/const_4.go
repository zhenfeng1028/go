package main

import "unsafe"

const S = "go" // S[1]-S[0] == 8

func main() {
	var x *[8][8]byte
	println(unsafe.Sizeof((*x)[S[1]-S[0]][S[1]-S[0]]))
}

/*
	Key points:

	For a constant string S, the expression S[i] is always treated as a non-constant. So the program compiles.

	unsafe.Sizeof calls are evaluated at compile time. In the evaluations, only the type information of the arguments are used.
	The expressions *x and S[1]-S[0] etc are never evaluated at run time. So no "index out of range" panics will occur.
*/
