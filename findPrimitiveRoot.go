package main

// // Fast modular exponentiation: (base^exp) % mod
func findPrimitiveRoot(N int64) int64 {
	phi := N - 1 // Euler's totient function for prime N is N-1

	// Find prime factors of phi(N)
	factors := factorize(phi)

	// Try different values for g
	for g := int64(2); g < N; g++ {
		valid := true
		for _, f := range factors {
			if modExp(g, phi/f, N) == 1 {
				valid = false
				break
			}
		}
		if valid {
			return g
		}
	}

	return -1 // Should never reach here if N is prime
}
