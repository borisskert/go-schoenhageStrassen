package main

import (
	"fmt"
)

func checkNttInttExample() {
	A := []int64{1, 2, 3} // Input array
	n := int64(len(A))
	mod := int64(4294967297)      // Example modulus
	m := log2(mod)                // Length of A
	omega, _ := findOmega(m, mod) // Example modulus
	fmt.Println("Omega:", omega, " is valid: ", isValidOmega(omega, m, mod))

	fmt.Println("Original A:", A)

	// Compute NTT
	A_ntt := NTT(padToPowerOfTwo(A), omega, mod)
	fmt.Println("After NTT :", A_ntt)

	omegaInv := modInverse(omega, mod)
	fmt.Println("omega:", omega, "omegaInv:", omegaInv)
	fmt.Println("omega ^ n =", modExp(omega, n, mod)) // should be 1 if ω is correct

	// Compute INTT
	A_recovered := INTT(A_ntt, omegaInv, mod)

	// Verify recovery
	fmt.Println("After INTT:", A_recovered) // Should match original A
}
