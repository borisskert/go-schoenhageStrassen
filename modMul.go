package main

import "math/bits"

func modMul(a, b, mod int64) int64 {
	if a == 0 || b == 0 || mod == a || mod == b {
		return 0
	}

	hi, lo := bits.Mul64(uint64(a%mod), uint64(b))

	_, _, r := divmod64(hi, lo, uint64(mod))

	return normalizeMod(int64(r), mod)
}
