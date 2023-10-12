package main

func f(n int) (r int) {
	a, r := n-1, n+1
	if a+a == r {
		c, r := n, n*n
		r = r - c
	}
	return r
}

func main() {
	println(f(3))
}

/*
	Key point:

	The r varialbe declared in the inner block is not a re-declaration of the one (the return result) declared in the function top block.
	It is totally a new declared variable. So the return result is not modified in the inner block.
*/

/*
	Go spec says: Unlike regular variable declarations, a short variable declaration may redeclare variables provided they were
	originally declared earlier in the same block (or the parameter lists if the block is the function body) with the same type,
	and at least one of the non-blank variables is new. As a consequence, redeclaration can only appear in a multi-variable short declaration.
	Redeclaration does not introduce a new variable; it just assigns a new value to the original.

		field1, offset := nextField(str, 0)
		field2, offset := nextField(str, offset)  // redeclares offset
*/
