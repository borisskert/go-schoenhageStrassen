package main

import "fmt"

func INTT(a []int64, omegaInv int64, M int64) []int32 {
	n := int64(len(a))

	fmt.Println("INTT input:", a)

	a = BitReverseCopy(a) // Ensure bit-reversed order

	result := make([]int64, len(a))
	copy(result, a)

	for length := int64(2); length <= n; length *= 2 {
		wLen := modExp(omegaInv, n/length, M)

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

	nInv := modInverseFermat(n, M)

	if modMul(n, nInv, M) != 1 {
		fmt.Println("ERROR: n * nInv % M != 1")
		return nil
	}

	if modMul(n, nInv, M) != 1 {
		fmt.Println("ERROR: n * nInv % M != 1")
		return nil
	}

	result32 := make([]int32, len(result))
	size := len(result)
	shrinking := true
	for i := len(result) - 1; i >= 0; i-- {
		result32[i] = int32(modMul(result[i], nInv, M))

		if shrinking && result32[i] == 0 {
			size--
		} else {
			shrinking = false
		}
	}

	fmt.Println("INTT output:", result32)

	return result32[:size]
}

func intt(a []int64, omegaInv int64, M int64) []int64 {
	n := int64(len(a))

	a = BitReverseCopy(a) // Ensure bit-reversed order
	result := make([]int64, len(a))
	copy(result, a)

	for length := int64(2); length <= n; length *= 2 {
		wLen := modExp(omegaInv, n/length, M)
		for i := int64(0); i < n; i += length {
			w := int64(1)
			for j := int64(0); j < length/2; j++ {
				u := result[i+j]
				v := modMul(result[i+j+length/2], w, M)
				result[i+j] = modAdd(u, v, M)
				result[i+j+length/2] = modAdd(modSub(u, v, M), M, M)
				w = modMul(w, wLen, M)
			}
		}
	}

	nInv := modExp(n, M-2, M) // Modular Inverse of n
	for i := range result {
		result[i] = (result[i] * nInv) % M
	}

	for i := len(result) - 1; i >= 1; i-- {
		if result[i] == 0 {
			result = result[:i]
		} else {
			break
		}
	}

	return result
}
