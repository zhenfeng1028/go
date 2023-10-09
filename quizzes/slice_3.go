package main

import "fmt"

var x = []int{2: 5, 6, 0: 7}

func main() {
	fmt.Println(x)
}

/*
	Key points:

	keyed elements (here they are 2: 5 and 0: 7) and un-keyed elements (here it is 6) may be mixed up in array and slice composite literals.

	the index of an un-keyed element is determined by the index of the last keyed element.
	If all previous elements are un-keyed, then its index is its order number (zero based). So the index of the element 6 is 3.

	the length of an array or slice composite literal is the largest element index plus one.

	unspecified elements are viewed as the zero value of the element type.
*/
