package main

import "fmt"

func INTT(a []int64, omegaInv int64, mod int64) []int64 {
	n := int64(len(a))
	res := NTT(a, omegaInv, mod) // Apply NTT with omegaInv

	nInv := modInverse(n, mod)                       // Compute modular inverse of n
	fmt.Println("n:", n, "mod:", mod, "nInv:", nInv) // Debugging

	// Normalize coefficients
	fmt.Println("Before scaling:", res) // Debugging
	for i := 0; i < len(res); i++ {
		res[i] = modMul(res[i], nInv, mod) // Use safe multiplication
		if res[i] < 0 {
			res[i] += mod // Ensure non-negative result
		}
	}
	fmt.Println("After scaling:", res) // Debugging
	return res
}
