package main

import "fmt"

// Schönhage–Strassen Multiplication
func schoenhageStrassenMultiply32(A, B []uint32) []uint32 {

	fmt.Println("A:", A)
	fmt.Println("B:", B)

	// Step 1: Convert inputs to uint64
	A64 := make([]uint64, len(A))
	B64 := make([]uint64, len(B))
	for i := range A {
		A64[i] = uint64(A[i])
	}
	for i := range B {
		B64[i] = uint64(B[i])
	}

	// Step 2: Determine the required length
	n := len(A64)
	m := len(B64)
	requiredLength := n + m - 1

	// Find the smallest power of 2 >= requiredLength
	paddedLength := 1
	for paddedLength < requiredLength {
		paddedLength <<= 1
	}

	// Step 3: Zero-pad the inputs
	A_padded := make([]uint64, paddedLength)
	B_padded := make([]uint64, paddedLength)
	copy(A_padded, A64)
	copy(B_padded, B64)

	// Step 4: Find modulus and roots of unity for the padded length
	mod, omega, omegaInv, err := findModulusTwo(A_padded, B_padded)
	if err != nil {
		panic(err)
	}

	// Step 5: Apply NTT to both padded inputs
	A_ntt := NTT(A_padded, uint64(omega), uint64(mod))
	B_ntt := NTT(B_padded, uint64(omega), uint64(mod))

	// Step 6: Pointwise multiplication
	C_ntt := pointwiseMultiply(A_ntt, B_ntt, uint64(mod))

	// Step 7: Apply INTT to get the final coefficients
	C64 := INTT(C_ntt, uint64(omegaInv), uint64(mod))

	// Step 8: Convert the result to []uint32 and handle carries
	C := make([]uint32, len(C64)*2) // Each uint64 becomes 2 uint32s
	carry := uint64(0)
	for i := range C64 {
		// Add the carry to the current coefficient
		value := C64[i] + carry
		// Split the value into two uint32 coefficients
		C[2*i] = uint32(value & 0xFFFFFFFF)           // Lower 32 bits
		C[2*i+1] = uint32((value >> 32) & 0xFFFFFFFF) // Upper 32 bits
		// Calculate the carry for the next coefficient
		carry = value >> 64
	}

	// Step 9: If there's a carry left, append it to the result
	if carry > 0 {
		C = append(C, uint32(carry))
	}

	// Step 10: Trim leading zeros (optional)
	for len(C) > 0 && C[len(C)-1] == 0 {
		C = C[:len(C)-1]
	}

	fmt.Println("Final result:", C)

	return C
}

// Pointwise multiplication of two ntt-transformed polynomials
func pointwiseMultiply(A, B []uint64, mod uint64) []uint64 {
	n := len(A)
	C := make([]uint64, n)
	for i := 0; i < n; i++ {
		C[i] = modMul(A[i], B[i], mod)
	}
	return C
}

func multiplyExample() {
	// Example numbers (as base 2^m chunks)
	a := []uint32{0, 0, 0, 1} // Represents 1 + 2x + 3x^2 + 4x^3
	b := []uint32{0, 0, 0, 1} // Represents 2 + 3x + 4x^2 + 5x^3
	fmt.Println("A:", a)
	fmt.Println("B:", b)

	result := schoenhageStrassenMultiply32(a, b)

	fmt.Println("Multiplication Result:", result)
}

//func schoenhageStrassenMultiply16(A, B []uint16) []uint16 {
//	fmt.Println("A:", A)
//	fmt.Println("B:", B)
//
//	// Step 1: Determine the required length
//	n := len(A)
//	m := len(B)
//	requiredLength := n + m - 1
//
//	// Find the smallest power of 2 >= requiredLength
//	paddedLength := 1
//	for paddedLength < requiredLength {
//		paddedLength <<= 1
//	}
//
//	// Step 2: Zero-pad the inputs
//	A_padded := make([]uint64, paddedLength)
//	B_padded := make([]uint64, paddedLength)
//	for i := range A {
//		A_padded[i] = uint64(A[i])
//	}
//	for i := range B {
//		B_padded[i] = uint64(B[i])
//	}
//
//	// Step 3: Find modulus and roots of unity for the padded length
//	mod, omega, omegaInv, err := findModulusTwo(A_padded, B_padded)
//	if err != nil {
//		panic(err)
//	}
//
//	// Step 4: Apply NTT to both padded inputs
//	A_ntt := NTT(A_padded, uint64(omega), uint64(mod))
//	B_ntt := NTT(B_padded, uint64(omega), uint64(mod))
//
//	// Step 5: Pointwise multiplication
//	C_ntt := pointwiseMultiply(A_ntt, B_ntt, uint64(mod))
//
//	// Step 6: Apply INTT to get the final coefficients
//	C := INTT16(C_ntt, uint64(omegaInv), uint64(mod))
//
//	// Step 7: Convert the result to []uint16 and handle carries
//	//C32 := make([]uint32, len(C64))
//	//for i := range C64 {
//	//	C32[i] = uint32(C64[i])
//	//}
//	//
//	//C16 := splitUint32ToUint16(C32)
//
//	// Step 9: Trim leading zeros (optional)
//	for len(C) > 0 && C[len(C)-1] == 0 {
//		C = C[:len(C)-1]
//	}
//
//	fmt.Println("C:", C)
//
//	return C
//}

