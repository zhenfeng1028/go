package main

import "fmt"

type BuiltinFormat int

const (
	Binary BuiltinFormat = iota
	TextMap
	HTTPHeaders
)

func do(format interface{}) {
	switch format {
	case Binary:
		fmt.Println("Binary")
	case TextMap:
		fmt.Println("TextMap")
	case HTTPHeaders:
		fmt.Println("HTTPHeaders")
	default:
		fmt.Println("unknown format")
	}
}

func main() {
	do(Binary)
	do(HTTPHeaders)
	do(BuiltinFormat(3))
}
