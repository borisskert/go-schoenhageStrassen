package main

import "fmt"

// Compute modular inverse using Extended Euclidean Algorithm
func modInverse(n, mod int64) int64 {
	originalMod := mod
	x0, x1 := int64(0), int64(1)

	if mod == 1 {
		return 0 // No modular inverse
	}

	for n > 1 {
		if mod == 0 {
			panic("No modular inverse exists") // n and mod must be coprime
		}
		q := n / mod
		n, mod = mod, n%mod
		x0, x1 = x1-q*x0, x0
	}

	// Ensure result is positive
	if x1 < 0 {
		x1 += originalMod
	}
	return x1
}

func modInverseExample() {
	N, mod := int64(4), int64(4294967297) // M = 2^32 + 1
	result := modInverse(N, mod)
	fmt.Println("Modular inverse of", N, "mod", mod, "is", result)

	// Verify correctness
	fmt.Println("Check:", (N*result)%mod) // Should be 1
}
