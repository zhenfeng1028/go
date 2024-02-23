package main

import (
	"fmt"
)

func main() {
	a := struct {
		A string
		B int
	}{"lzf", 10}

	// General
	fmt.Printf("%v\n", a)  // default format
	fmt.Printf("%+v\n", a) // the plus flag (%+v) adds field names
	fmt.Printf("%#v\n", a) // Go-syntax representation of the value
	fmt.Printf("%T\n", a)  // Go-syntax representation of the type of the value

	// Boolean
	fmt.Printf("%t\n", 1 == 2) // the word true or false

	// Integer
	fmt.Printf("%b\n", 10)  // base 2
	fmt.Printf("%c\n", '😀') // the character represented by the corresponding Unicode code point
	fmt.Printf("%o\n", 10)  // base 8
	fmt.Printf("%O\n", 10)  // base 8 with 0o prefix
	fmt.Printf("%q\n", 48)  // a single-quoted character literal safely escaped with Go syntax
	fmt.Printf("%x\n", 10)  // base 16, with lower-case letters for a-f
	fmt.Printf("%X\n", 10)  // base 16, with upper-case letters for A-F
	fmt.Printf("%U\n", 255) // Unicode format: U+1234; same as "U+%04X"

	// Floating-point and complex constituents
	fmt.Printf("%b\n", 123.456) // decimalless scientific notation with exponent a power of two, e.g. -123456p-78
	fmt.Printf("%e\n", 123.456) // scientific notation, e.g. -1.234456e+78
	fmt.Printf("%E\n", 123.456) // scientific notation, e.g. -1.234456E+78
	fmt.Printf("%f\n", 123.456) // decimal point but no exponent, e.g. 123.456
	fmt.Printf("%F\n", 123.456) // synonym for %f
	fmt.Printf("%g\n", 123.456) // %e for large exponents, %f otherwise
	fmt.Printf("%G\n", 123.456) // %E for large exponents, %F otherwise
	fmt.Printf("%x\n", 123.456) // hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
	fmt.Printf("%X\n", 123.456) // upper-case hexadecimal notation, e.g. -0X1.23ABCP+20

	// String and slice of bytes
	fmt.Printf("%s\n", "lzf") // the uninterpreted bytes of the string or slice
	fmt.Printf("%q\n", "lzf") // a double-quoted string safely escaped with Go syntax
	fmt.Printf("%x\n", "lzf") // base 16, lower-case, two characters per byte
	fmt.Printf("%X\n", "lzf") // base 16, upper-case, two characters per byte

	// Slice
	fmt.Printf("%p\n", []int{1}) // address of 0th element in base 16 notation, with leading 0x

	// Pointer
	v := 10
	fmt.Printf("%p\n", &v)

	fmt.Printf("%p\n", &struct{}{})
	fmt.Printf("%p\n", &[1]int{3})
	fmt.Printf("%p\n", &[]int{3})
	fmt.Printf("%p\n", &map[int]int{1: 1})

	// Width and Precision
	fmt.Printf("%f\n", 12.3456789)       // default width, default precision
	fmt.Printf("%9f\n", 12.3456789)      // width 9, default precision
	fmt.Printf("%.2f\n", 12.3456789)     // default width, precision 2
	fmt.Printf("%9.2f\n", 12.3456789)    // width 9, precision 2
	fmt.Printf("%9.f\n", 12.3456789)     // width 9, precision 0
	fmt.Printf("%*.2f\n", 9, 12.3456789) // Either or both of the flags may be replaced with the character '*', causing their values to
	fmt.Printf("%9.*f\n", 2, 12.3456789) // be obtained from the next operand (preceding the one to format), which must be of type int.

	vbArr := [11]byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Printf("%.5s\n", "hello world")         // For strings, byte slices and byte arrays, however,
	fmt.Printf("%.5s\n", []byte("hello world")) // precision limits the length of the input to be formatted
	fmt.Printf("%.5s\n", vbArr)                 // (not the size of the output), truncating if necessary.

	fmt.Printf("%.3g\n", 12.34500) // %g/%G precision sets the maximum number of significant digits (trailing zeros are removed)
	fmt.Printf("%.3G\n", 12.34500)
	fmt.Printf("%f\n", 1.2+3.4i) // the width and precision apply to the two components independently and the result is parenthesized

	fmt.Printf("%+d\n", 10)           // '+' always print a sign for numeric values
	fmt.Printf("%-9.f\n", 12.3456789) // '-' pad with spaces on the right rather than the left (left-justify the field)

	fmt.Printf("%#q\n", "lzf")    // print a raw (backquoted) string
	fmt.Printf("%#g\n", 12.34500) // do not remove trailing zeros for %g and %G
	fmt.Printf("%#U\n", 97)       // write e.g. U+0078 'x' if the character is printable

	fmt.Printf("% d\n", 10)    // leave a space for elided sign in numbers
	fmt.Printf("% x\n", "lzf") // put spaces between bytes printing strings or slices in hex

	fmt.Printf("%09.f\n", 12.3456789) // pad with leading zeros rather than spaces

	fmt.Printf("%#q\n", []string{"lzf", "ggg"})        // For compound operands such as slices and structs, the format applies to
	fmt.Printf("%6.2f\n", []float64{12.3234, 13.2324}) // the elements of each operand, recursively, not to the operand as a whole.
	fmt.Printf("%#q\n", a)
}
