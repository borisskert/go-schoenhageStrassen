package arithmetic

import "math/bits"

func ModAdd(a, b, mod uint64) uint64 {
	sum, _ := bits.Add64(a%mod, b%mod, 0)
	return sum % mod
}
