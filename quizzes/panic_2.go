package main

func main() {
	defer func() {
		println(recover().(int))
	}()
	defer func() {
		defer func() {
			recover()
		}()
		defer recover() // no-op
		panic(3)
	}()
	defer func() {
		defer func() {
			defer func() {
				recover() // no-op
			}()
		}()
		defer recover() // no-op
		panic(2)
	}()
	panic(1)
}

/*
	Key points:

	The recover calls at line 9 recovers the panic 3.

	The recover calls at line 5 recovers the panic 2.

	The the panic 1 is never recovered, but it is suppressed by the panic 2.
*/
