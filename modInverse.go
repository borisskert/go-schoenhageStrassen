package main

import "fmt"

// Fermat’s Little Theorem (When mod is Prime)
func modInverseFermat(n, mod int64) int64 {
	return modExp(n, mod-2, mod) // Uses modExp function from before
}

// Compute modular inverse using Extended Euclidean Algorithm
func modInverse(n, mod int64) int64 {
	g, x, _ := extendedGCD(n, mod)
	if g != 1 {
		panic("Inverse doesn't exist") // Only exists if gcd(n, mod) = 1
	}
	return (x%mod + mod) % mod // Ensure non-negative result
}

func extendedGCD(a, b int64) (int64, int64, int64) {
	if a == 0 {
		return b, 0, 1
	}
	g, x1, y1 := extendedGCD(b%a, a)
	x := y1 - (b/a)*x1
	y := x1
	return g, x, y
}

//func modInverse(n, mod int64) int64 {
//	originalMod := mod
//	x0, x1 := int64(0), int64(1)
//
//	if mod == 1 {
//		return 0 // No modular inverse
//	}
//
//	for n > 1 {
//		if mod == 0 {
//			panic("No modular inverse exists") // n and mod must be coprime
//		}
//		q := n / mod
//		n, mod = mod, n%mod
//		x0, x1 = x1-q*x0, x0
//	}
//
//	// Ensure result is positive
//	for x1 < 0 {
//		x1 += originalMod
//	}
//	return x1
//}

func modInverseExample() {
	N, mod := int64(4), int64(4294967297) // M = 2^32 + 1
	result := modInverse(N, mod)
	fmt.Println("Modular inverse of", N, "mod", mod, "is", result)

	// Verify correctness
	fmt.Println("Check:", (N*result)%mod) // Should be 1
}
