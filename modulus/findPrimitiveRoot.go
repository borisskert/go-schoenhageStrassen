package modulus

import (
	"errors"
	. "go-schoenhageStrassen/arithmetic"
	"go-schoenhageStrassen/integer"
)

// // Fast modular exponentiation: (base^exp) % mod
func findPrimitiveRoot(N uint64) (*uint64, error) {
	phi := N - 1 // Euler's totient function for prime N is N-1

	// Find prime factors of phi(N)
	factors := integer.Factorize(phi)

	// Try different values for g
	for g := uint64(2); g < N; g++ {
		valid := true
		for _, f := range factors {
			if ModExp(g, phi/f, N) == 1 {
				valid = false
				break
			}
		}
		if valid {
			return &g, nil
		}
	}

	return nil, errors.New("Primitive root not found")
}
