package main

func f() bool {
	return false
}

func main() {
	switch f() {
	case true:
		println(1)
	case false:
		println(0)
	}
}
