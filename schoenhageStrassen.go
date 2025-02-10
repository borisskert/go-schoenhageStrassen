package main

// Schönhage–Strassen Multiplication
func schoenhageStrassenMultiply(a, b []int64, mod int64) []int64 {
	n := len(a)
	g, err := findPrimitiveRoot(mod)
	if err != nil {
		panic(err)
	}

	omega := modExp(g, (mod-1)/int64(n), mod) // Primitive root
	//omega := int64(3383)

	// Step 1: Forward NTT
	A := NTT(a, omega, mod)
	B := NTT(b, omega, mod)

	// Step 2: Pointwise multiplication
	C := pointwiseMultiplication(A, B, mod)

	// Step 3: Inverse NTT
	result := INTT(C, omega, mod)
	return result
}

func pointwiseMultiplication(a, b []int64, mod int64) []int64 {
	n := len(a)

	m := uint64(mod)

	result := make([]int64, n)

	for i := 0; i < n; i++ {
		result[i] = int64(uint64(a[i]) * uint64(b[i]) % m)
	}

	return result
}
