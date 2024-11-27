package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := map[string]interface{}{
		"name": "lzf",
		"age":  28,
	}
	v := reflect.ValueOf(m)
	if v.Kind() == reflect.Map {
		keys := v.MapKeys()
		for _, k := range keys {
			value := v.MapIndex(k)
			fmt.Println(value.Kind())
		}
	}
}
