package main

import "reflect"

type I interface {
	foo()
	Bar()
}

type T struct{}

func (T) foo() {}
func (T) Bar() {}

func main() {
	var t T
	var i I = t
	x := reflect.ValueOf(t)
	y := reflect.ValueOf(&i).Elem()
	println(x.NumMethod(), y.NumMethod())
}

/*
	Key point:

	the reflect.Value.NumMethod method counts unexported methods for interface types/values,
	but only counts exported methods for non-interface types/values.
*/
