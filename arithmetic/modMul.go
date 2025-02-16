package arithmetic

import "math/bits"

func ModMul(a, b, mod uint64) uint64 {
	if a == 0 || b == 0 || mod == a || mod == b {
		return 0
	}

	hi, lo := bits.Mul64(a%mod, b)
	_, _, r := divmod64(hi, lo, mod)

	return r
}
