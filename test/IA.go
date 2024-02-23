package main

import "fmt"

type A struct {
	Name string
}

type IA interface {
	method() IA
}

func (a A) method() IA {
	return a
}

func main() {
	a := &A{Name: "lzf"}
	i := a.method()
	aa, ok := i.(A)
	if ok {
		fmt.Println(aa.Name)
	}
}
