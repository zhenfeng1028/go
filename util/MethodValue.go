package main

import (
	"fmt"
)

type Thing interface {
	GetItem() float64
	SetItem(float64)
}

type ProtoThing struct {
	Item          float64
	getItemMethod func() float64
	setItemMethod func(float64)
}

func (t ProtoThing) GetItem() float64  { return t.getItemMethod() }
func (t ProtoThing) SetItem(x float64) { t.setItemMethod(x) }

func main() {
	t := ProtoThing{}

	t.getItemMethod = func() float64 { return t.Item }
	t.setItemMethod = func(x float64) { t.Item = x }

	t.setItemMethod(2.0)

	fmt.Println(t.getItemMethod())
}
