package main

import (
	"fmt"
)

func main() {
	multiplyExample()
}

func multiplyExample() {
	// Example numbers (as base 2^m chunks)
	a := []int64{1, 2, 3, 4} // Represents 1 + 2x + 3x^2 + 4x^3
	b := []int64{2, 3, 4, 5} // Represents 2 + 3x + 4x^2 + 5x^3

	//// Choose modulus M = 2^m + 1 (for m = 32, M = 4294967297)
	mod := int64(4294967297)
	//mod := int64(4294967311)

	// Multiply using Schönhage–Strassen
	result := schoenhageStrassenMultiply(a, b, mod)

	fmt.Println("Multiplication Result:", result)
}
