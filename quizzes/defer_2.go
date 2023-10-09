package main

func bar() (r int) {
	defer func() {
		r += 4
		if recover() != nil {
			r += 8
		}
	}()

	var f func()
	defer f()
	f = func() {
		r += 2
	}

	return 1
}

func main() {
	println(bar())
}

/*
	Key points:

	Line 12: when the call f() is pushed to the defer call queue, f is evaluated as nil.
	Calling a nil function produces a panic. But the panic doesn't happen at the push time.
	Instead, it will panic at the execution time later during the exiting phase of the bar() call.

	Line 13-15: the assignment to f has not any impacts on the execution effect. In other words, this assignment is actually a no-op.

	Line 4-9: a defer call could modify the named return result of its containing function (here the bar function).
*/
