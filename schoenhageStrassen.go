package main

import "fmt"

// Schönhage–Strassen Multiplication
func schoenhageStrassenMultiply(A, B []int64) []int32 {
	n := int64(len(A))
	n = NextPowerOf2(n) // Ensure n is a power of 2

	A = append(A, make([]int64, n-int64(len(A)))...) // Pad with zeros
	B = append(B, make([]int64, n-int64(len(B)))...) // Pad with zeros

	mod, _, omega, omegaInv, err := findModulus2(A, B)
	if err != nil {
		panic(err)
	}

	// Step 1: Apply NTT to both inputs
	A_ntt := NTT(A, int64(omega), int64(mod))
	fmt.Println("A_ntt:", A_ntt)

	B_ntt := NTT(B, int64(omega), int64(mod))
	fmt.Println("B_ntt:", B_ntt)

	// Step 2: Pointwise multiplication
	C_ntt := pointwiseMultiply(A_ntt, B_ntt, int64(mod))
	fmt.Println("C_ntt:", C_ntt)

	// Step 3: Apply INTT to get the final coefficients
	C := INTT(C_ntt, int64(omegaInv), int64(mod))
	fmt.Println("C:", C)

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
	a := []int64{1, 2, 3, 4} // Represents 1 + 2x + 3x^2 + 4x^3
	b := []int64{2, 3, 4, 5} // Represents 2 + 3x + 4x^2 + 5x^3
	fmt.Println("A:", a)
	fmt.Println("B:", b)

	result := schoenhageStrassenMultiply(a, b)

	fmt.Println("Multiplication Result:", result)
}
