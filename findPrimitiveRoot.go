package main

import "fmt"

// Modular exponentiation: (base^exp) % mod
func modExp3(base, exp, mod int64) int64 {
	if base == 0 {
		return 0
	}

	if base == 1 {
		return 1
	}

	result := int64(1)
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exp /= 2

		if base == 0 {
			return 0
		}
	}
	return result
}

// Check if g is a primitive root modulo M
func isPrimitiveRoot(g, mod int64) bool {
	order := mod - 1 // Since M = 2^m + 1, order = 2^m
	// Factorize order (for M = 2^32 + 1, only factor is 2)
	factors := []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
		31, 37, 41, 43, 47, 53, 59, 61, 67, 71}
	for _, f := range factors {
		if modExp3(g, order/f, mod) == 1 {
			return false
		}
	}
	return true
}

// Find the smallest primitive root modulo M
func findPrimitiveRoot(mod int64) int64 {
	for g := int64(2); g < mod; g++ {
		if isPrimitiveRoot(g, mod) {
			return g
		}
	}

	panic("Should never happen for valid M = 2^m + 1")
}

func findPrimitiveRootExample() {
	mod := int64(4294967297) // M = 2^32 + 1
	//mod := findLargePrime()
	g := findPrimitiveRoot(mod)
	fmt.Println("Smallest Primitive Root g:", g)
}
