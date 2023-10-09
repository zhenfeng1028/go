package main

func main() {
	m := make(map[string]int, 3)
	x := len(m)
	m["Go"] = m["Go"]
	y := len(m)
	println(x, y)
}

/*
	Key points:

	when using the make function to create a map, the second argument is neither the initial length nor the capacity of the result map.
	It is just a hint for Go runtime to allocate a large enough backing array to hold at least the specified number of entries.

	The length (number of entries) of the result map is zero.
	m["Go"] = m["Go"] is equivalent to m["Go"] = 0. After the assignment, the map contains one entry.
*/
