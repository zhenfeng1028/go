package main

func main() {
	var y = []string{"A", "B", "C", "D"}
	var x = y[:3]

	for i, s := range x {
		print(i, s, ",")
		x = append(x, "Z")
		x[i+1] = "Z"
	}
}

/*
	Key points:

	The first append call doesn't create a new backing array, so the assignment x[i+1] = "Z"
	in the first loop has effect on the initial slice x (and its copy used in the element iteration).

	The second append call creates a new backing array, so subsequent x[i+1] = "Z"
	assignments have no effects on the initial slice x (and its copy used in the element iteration).
*/
