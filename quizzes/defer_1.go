package main

type Foo struct{}

func MakeFoo(n *int) Foo {
	println(*n)
	return Foo{}
}

func (Foo) Bar(n *int) {
	println(*n)
}

func main() {
	x := 1
	p := &x
	defer MakeFoo(p).Bar(p) // line 17
	x = 2
	p = new(int) // line 19
	MakeFoo(p)
}

/*
	Key points:

	When a function call is pushed into the deferred call queue, the called function value and all the arguments are evaluated.
	The evaluated values will be used when the call is executed later in the existing phase of its caller function.

	So, at line 17, when the deferred call Bar is pushed into the deferred call queue,
	its arguments MakeFoo(p) (as the receiver argument) and p are evaluated. In evaluating MakeFoo(p), 1 is printed.

	The later modification at line 19 doesn't affect the evaluation results at line 17,
	which means the argument passed to the Bar function call is still a pointer to the variable x.
*/
