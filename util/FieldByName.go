// Golang program to illustrate
// reflect.FieldByName() Function

package main

import (
	"fmt"
	"reflect"
)

type Struct1 struct {
	Var1 string
	Var2 string
	Var3 float64
	Var4 float64
}

// Main function
func main() {
	NewMap := make(map[string]*Struct1)
	NewMap["abc"] = &Struct1{"abc", "def", 1.0, 2.0}
	subvalMetric := "Var1"

	for _, Value := range NewMap {
		s := reflect.ValueOf(&Value).Elem()
		println(s.String())
		println(s.Elem().String())

		// use of FieldByName() method
		metric := s.Elem().FieldByName(subvalMetric).Interface()
		fmt.Println(metric)
	}
}
