package main

import "fmt"

func INTT(a []int64, omegaInv int64, M int64) []int64 {
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
				result[i+j+length/2] = modSub(u, v, M) // ✅ FIXED
				w = modMul(w, wLen, M)
			}
		}
	}

	nInv := modInverseFermat(n, M) // Ensure M is prime

	fmt.Println("nInv:", nInv, " for n:", n)

	if modMul(n, nInv, M) != 1 {
		fmt.Println("ERROR: n * nInv % M != 1")
		return nil
	}

	fmt.Println("n * nInv % M:", modMul(n, nInv, M))

	if modMul(n, nInv, M) != 1 {
		fmt.Println("ERROR: n * nInv % M != 1")
		return nil
	}

	for i := 0; i < len(result); i++ {
		result[i] = modMul(result[i], nInv, M)
		if result[i] < 0 {
			result[i] += M // ✅ Fix negative values
		}
	}

	fmt.Println("INTT Result:", result)

	return result
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
