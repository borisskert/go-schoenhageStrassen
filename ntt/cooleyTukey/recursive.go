package cooleyTukey

import (
	. "github.com/borisskert/go-schoenhageStrassen/arithmetic"
	"github.com/borisskert/go-schoenhageStrassen/misc"
	"github.com/borisskert/go-schoenhageStrassen/ntt"
)

type cooleyTukeyRecursive struct{}

func (c cooleyTukeyRecursive) NTT(a []uint16, omega uint64, mod uint64) []uint64 {
	a64 := make([]uint64, len(a))
	for i := range a {
		a64[i] = uint64(a[i])
	}

	return cooleyTukeyRecursiveNTT(a64, omega, mod)
}

func (c cooleyTukeyRecursive) INTT(a []uint64, omega uint64, mod uint64) []uint16 {
	return cooleyTukeyINTT(a, omega, mod)
}

func RecursiveAlgorithm() ntt.NumberTheoreticTransforms {
	return &cooleyTukeyRecursive{}
}

func cooleyTukeyRecursiveNTT(a []uint64, omega uint64, mod uint64) []uint64 {
	n := misc.NextPowerOf2(len(a))

	// Base case: If the input size is 1, return the input as is.
	if n == 1 {
		return a
	}

	// Pad the input with zeros to the next power of 2.
	a = append(a, make([]uint64, n-(len(a)))...)

	// Split the input into even and odd indices.
	even := make([]uint64, n/2)
	odd := make([]uint64, n/2)
	for i := 0; i < n/2; i++ {
		even[i] = a[2*i]
		odd[i] = a[2*i+1]
	}

	// Recursively compute the NTT of the even and odd halves.
	omegaSq := ModExp(omega, 2, mod) // omega^2 is the primitive root for the smaller NTT.
	evenNTT := cooleyTukeyRecursiveNTT(even, omegaSq, mod)
	oddNTT := cooleyTukeyRecursiveNTT(odd, omegaSq, mod)

	// Combine the results.
	result := make([]uint64, n)
	for k := 0; k < n/2; k++ {
		// Twiddle factor: omega^k
		twiddle := ModExp(omega, uint64(k), mod)
		// Combine even and odd results
		result[k] = ModAdd(evenNTT[k], ModMul(twiddle, oddNTT[k], mod), mod)
		result[k+n/2] = ModAdd(evenNTT[k], ModMul(mod-twiddle, oddNTT[k], mod), mod)
	}

	return result
}

// cooleyTukeyINTT computes the INTT of the input array `a` using the Cooley-Tukey algorithm.
// The input is a []uint64 array, and the output is a []uint16 array.
func cooleyTukeyINTT(a []uint64, omegaInv uint64, mod uint64) []uint16 {
	// Compute the NTT with omegaInv (inverse of omega).
	nttResult := cooleyTukeyRecursiveNTT(a, omegaInv, mod)

	n := uint64(len(nttResult))

	// Scale the result by nInv (modular inverse of n).
	nInv := ModInverseFermat(n, mod) // nInv = n^(-1) mod mod

	result64 := make([]uint64, n)
	for i := range nttResult {
		result64[i] = ModMul(nttResult[i], nInv, mod)
	}

	// Convert the result to []uint16.
	result16 := make([]uint16, n)
	carry := uint64(0)
	for i := range result64 {
		mul := result64[i] + carry
		result16[i] = uint16(mul & 0xFFFF) // Store the lower 16 bits
		carry = mul >> 16                  // Shift to get the carry (if any)
	}

	// If there's any carry left after the loop
	if carry > 0 {
		panic("ERROR: carry > 0")
	}

	return result16
}
