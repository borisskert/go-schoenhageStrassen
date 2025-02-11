package main

import "fmt"

// Schönhage–Strassen Multiplication
func schoenhageStrassenMultiply(A, B []int64) []int32 {
	n := int64(len(A))

	mod, _, omega, err := findModulus2(A, B)
	if err != nil {
		panic(err)
	}

	// Step 1: Apply NTT to both inputs
	A_ntt := NTT(A, n, int64(omega), int64(mod))
	B_ntt := NTT(B, n, int64(omega), int64(mod))

	// Step 2: Pointwise multiplication
	C_ntt := pointwiseMultiply(A_ntt, B_ntt, int64(mod))

	// Step 3: Apply INTT to get the final coefficients
	omegaInv := modInverse(int64(omega), int64(mod))
	C := INTT(C_ntt, n, omegaInv, int64(mod))

	// Ensure the final result is non-negative
	for i := 0; i < len(C); i++ {
		if C[i] < 0 {
			C[i] += int64(mod) // Fix negative results
		}
	}

	result := make([]int32, len(C))
	for i := 0; i < len(C); i++ {
		result[i] = int32(C[i])
	}

	return result
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
	//a := []int64{1, 2, 3, 4} // Represents 1 + 2x + 3x^2 + 4x^3
	//b := []int64{2, 3, 4, 5} // Represents 2 + 3x + 4x^2 + 5x^3
	// X=(4,1,4,2,1,3,5,6) and Y=(6,1,8,0,3,3,9,8)
	a := []int64{4, 1, 4, 2, 1, 3, 5, 6}
	b := []int64{6, 1, 8, 0, 3, 3, 9, 8}

	//// Choose modulus M = 2^m + 1 (for m = 32, M = 4294967297)
	//mod := int64(7681)
	//mod := int64(11)
	//mod := int64(4294967297)

	// Multiply using Schönhage–Strassen
	result := schoenhageStrassenMultiply(a, b)

	fmt.Println("Multiplication Result:", result)
}
