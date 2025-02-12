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

func BitReverseCopyN(a []int32, n int64) []int64 {
	result := make([]int64, n)
	for i := int64(0); i < n; i++ {
		if i < int64(len(a)) {
			rev := ReverseBits(i, n)
			result[rev] = int64(a[i])
		}
	}
	return result
}

// Reverse the bits of i considering the length n (power of 2)
func ReverseBits(i, n int64) int64 {
	rev := int64(0)
	bitCount := CountBits(n) - 1 // Avoid negative shifts
	for j := int64(0); j <= bitCount; j++ {
		rev |= ((i >> j) & 1) << (bitCount - j)
	}
	return rev
}

// Counts bits needed for `n`
func CountBits(n int64) int64 {
	count := int64(0)
	for n > 1 {
		count++
		n >>= 1
	}
	return count
}
