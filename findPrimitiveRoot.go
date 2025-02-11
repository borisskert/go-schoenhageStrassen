package main

import (
	"fmt"
	"math/big"
)

//// Fast modular exponentiation: (base^exp) % mod
//func modExp(base, exp, mod int64) int64 {
//	result := int64(1)
//	base = base % mod
//	for exp > 0 {
//		if exp%2 == 1 {
//			result = (result * base) % mod
//		}
//		base = (base * base) % mod
//		exp /= 2
//	}
//	return result
//}

//// Factorize n into its prime factors
//func factorize(n int64) []int64 {
//	factors := []int64{}
//	for i := int64(2); i*i <= n; i++ {
//		if n%i == 0 {
//			factors = append(factors, i)
//			for n%i == 0 {
//				n /= i
//			}
//		}
//	}
//	if n > 1 {
//		factors = append(factors, n)
//	}
//	return factors
//}

// Check if g is a primitive root modulo M
func isPrimitiveRoot(g, mod int64) bool {
	order := mod - 1 // φ(M) for prime M is M-1
	factors := factorize(order)

	// g^((M-1)/p) should not be 1 for any prime factor p
	for _, p := range factors {
		if modExp(g, order/p, mod) == 1 {
			return false
		}
	}
	return true
}

// Find the smallest primitive root modulo M (for any prime M)
func findPrimitiveRoot(mod int64) (int64, error) {
	// Ensure M is prime
	if !big.NewInt(mod).ProbablyPrime(20) {
		return -1, fmt.Errorf("M is not prime! No primitive root exists.")
	}

	for g := int64(2); g < mod; g++ {
		if isPrimitiveRoot(g, mod) {
			return g, nil
		}
	}
	return -1, fmt.Errorf("No primitive root found (should not happen for valid prime M)")
}

func findPrimitiveRootExample() {
	// Test with different prime numbers (not just Fermat primes)
	primes := []int64{17, 101, 65537, 7919, 982451653, 4294967297}

	for _, mod := range primes {
		primitiveRoot, err := findPrimitiveRoot(mod)
		if err != nil {
			fmt.Println("Error for M =", mod, ":", err)
		} else {
			fmt.Println("Primitive Root of", mod, ":", primitiveRoot)
		}
	}
}
