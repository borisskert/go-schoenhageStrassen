package main

func BitReverseCopy(a []uint64) []uint64 {
	n := uint64(len(a))
	result := make([]uint64, n)
	for i := uint64(0); i < n; i++ {
		rev := ReverseBits(i, n)
		result[rev] = a[i]
	}
	return result
}

func bitReverseCopy32N(a []uint32, n uint64) []uint64 {
	result := make([]uint64, n)
	for i := uint64(0); i < n; i++ {
		rev := ReverseBits(i, n)

		if i < uint64(len(a)) {
			result[rev] = uint64(a[i])
		}
	}
	return result
}

func bitReverseCopyN(a []uint64, n uint64) []uint64 {
	result := make([]uint64, n)
	for i := uint64(0); i < n; i++ {
		rev := ReverseBits(i, n)

		if i < uint64(len(a)) {
			result[rev] = uint64(a[i])
		}
	}
	return result
}

func bitReverseCopy16N(a []uint16, n uint64) []uint64 {
	result := make([]uint64, n)
	for i := uint64(0); i < n; i++ {
		rev := ReverseBits(i, n)

		if i < uint64(len(a)) {
			result[rev] = uint64(a[i])
		}
	}
	return result
}

// Reverse the bits of i considering n (which should be a power of 2)
func ReverseBits(i, n uint64) uint64 {
	rev := uint64(0)
	bitCount := CountBits(n) - 1 // Ensure correct bit shifting
	for j := uint64(0); j < bitCount; j++ {
		rev |= ((i >> j) & 1) << (bitCount - 1 - j)
	}
	return rev
}

func ReverseBits16(a []uint16) []uint16 {
	n := uint64(len(a))
	result := make([]uint16, n)
	for i := uint64(0); i < n; i++ {
		result[i] = a[n-1-i]
	}
	return result
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
