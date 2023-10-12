package main

const (
	_    = 6
	A, _ = iota, iota + 10
	_, _
	_, B
)

func main() {
	println(A, B)
}

/*
	Key points:

	Go spec says: Within a parenthesized const declaration list the expression list may be omitted from any but the first ConstSpec.
	Such an empty list is equivalent to the textual substitution of the first preceding non-empty expression list and its type if any.

	the value of the prededeclared iota is the constant specification order id (0-based) in a constant declaration.

	The constant declaration in the quiz code contains 4 constant specifications.
	By the just mentioned rules, the declaration is equivalent to the following one:

		const (
			_    = 6
			A, _ = 1, 1 + 10
			_, _ = 2, 2 + 10
			_, B = 3, 3 + 10
		)
*/
