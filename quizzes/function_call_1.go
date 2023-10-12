package main

func f(vs ...interface{}) {
	print(len(vs))
}

func main() {
	f()
	f(nil)
	f(nil...)
}

/*
	Key points:

	The nil in f(nil...) is treated as a nil slice (of type []interface{}).

	The nil in f(nil) is treated as an element in a slice. The call f(nil) is equivalent to f([]interface{}{nil}...).

	The call f() is equivalent to f(nil...).
*/
