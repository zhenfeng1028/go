package main

/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: -L . -lhello
#include "hello.h"
*/
import "C"

func main() {
	C.hello()
}
