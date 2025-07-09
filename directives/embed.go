package main

import (
	"embed"
	_ "embed"
)

//go:embed hello.txt
var s string

//go:embed hello.txt
var b []byte

//go:embed hello.txt
var f embed.FS // read-only

func main() {
	println(s)

	println(string(b))

	data, _ := f.ReadFile("hello.txt")
	println(string(data))
}
