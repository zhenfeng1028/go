package main

func f() bool {
	return false
}

func main() {
	switch f(); // this semicolon can be omitted
	{
	case true:
		println(1)
	case false:
		println(0)
	}
}

/*
	Key points:

	Go compilers will insert semicolons automatically for Go code in compiling.

	The default switch expression is true (of the builtin type bool).
*/