func schoenhageStrassenMultiply16(a, b []uint16) []uint16 {
	fmt.Println("A:", a)
	fmt.Println("B:", b)

	// Ensure both inputs are of the same length
	n := 1
	//requiredN := len(a) + len(b) - 1
	for n < len(a) || n < len(b) {
		//for n < requiredN {
		n <<= 1
	}

	// Pad the inputs with zeros to make them length n
	aPadded := make([]uint16, n)
	bPadded := make([]uint16, n)
	copy(aPadded, a)
	copy(bPadded, b)

	modulus, omega, omegaInv, err := findModulus16Two(aPadded, bPadded)
	if err != nil {
		panic(err)
	}

	// Compute NTT of both inputs
	nttA := NTT2(aPadded, uint64(omega), uint64(modulus))
	nttB := NTT2(bPadded, uint64(omega), uint64(modulus))
	//nttA := NTT16(aPadded, uint64(omega), uint64(modulus))
	//nttB := NTT16(bPadded, uint64(omega), uint64(modulus))

	// Pointwise multiplication in NTT domain
	nttC := make([]uint64, n)
	for i := range nttC {
		//nttC[i] = (nttA[i] * nttB[i]) % modulus
		nttC[i] = modMul(nttA[i], nttB[i], uint64(modulus))
	}

	// Compute INTT to get the result in the original domain
	result := INTT2a(nttC, uint64(omegaInv), uint64(modulus))

	fmt.Println("Result:", result)

	// Normalize the result (remove leading zeros)
	for len(result) > 1 && result[len(result)-1] == 0 {
		result = result[:len(result)-1]
	}

	fmt.Println("Result (normalized):", result)

	return result
}

func schoenhageStrassenMultiply16a(A, B []uint16) []uint16 {
	fmt.Println("A:", A)
	fmt.Println("B:", B)

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
	A_padded := make([]uint16, paddedLength)
	B_padded := make([]uint16, paddedLength)
	for i := range A {
		A_padded[i] = uint16(A[i])
	}
	for i := range B {
		B_padded[i] = uint16(B[i])
	}

	// Step 3: Find modulus and roots of unity for the padded length
	mod, omega, omegaInv, err := findModulus16Two(A_padded, B_padded)
	if err != nil {
		panic(err)
	}

	// Step 4: Apply NTT to both padded inputs
	A_ntt := NTT16(A_padded, uint64(omega), uint64(mod))
	B_ntt := NTT16(B_padded, uint64(omega), uint64(mod))

	// Step 5: Pointwise multiplication
	C_ntt := pointwiseMultiply(A_ntt, B_ntt, uint64(mod))

	// Step 6: Apply INTT to get the final coefficients
	C := INTT16(C_ntt, uint64(omegaInv), uint64(mod))

	// Step 7: Convert the result to []uint16 and handle carries
	//C32 := make([]uint32, len(C64))
	//for i := range C64 {
	//	C32[i] = uint32(C64[i])
	//}
	//
	//C16 := splitUint32ToUint16(C32)

	C = ReverseBits16(C)

	// Step 9: Trim leading zeros (optional)
	for len(C) > 0 && C[len(C)-1] == 0 {
		C = C[:len(C)-1]
	}

	fmt.Println("C:", C)

	return C
}

func schoenhageStrassenMultiply64(A, B []uint64) []uint64 {
	fmt.Println("A:", A)
	fmt.Println("B:", B)

	a16 := splitUint64ToUint16(A)
	b16 := splitUint64ToUint16(B)

	// Step 1: Determine the required length
	n := len(a16)
	m := len(b16)
	requiredLength := n + m - 1

	// Find the smallest power of 2 >= requiredLength
	paddedLength := 1
	for paddedLength < requiredLength {
		paddedLength <<= 1
	}

	// Step 2: Zero-pad the inputs
	A_padded := copy16to64N(a16, paddedLength)
	B_padded := copy16to64N(b16, paddedLength)

	// Step 3: Find modulus and roots of unity for the padded length
	mod, omega, omegaInv, err := findModulusTwo(A_padded, B_padded)
	if err != nil {
		panic(err)
	}

	// Step 4: Apply NTT to both padded inputs
	A_ntt := NTT(A_padded, uint64(omega), uint64(mod))
	B_ntt := NTT(B_padded, uint64(omega), uint64(mod))

	// Step 5: Pointwise multiplication
	C_ntt := pointwiseMultiply(A_ntt, B_ntt, uint64(mod))

	// Step 6: Apply INTT to get the final coefficients
	C := INTT(C_ntt, uint64(omegaInv), uint64(mod))

	// Step 7: Convert the result to []uint16 and handle carries
	//C := make([]uint16, len(C64)*4) // Each uint64 becomes 4 uint16s
	//carry := uint64(0)
	//for i := range C64 {
	//	// Add the carry to the current coefficient
	//	value := C64[i] + carry
	//	// Split the value into four uint16 coefficients
	//	C[4*i] = uint16(value & 0xFFFF)           // Least significant 16 bits
	//	C[4*i+1] = uint16((value >> 16) & 0xFFFF) // Next 16 bits
	//	C[4*i+2] = uint16((value >> 32) & 0xFFFF) // Next 16 bits
	//	C[4*i+3] = uint16((value >> 48) & 0xFFFF) // Most significant 16 bits
	//	// Calculate the carry for the next coefficient
	//	carry = value >> 64
	//}
	//
	//// Step 8: If there's a carry left, append it to the result
	//if carry > 0 {
	//	C = append(C, uint16(carry))
	//}

	// Step 9: Trim leading zeros (optional)
	for len(C) > 0 && C[len(C)-1] == 0 {
		C = C[:len(C)-1]
	}

	fmt.Println("Final result:", C)

	return C
}
