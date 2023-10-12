package main

import "reflect"

type T struct {
	A int
	b int
}

func (T) M() {}
func (T) m() {}

func main() {
	v := reflect.ValueOf(T{})
	println(v.NumField(), v.NumMethod())
}

/*
	Key point:

	the reflect.Value.NumMethod method only counts exported methods for non-interface types/values.

	the reflect.Value.NumField method counts both exported and unexported fields.
*/
