package main

func main() {
	i, s := 9, []int{}

	for i = range s {
	}
	print(i) // 9

	for i = 0; i < len(s); i++ {
	}
	print(i) // 0

	s = append(s, 1, 2, 3, 4, 5)

	for i = range s {
	}
	print(i) // 4

	for i = 0; i < len(s); i++ {
	}
	println(i) // 5
}
