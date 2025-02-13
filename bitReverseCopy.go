package main

func BitReverseCopy(a []int64) []int64 {
	n := int64(len(a))
	result := make([]int64, n)
	for i := int64(0); i < n; i++ {
		rev := ReverseBits(i, n)
		result[rev] = a[i]
	}
	return result
}

func bitReverseCopyN(a []int32, n int64) []int64 {
	result := make([]int64, n)
	for i := int64(0); i < n; i++ {
		rev := ReverseBits(i, n)

		if i < int64(len(a)) {
			result[rev] = int64(a[i])
		}
	}
	return result
}

// Reverse the bits of i considering n (which should be a power of 2)
func ReverseBits(i, n int64) int64 {
	rev := int64(0)
	bitCount := CountBits(n) - 1 // Ensure correct bit shifting
	for j := int64(0); j < bitCount; j++ {
		rev |= ((i >> j) & 1) << (bitCount - 1 - j)
	}
	return rev
}

// Counts the number of bits in `n`
func CountBits(n int64) int64 {
	count := int64(0)
	for n > 1 {
		count++
		n >>= 1
	}
	return count + 1 // Ensure correct count
}
