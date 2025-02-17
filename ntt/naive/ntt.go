package naive

import (
	. "github.com/borisskert/go-schoenhageStrassen/arithmetic"
)

// NaiveNTT computes the Number Theoretic Transform of the input slice a.
func nttNaive(a []uint16, omega uint64, mod uint64) []uint64 {
	n := len(a)

	result := make([]uint64, n)

	for i := uint64(0); i < uint64(n); i += 1 {
		yi := uint64(0)

		for j := uint64(0); j < uint64(n); j += 1 {
			u := ModExp(omega, i*j, mod)
			v := ModMul(uint64(a[j]), u, mod)
			yi = ModAdd(yi, v, mod)
		}

		result[i] = yi
	}

	return result
}
