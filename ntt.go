package main

// Modular exponentiation: (base^exp) % mod
func modExp(base, exp, mod int64) int64 {
	result := uint64(1)

	m := uint64(mod)
	base64 := uint64(base)

	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base64) % m
		}
		base64 = (base64 * base64) % m
		exp /= 2
	}

	return int64(result)
}

// Number-Theoretic Transform (NTT) over Z/(2^m + 1)
func NTT(a []int64, omega, mod int64) []int64 {
	n := len(a)
	if n == 1 {
		return a
	}

	// Split into even and odd parts
	even := make([]int64, n/2)
	odd := make([]int64, n/2)
	for i := 0; i < n/2; i++ {
		even[i] = a[i*2]
		odd[i] = a[i*2+1]
	}

	m := uint64(mod)
	omega2 := (uint64(omega) * uint64(omega)) % m

	// Recursive calls
	fftEven := NTT(even, int64(omega2), mod)
	fftOdd := NTT(odd, int64(omega2), mod)

	// Combine step (Butterfly operation)
	result := make([]int64, n)
	w := uint64(1)

	for k := 0; k < n/2; k++ {
		t := (w * uint64(fftOdd[k])) % m
		result[k] = int64((uint64(fftEven[k]) + t) % m)
		result[k+n/2] = int64((uint64(fftEven[k]) - t + m) % m)
		w = (w * uint64(omega)) % m
	}

	return result
}

// Inverse NTT (INTT)
func INTT(a []int64, omega, mod int64) []int64 {
	n := int64(len(a))
	omegaInv := modExp(omega, mod-2, mod) // Modular inverse of ω
	result := NTT(a, omegaInv, mod)

	// Compute modular inverse of n
	nInv := modExp(n, mod-2, mod)

	// Normalize
	for i := range result {
		result[i] = (result[i] * nInv) % mod
	}
	return result
}
