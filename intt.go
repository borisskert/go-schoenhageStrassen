package main

func INTT(a []int64, n int64, omegaInv int64, M int64) []int64 {
	a = BitReverseCopy(a) // Ensure bit-reversed order
	result := make([]int64, len(a))
	copy(result, a)

	for length := int64(2); length <= n; length *= 2 {
		wLen := modExp(omegaInv, n/length, M)
		for i := int64(0); i < n; i += length {
			w := int64(1)
			for j := int64(0); j < length/2; j++ {
				u := result[i+j]
				v := (result[i+j+length/2] * w) % M
				result[i+j] = (u + v) % M
				result[i+j+length/2] = (u - v + M) % M
				w = (w * wLen) % M
			}
		}
	}

	nInv := modExp(n, M-2, M) // Modular Inverse of n
	for i := range result {
		result[i] = (result[i] * nInv) % M
	}
	return result
}
