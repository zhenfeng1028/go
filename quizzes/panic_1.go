package main

import "fmt"

func main() {
	defer func() {
		fmt.Print(recover())
	}()
	defer func() {
		defer fmt.Print(recover())
		defer panic(1)
		recover() // no-op
	}()
	defer recover() // no-op
	panic(2)
}

/*
	Key points:

	The recover calls at line 10 catches the panic 2.

	The recover calls at line 7 catches the panic 1.
*/
