package ntt

import (
	"fmt"
	. "go-schoenhageStrassen/arithmetic"
)

// NttNaive computes the Number Theoretic Transform of the input slice a.
func NttNaive(a []uint16, omega uint64, mod uint64) []uint64 {
	fmt.Println("NTT input:", a)

	n := len(a)

	result := make([]uint64, n)

	// Iterative NTT using butterfly operations
	for i := uint64(0); i < uint64(n); i += 1 {
		yi := uint64(0)

		for j := uint64(0); j < uint64(n); j += 1 {
			u := ModExp(omega, i*j, mod)
			v := ModMul(uint64(a[j]), u, mod)
			yi = ModAdd(yi, v, mod)
		}

		result[i] = yi
	}

	fmt.Println("NTT output:", result)

	return result
}
