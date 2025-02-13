package arithmetic

import "math/bits"

func ModAdd(a, b, mod uint64) uint64 {
	sum, _ := bits.Add64(uint64(a%mod), uint64(b%mod), 0)
	return normalizeMod(uint64(sum%uint64(mod)), mod)
}
