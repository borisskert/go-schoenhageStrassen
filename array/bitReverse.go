package array

// bitReverse reverses the bits of x with respect to the highest power of two less than n.
func bitReverse(x, n int) int {
	result := 0
	for n > 1 {
		result = (result << 1) | (x & 1)
		x >>= 1
		n >>= 1
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

func bitwiseReverseArray16(a []uint16) []uint16 {
	n := len(a)
	result := make([]uint16, n)
	for i := 0; i < n; i++ {
		result[i] = bitwiseReverse16(a[n-i-1])
	}
	return result
}

func bitwiseReverse16(x uint16) uint16 {
	x = (x&0x5555)<<1 | (x&0xAAAA)>>1
	x = (x&0x3333)<<2 | (x&0xCCCC)>>2
	x = (x&0x0F0F)<<4 | (x&0xF0F0)>>4
	x = (x&0x00FF)<<8 | (x&0xFF00)>>8

	return x
}
