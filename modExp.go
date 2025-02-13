package main

import "math/bits"

// Modular exponentiation: (base^exp) % mod
func modExp(base, exp, mod int64) int64 {
	if mod == 1 {
		return 0 // Special case: Any number mod 1 is always 0
	}

	result := int64(1)
	b := uint64(base % mod) // Ensure base is within [0, mod)
	m := uint64(mod)

	for exp > 0 {
		if exp&1 == 1 { // If exp is odd, multiply result
			result = int64(safeModMul(result, b, m))
		}
		b = safeModMul(int64(b), b, m) // Square the base safely
		exp >>= 1
	}

	return normalizeMod(result, mod) // Ensure result is non-negative
}

// Safe modular multiplication using 128-bit arithmetic
func safeModMul(a int64, b uint64, mod uint64) uint64 {
	hi, lo := bits.Mul64(uint64(a), b) // Perform 128-bit multiplication
	_, _, r := divmod64(hi, lo, mod)
	return uint64(normalizeMod(int64(r), int64(mod)))
}

// Ensure non-negative modulo result
func normalizeMod(x, mod int64) int64 {
	x %= mod

	for x < 0 {
		x += mod
	}

	return x
}
