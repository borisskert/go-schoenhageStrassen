package main

import "fmt"

// NTT computes the Number Theoretic Transform of the input vector.
func NTT2(a []uint16, omega, modulus uint64) []uint64 {
	fmt.Println("NTT input:", a)

	//n := len(a)
	//if n == 0 {
	//	return nil
	//}
	n := nextPowerOf2a(len(a))
	//n := len(a)

	// Convert input to uint64
	a64 := make([]uint64, n)
	for i := range a {
		a64[i] = uint64(a[i])
	}

	// Bit-reverse permutation
	rev := bitReverseCopyN(a64, uint64(n))

	// Cooley-Tukey NTT
	for s := 1; s < n; s <<= 1 {
		m := s << 1
		wm := powMod(omega, uint64(n/m), modulus)

		for k := 0; k < n; k += m {
			w := uint64(1)
			for j := 0; j < s; j++ {
				//t := w * rev[k+j+s] % modulus
				t := modMul(w, rev[k+j+s], modulus)
				u := rev[k+j]
				//rev[k+j] = (u + t) % modulus
				rev[k+j] = modAdd(u, t, modulus)
				//rev[k+j+s] = (u - t + modulus) % modulus
				rev[k+j+s] = modSub(u, t, modulus)
				//w = w * wm % modulus
				w = modMul(w, wm, modulus)
			}
		}
	}

	fmt.Println("NTT output:", rev)

	return rev
}

// INTT computes the Inverse Number Theoretic Transform of the input vector.
func INTT2(a []uint64, omega, modulus uint64) []uint16 {
	fmt.Println("INTT input:", a)

	n := len(a)
	if n == 0 {
		return nil
	}

	// Compute the inverse of n modulo `modulus`
	nInv := invMod(uint64(n), modulus)

	// Compute the inverse of omega modulo `modulus`
	omegaInv := invMod(omega, modulus)

	// Perform NTT with omegaInv
	inttResult := nttWithOmega(a, omegaInv, modulus)

	// Multiply by nInv and convert to uint16
	result := make([]uint16, n)
	for i := range inttResult {
		result[i] = uint16((inttResult[i] * nInv) % modulus)
	}

	fmt.Println("INTT output:", result)

	return result
}

func INTT2a(a []uint64, omegaInv, modulus uint64) []uint16 {
	fmt.Println("INTT input:", a)

	n := len(a)
	if n == 0 {
		return nil
	}

	// Compute the inverse of n modulo `modulus`
	nInv := invMod(uint64(n), modulus)

	// Perform NTT with omegaInv
	inttResult := nttWithOmega(a, omegaInv, modulus)

	// Multiply by nInv and convert to uint16
	result := make([]uint16, n)
	for i := range inttResult {
		result[i] = uint16((inttResult[i] * nInv) % modulus)
	}

	//result = trimZeros16(result)

	fmt.Println("INTT output:", result)

	return result
}

func nttWithOmega(a []uint64, omega, modulus uint64) []uint64 {
	n := len(a)
	if n == 0 {
		return nil
	}

	// Bit-reverse permutation
	rev := bitReverseCopyN(a, uint64(n))

	// Cooley-Tukey NTT
	for s := 1; s < n; s <<= 1 {
		m := s << 1
		wm := powMod(omega, uint64(n/m), modulus)

		for k := 0; k < n; k += m {
			w := uint64(1)
			for j := 0; j < s; j++ {
				//t := w * rev[k+j+s] % modulus
				t := modMul(w, rev[k+j+s], modulus)
				u := rev[k+j]
				//rev[k+j] = (u + t) % modulus
				rev[k+j] = modAdd(u, t, modulus)
				//rev[k+j+s] = (u - t + modulus) % modulus
				rev[k+j+s] = modSub(u, t, modulus)
				//w = w * wm % modulus
				w = modMul(w, wm, modulus)
			}
		}
	}

	return rev
}

// reverseBits returns the bit-reversed value of x with respect to n.
func reverseBits(x, n uint32) uint32 {
	var r uint32
	for i := uint32(0); i < n; i++ {
		r = (r << 1) | (x & 1)
		x >>= 1
	}
	return r
}

// powMod computes (base^exp) % mod efficiently.
func powMod(base, exp, mod uint64) uint64 {
	result := uint64(1)
	base = base % mod
	for exp > 0 {
		if exp&1 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exp >>= 1
	}
	return result
}

// invMod computes the modular inverse of a modulo mod.
func invMod(a, mod uint64) uint64 {
	// Using Fermat's little theorem for prime mod
	return powMod(a, mod-2, mod)
}
