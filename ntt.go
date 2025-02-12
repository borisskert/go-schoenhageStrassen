package main

import "fmt"

func NTT(A []int32, omega int64, M int64) []int64 {
	fmt.Println("NTT input:", A)

	n := NextPowerOf2(int64(len(A)))

	result := BitReverseCopyN(A, n) // Ensure bit-reversed order

	for length := int64(2); length <= n; length *= 2 {
		wLen := modExp(omega, n/length, M)
		for i := int64(0); i < n; i += length {
			w := int64(1)
			for j := int64(0); j < length/2; j++ {
				u := result[i+j]
				v := modMul(result[i+j+length/2], w, M)
				result[i+j] = modAdd(u, v, M)
				result[i+j+length/2] = modSub(u, v, M)
				w = modMul(w, wLen, M)
			}
		}
	}

	for i := int64(0); i < n; i++ {
		result[i] = result[i] % M

		if result[i] < 0 {
			result[i] += M
		}
	}

	fmt.Println("NTT output:", result)

	return result
}
