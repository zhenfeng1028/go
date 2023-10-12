package main

const X = 3

func main() {
	const (
		X = X + X
		Y
	)

	println(X, Y)
}

/*
Key points:

	a local identifier will shadow the global identifier with the same name.

	The consntant declaration in the quiz code is equivalent to the following one:

		const (
			X = X + X // here the two "X" are both the global one
			Y = X + X // here the two "X" are both the local one
		)

	The local X is evaluated as 6 at compile time, so the constant Y is evaluaed as 12 (also at compile time).

	Please note that, the output result was 6 6 when using Go toolchain v1.17-. The bug has been fixed since Go toochain v1.18.

	Similarly, since Go toolchain 1.18, the following program prints 0 0. Before Go toolchain 1.18, it printed 1 2 (a bug).

		package main

		func main() {
			const (
				iota = iota
				X
				Y
			)
			println(X, Y)
		}
*/
