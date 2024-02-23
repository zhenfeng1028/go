package main

func main() {
	x := []string{"A", "B", "C"}

	for i, s := range x {
		print(i, s, " ")
		x = append(x, "Z")
		x[i+1] = "Z"
	}
}

// the ranged container is a copy of the initial x, and elements of the copy are never changed.
