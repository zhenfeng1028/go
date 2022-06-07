// %p	表示为十六进制，并加上前导的0x

package main

import "fmt"

func main() {
	a := 18
	fmt.Printf("%p\n", &a)
	fmt.Printf("%#p\n", &a)
}
