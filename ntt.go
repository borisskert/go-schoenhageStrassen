package main

import "fmt"

func NTT(A []uint64, omega uint64, M uint64) []uint64 {
	fmt.Println("NTT input:", A)

	n := nextPowerOf2(uint64(len(A)))
	result := bitReverseCopyN(A, n)

	for length := uint64(2); length <= n; length *= 2 {
		wLen := modExp(omega, n/length, M)
		for i := uint64(0); i < n; i += length {
			w := uint64(1)
			for j := uint64(0); j < length/2; j++ {
				u := result[i+j]
				v := modMul(result[i+j+length/2], w, M)
				result[i+j] = modAdd(u, v, M)
				result[i+j+length/2] = modSub(u, v, M)
				w = modMul(w, wLen, M)
			}
		}
	}

	for i := uint64(0); i < n; i++ {
		result[i] = result[i] % M

		if result[i] < 0 {
			result[i] += M
		}
	}

	fmt.Println("NTT output:", result)

	return result
}

func NTT16(a []uint16, omega, modulus uint64) []uint64 {
	fmt.Println("NTT input:", a)

	n := nextPowerOf2(uint64(len(a)))
	result := make([]uint64, n)
	for i, v := range a {
		result[i] = uint64(v)
	}

	for length := uint64(1); length < n; length <<= 1 {
		wLen := modExp(omega, (modulus-1)/(length<<1), modulus)
		for i := uint64(0); i < n; i += length << 1 {
			w := uint64(1)
			for j := uint64(0); j < length; j++ {
				u := result[i+j]
				v := (result[i+j+length] * w) % modulus
				result[i+j] = (u + v) % modulus
				result[i+j+length] = (u + modulus - v) % modulus
				w = (w * wLen) % modulus
			}
		}
	}

	fmt.Println("NTT output:", result)

	return result
}
