package main

func f() {
	var a = [2]int{5, 7}
	for i, v := range a {
		if i == 0 {
			a[1] = 9
		} else {
			print(v)
		}
	}
}

func g() {
	var a = [2]int{5, 7}
	for i, v := range a[:] {
		if i == 0 {
			a[1] = 9
		} else {
			print(v)
		}
	}
}

func main() {
	f()
	g()
}

/*
	Key points:

	When iterating the elements of a container (array/slice/map),
	the elements of a copy of the container is actually iterated.

	An array owns its elements. When copying an array, its elements are also copied.
	Modifying elements in an array doesn't affect the elements in the copy of the array.

	A slice just references its elements. When copying a slice (through assignment,
	not call the builtin copy function), its elements are not copied.
	After copying, the copy of the slice shares elements with the original slice.
	So modifying elements in a slice means modifying the elements in the copy of the slice.

	So in function f, the modification of a[1] is not reflected in the loop.
	But the modification of a[1] in function g is reflected in the loop.
*/
