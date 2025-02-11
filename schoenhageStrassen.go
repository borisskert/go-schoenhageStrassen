package main

import "fmt"

// Schönhage–Strassen Multiplication
func schoenhageStrassenMultiply(A, B []int64, mod int64) []int64 {
	//n := int64(len(A))

	omega, err := findOmega(32, mod) // Find a primitive n-th root of unity
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	// Step 1: Apply NTT to both inputs
	A_ntt := NTT(A, omega, mod)
	B_ntt := NTT(B, omega, mod)

	// Step 2: Pointwise multiplication
	C_ntt := pointwiseMultiply(A_ntt, B_ntt, mod)

	// Step 3: Apply INTT to get the final coefficients
	omegaInv := modInverse(omega, mod)
	C := INTT(C_ntt, omegaInv, mod)

	// Ensure the final result is non-negative
	for i := 0; i < len(C); i++ {
		if C[i] < 0 {
			C[i] += mod // Fix negative results
		}
	}

	return C
}

// Pointwise multiplication of two NTT-transformed polynomials
func pointwiseMultiply(A, B []int64, mod int64) []int64 {
	n := len(A)
	C := make([]int64, n)
	for i := 0; i < n; i++ {
		C[i] = (A[i] * B[i]) % mod
		if C[i] < 0 {
			C[i] += mod // Ensure non-negative result
		}
	}
	return C
}

func multiplyExample() {
	// Example numbers (as base 2^m chunks)
	a := []int64{1, 2, 3, 4} // Represents 1 + 2x + 3x^2 + 4x^3
	b := []int64{2, 3, 4, 5} // Represents 2 + 3x + 4x^2 + 5x^3

	//// Choose modulus M = 2^m + 1 (for m = 32, M = 4294967297)
	//mod := int64(7681)
	//mod := int64(11)
	mod := int64(4294967297)

	// Multiply using Schönhage–Strassen
	result := schoenhageStrassenMultiply(a, b, mod)

	fmt.Println("Multiplication Result:", result)
}
