package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := fractionToBinary(6.986, 8)
	fmt.Printf("%s\n", result)
}

func fractionToBinary(fraction float64, precision int) string {
	// Split the fraction into its integer and fractional parts.
	integerPart := int(fraction)
	fractionalPart := fraction - float64(integerPart)

	// Convert the integer part to binary.
	integerPartBinary := strconv.FormatInt(int64(integerPart), 2)

	// Convert the fractional part to binary.
	fractionalPartBinary := ""
	for precision > 0 {
		fractionalPart *= 2
		if fractionalPart >= 1 {
			fractionalPartBinary += "1"
			fractionalPart -= 1
		} else {
			fractionalPartBinary += "0"
		}
		precision--
	}

	// Return the binary representation of the fraction.
	return integerPartBinary + "." + fractionalPartBinary
}
