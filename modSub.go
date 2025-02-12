package main

import "math/bits"

func modSub(a, b, mod int64) int64 {
	diff, _ := bits.Sub64(uint64(a%mod), uint64(b%mod), 0)
	return normalizeMod(int64(diff), mod)
}
