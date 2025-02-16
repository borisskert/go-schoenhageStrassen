package array

func BitwiseReverse64(a []uint64) []uint64 {
	n := len(a)
	result := make([]uint64, n)

	bits := uint64(0)
	for length := n; length > 1; length >>= 1 {
		bits++
	}

	for i := uint64(0); i < uint64(n); i++ {
		rev := reverseBits(i, bits)
		result[rev] = a[i]
	}

	return result
}

func reverseBits(x uint64, bits uint64) uint64 {
	var r uint64 = 0
	for i := uint64(0); i < bits; i++ {
		r = (r << 1) | (x & 1)
		x >>= 1
	}
	return r
}

func Reverse64(a []uint64) []uint64 {
	n := len(a)
	result := make([]uint64, n)

	for i := 0; i < n; i++ {
		result[i] = a[n-i-1]
	}

	return result
}

// bitReverse performs bit-reversal permutation on the input slice.
func bitwiseReverseInPlace64(a []uint64) {
	n := len(a)
	j := 0
	for i := 0; i < n; i++ {
		bit := n >> 1
		for ; j&bit != 0; bit >>= 1 {
			j ^= bit
		}
		j ^= bit
		if i < j {
			a[i], a[j] = bitwiseReverse64(a[j]), bitwiseReverse64(a[i])
		}
	}
}

func bitwiseReverse64(x uint64) uint64 {
	x = (x&0x5555555555555555)<<1 | (x&0xAAAAAAAAAAAAAAAA)>>1
	x = (x&0x3333333333333333)<<2 | (x&0xCCCCCCCCCCCCCCCC)>>2
	x = (x&0x0F0F0F0F0F0F0F0F)<<4 | (x&0xF0F0F0F0F0F0F0F0)>>4
	x = (x&0x00FF00FF00FF00FF)<<8 | (x&0xFF00FF00FF00FF00)>>8

	return x
}

func BitReverseCopyN(a []uint64, n uint64) []uint64 {
	result := make([]uint64, n)
	for i := uint64(0); i < n; i++ {
		rev := ReverseBits(i, n)

		if i < uint64(len(a)) {
			result[rev] = a[i]
		}
	}
	return result
}

// ReverseBits Reverse the bits of i considering n (which should be a power of 2)
func ReverseBits(i, n uint64) uint64 {
	rev := uint64(0)
	bitCount := CountBits(n) - 1 // Ensure correct bit shifting
	for j := uint64(0); j < bitCount; j++ {
		rev |= ((i >> j) & 1) << (bitCount - 1 - j)
	}
	return rev
}

// Counts the number of bits in `n`
func CountBits(n uint64) uint64 {
	count := uint64(0)
	for n > 1 {
		count++
		n >>= 1
	}
	return count + 1 // Ensure correct count
}
