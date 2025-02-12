package main

import "math/bits"

func modAdd(a, b, mod int64) int64 {
	sum, _ := bits.Add64(uint64(a%mod), uint64(b%mod), 0)
	return normalizeMod(int64(sum%uint64(mod)), mod)
}
