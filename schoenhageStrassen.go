package main

import "fmt"

// Schönhage–Strassen Multiplication
func schoenhageStrassenMultiply(A, B []int32) []int32 {
	// Step 1: Determine the required length
	n := len(A)
	m := len(B)
	requiredLength := n + m - 1

	// Find the smallest power of 2 >= requiredLength
	paddedLength := 1
	for paddedLength < requiredLength {
		paddedLength <<= 1
	}

	// Step 2: Zero-pad the inputs
	A_padded := make([]int32, paddedLength)
	B_padded := make([]int32, paddedLength)
	copy(A_padded, A)
	copy(B_padded, B)

	// Step 3: Find modulus and roots of unity for the padded length
	mod, omega, omegaInv, err := findModulus32xTwo(A_padded, B_padded)
	if err != nil {
		panic(err)
	}

	// Step 4: Apply NTT to both padded inputs
	A_ntt := NTT(A_padded, int64(omega), int64(mod))
	B_ntt := NTT(B_padded, int64(omega), int64(mod))

	// Step 5: Pointwise multiplication
	C_ntt := pointwiseMultiply(A_ntt, B_ntt, int64(mod))

	// Step 6: Apply INTT to get the final coefficients
	C := INTT(C_ntt, int64(omegaInv), int64(mod))

	// Step 7: Trim leading zeros (optional)
	for len(C) > 0 && C[len(C)-1] == 0 {
		C = C[:len(C)-1]
	}

	return C
}

func propagateCarry(C []int32, base int32) []int32 {
	carry := int32(0)
	for i := 0; i < len(C); i++ {
		C[i] += carry
		carry = C[i] / base
		C[i] %= base
	}
	// If there's a carry left, append it to the result
	if carry > 0 {
		C = append(C, carry)
	}
	return C
}

// Pointwise multiplication of two ntt-transformed polynomials
func pointwiseMultiply(A, B []int64, mod int64) []int64 {
	n := len(A)
	C := make([]int64, n)
	for i := 0; i < n; i++ {
		C[i] = modMul(A[i], B[i], mod)
	}
	return C
}

func multiplyExample() {
	// Example numbers (as base 2^m chunks)
	a := []int32{0, 0, 0, 1} // Represents 1 + 2x + 3x^2 + 4x^3
	b := []int32{0, 0, 0, 1} // Represents 2 + 3x + 4x^2 + 5x^3
	fmt.Println("A:", a)
	fmt.Println("B:", b)

	result := schoenhageStrassenMultiply(a, b)

	fmt.Println("Multiplication Result:", result)
}
