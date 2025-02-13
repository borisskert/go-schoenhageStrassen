package main

import "fmt"

func INTT(a []uint64, omegaInv uint64, M uint64) []uint64 {
	fmt.Println("INTT input:", a)

	n := uint64(len(a))
	a = BitReverseCopy(a)

	result := make([]uint64, len(a))
	copy(result, a)

	for length := uint64(2); length <= n; length *= 2 {
		wLen := modExp(omegaInv, n/length, M)

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

	nInv := modInverseFermat(n, M)

	if modMul(n, nInv, M) != 1 {
		fmt.Println("ERROR: n * nInv % M != 1")
		return nil
	}

	if modMul(n, nInv, M) != 1 {
		fmt.Println("ERROR: n * nInv % M != 1")
		return nil
	}

	result64 := make([]uint64, len(result))
	size := len(result)
	shrinking := true
	for i := len(result) - 1; i >= 0; i-- {
		result64[i] = uint64(modMul(result[i], nInv, M))

		if shrinking && result64[i] == 0 {
			size--
		} else {
			shrinking = false
		}
	}

	fmt.Println("INTT output:", result64)

	return result64[:size]
}

func INTT16(a []uint64, omegaInv uint64, modulus uint64) []uint16 {
	fmt.Println("INTT input:", a)

	n := uint64(len(a))
	result := make([]uint64, n)
	copy(result, a)

	//invOmega := modExp(omega, modulus-2, modulus)
	for length := n >> 1; length > 0; length >>= 1 {
		wLen := modExp(omegaInv, (modulus-1)/(length<<1), modulus)
		for i := uint64(0); i < n; i += length << 1 {
			w := uint64(1)
			for j := uint64(0); j < length; j++ {
				u := result[i+j]
				v := result[i+j+length]
				//result[i+j] = (u + v) % modulus
				result[i+j] = modAdd(u, v, modulus)
				//result[i+j+length] = ((u + modulus - v) * w) % modulus
				result[i+j+length] = modMul(modSub(u, v, modulus), w, modulus)
				//w = (w * wLen) % modulus
				w = modMul(w, wLen, modulus)
			}
		}
	}

	//invN := modExp(n, modulus-2, modulus)
	invN := modInverseFermat(n, modulus)

	if modMul(n, invN, modulus) != 1 {
		fmt.Println("ERROR: n * invN % modulus != 1")
		return nil
	}

	output := make([]uint16, n)
	for i := range result {
		//output[i] = uint16((result[i] * invN) % modulus)
		output[i] = uint16(modMul(result[i], invN, modulus))
	}

	//output = trimZeros16(output)

	fmt.Println("INTT output:", output)

	return output
}
