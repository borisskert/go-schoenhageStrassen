package main

import "math/bits"

// Modular exponentiation: (base^exp) % mod
func modExp(base, exp, mod int64) int64 {
	if exp == 0 {
		return 1
	}

	if base == 0 {
		return 0
	}

	if base == 1 {
		return 1
	}

	m := uint64(mod)
	result := uint64(1)
	b := uint64(base % mod)

	for exp > 0 {
		if exp%2 == 1 {
			result = (result * b) % m
		}

		hi, lo := bits.Mul64(b, b)
		_, _, b = divmod64(hi, lo, m)

		exp /= 2
	}

	return int64(result)
}

func divmod64(aHi, aLo uint64, b uint64) (uint64, uint64, uint64) {
	if uint64(b) == 0 {
		panic("division by zero")
	}

	// Step 1: Divide the high 64-bit part
	qHi, r := bits.Div64(0, aHi, b)

	// Step 2: Divide the lower 64-bit part using the remainder from step 1
	qLo, r := bits.Div64(r, aLo, b)

	return qHi, qLo, r
}
