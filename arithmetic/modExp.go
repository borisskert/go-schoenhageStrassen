package arithmetic

import "math/bits"

// ModExp Modular exponentiation: (base^exp) % mod
func ModExp(base, exp, mod uint64) uint64 {
	if mod == 1 {
		return 0 // Special case: Any number mod 1 is always 0
	}

	result := uint64(1)
	b := base % mod // Ensure base is within [0, mod)
	m := mod

	for exp > 0 {
		if exp&1 == 1 { // If exp is odd, multiply result
			result = safeModMul(result, b, m)
		}
		b = safeModMul(b, b, m) // Square the base safely
		exp >>= 1
	}

	return result
}

// safeModMul Safe modular multiplication using 128-bit arithmetic
func safeModMul(a uint64, b uint64, mod uint64) uint64 {
	hi, lo := bits.Mul64(uint64(a), b) // Perform 128-bit multiplication
	_, _, r := divmod64(hi, lo, mod)
	return r
}
