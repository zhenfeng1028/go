package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	s1 := []int{3, 2, 5, 7, 4, 1, 6}
	s2 := []string{"lzf", "ggg", "zxj"}
	s3 := []Person{
		{"lzf", 29}, {"ggg", 27}, {"zxj", 26},
	}

	gensort(len(s1),
		func(i, j int) bool {
			return s1[i] < s1[j]
		},
		func(i, j int) {
			s1[i], s1[j] = s1[j], s1[i]
		},
	)
	fmt.Println(s1)

	gensort(len(s2),
		func(i, j int) bool {
			return s2[i] < s2[j]
		},
		func(i, j int) {
			s2[i], s2[j] = s2[j], s2[i]
		},
	)
	fmt.Println(s2)

	gensort(len(s3),
		func(i, j int) bool {
			return s3[i].Age < s3[j].Age
		},
		func(i, j int) {
			s3[i], s3[j] = s3[j], s3[i]
		},
	)
	fmt.Println(s3)
}

type sortType struct {
	length int
	less   func(int, int) bool
	swap   func(int, int)
}

func (s *sortType) Len() int {
	return s.length
}

func (s *sortType) Less(i, j int) bool {
	return s.less(i, j)
}

func (s *sortType) Swap(i, j int) {
	s.swap(i, j)
}

func gensort(Len int, Less func(int, int) bool, Swap func(int, int)) {
	sort.Sort(&sortType{length: Len, less: Less, swap: Swap})
}
