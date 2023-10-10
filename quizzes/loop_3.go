package main

type Book struct {
	Pages int
}

func f() int {
	var books = []Book{{555}}
	for _, book := range books {
		book.Pages = 999
	}
	return books[0].Pages
}

func g() int {
	var books = []*Book{{555}}
	for _, book := range books {
		book.Pages = 999
	}
	return books[0].Pages
}

func main() {
	println(f(), g())
}

/*
	Key points:

	In a loop, each container element is copied to the second iteration variable,
	modifying the iteration variable doesn't change the values of container elements.

	A struct value owns its fields. When a struct value is copied, all of its fields are also copied.

	A pointer to a struct value doesn't own the fields of the struct value.
	It just references the struct value, a.k.a. it just references the fields of the struct value.
	When a pointer is copied, the values referenced by the pointer are not copied.
	After copying, the copy of the pointer references the same values as the original pointer.
*/
