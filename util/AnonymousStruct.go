package main

import (
	"fmt"
)

func main() {
	data := struct {
		Number int
		Text   string
	}{42, "Hello world!"} // anonymous struct
	fmt.Printf("%+v\n", data)
}
