package main

import "fmt"

func INTT(a []int64, omegaInv int64, M int64) []int32 {
	n := int64(len(a))

	fmt.Println("Before BitReverseCopy (INTT):", a)
	a = BitReverseCopy(a) // Ensure bit-reversed order
	fmt.Println("After BitReverseCopy (INTT):", a)

	result := make([]int64, len(a))
	copy(result, a)

	for length := int64(2); length <= n; length *= 2 {
		wLen := modExp(omegaInv, n/length, M)
		fmt.Println("wLen:", wLen, "for length:", length)

		for i := int64(0); i < n; i += length {
			w := int64(1)
			for j := int64(0); j < length/2; j++ {
				u := result[i+j]
				v := modMul(result[i+j+length/2], w, M)

				fmt.Printf("Step: i=%d, j=%d, u=%d, v=%d, w=%d\n", i, j, u, v, w)

				fmt.Println("modAdd debug:", u, v, modAdd(u, v, M))
				fmt.Println("modSub debug:", u, v, modSub(u, v, M))

				result[i+j] = modAdd(u, v, M)
				result[i+j+length/2] = modSub(u, v, M)

				w = modMul(w, wLen, M)

				// DEBUG: Print the w progression
				fmt.Printf("New w: %d\n", w)
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

	result32 := make([]int32, len(result))
	size := len(result)
	shrinking := true
	for i := len(result) - 1; i >= 0; i-- {
		fmt.Println("Before nInv Mul:", result[i])
		result32[i] = int32(modMul(result[i], nInv, M))
		fmt.Println("After nInv Mul:", result[i])

		if shrinking && result32[i] == 0 {
			size--
		} else {
			shrinking = false
		}
	}

	fmt.Println("INTT Result:", result32)

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
