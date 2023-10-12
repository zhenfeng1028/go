package main

var x, y = true, false

func o(b bool) bool {
	print(b)
	return !b
}

func main() {
	_ = x || o(x)
	_ = y && o(y)
}

/*
	Key points:

	Go spec says: The right operand is evaluated conditionally.

	When evaluating a || b, the expression a is evaluated firstly and the expression b will be only evaluated if a is evaluated as false.

	When evaluating a && b, the expression a is evaluated firstly and the expression b will be only evaluated if a is evaluated as true.
*/
