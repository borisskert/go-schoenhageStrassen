package main

// Modular exponentiation: (base^exp) % mod
func modExp(base, exp, mod int64) int64 {
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
		b = (b * b) % m
		exp /= 2
	}

	return int64(result)
}
