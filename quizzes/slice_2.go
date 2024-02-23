package main

func main() {
	x := []string{"A", "B", "C"}

	for i, s := range x {
		print(i, s, ",")
		x[i+1] = "M"
		x = append(x, "Z")
		x[i+1] = "Z"
	}
}

/*
	Key points:

	Ranging over a container iterates the elements of a copy of the container (array/slice/map).
	The evaluation of the copy happens before executing the loop, so the length and capacity of the copy are never changed.

	A slice references its elements on a backing array.
	So a copy of a slice shares the same elements (and the backing array) with the slice.
	The assignment x[i+1] = "M" in the first loop step modifies the second element of the initial slice x and the copy of the initial slice x.

	If the free element slots of the first argument slice of an append call are insufficient to hold all the appended elements,
	the append call will create a new backing array to hold all the elements of the first argument slice and the appended elements.
	So, at the end of the first loop step, the backing array of the slice x is changed.
	However, the change doesn't affect the slice copy used in the element iteration.
	All subsequent element modifications apply to the new backing array, so they have no effects on the copy used in the element iteration.
*/
