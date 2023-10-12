package main

var f = func(x int) {}

func Bar() {
	f := func(x int) {
		if x >= 0 {
			print(x)
			f(x - 1)
		}
	}
	f(2)
}

func Foo() {
	f = func(x int) {
		if x >= 0 {
			print(x)
			f(x - 1)
		}
	}
	f(2)
}

func main() {
	Bar()
	print(" | ")
	Foo()
}

/*
	Key points:

	In the Bar function, the second f (the one in the call f(x - 1)) is the package-level f.

	In the Foo function, all three f are the package-level f.
*/
