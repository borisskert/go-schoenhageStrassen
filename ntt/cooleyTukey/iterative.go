package cooleyTukey

import (
	. "go-schoenhageStrassen/arithmetic"
	"go-schoenhageStrassen/array"
	"go-schoenhageStrassen/misc"
	"go-schoenhageStrassen/ntt"
)

type cooleyTukeyIterative struct{}

func (c cooleyTukeyIterative) NTT(a []uint16, omega uint64, mod uint64) []uint64 {
	n := misc.NextPowerOf2a(len(a))

	a64 := make([]uint64, n)
	for i := range a {
		a64[i] = uint64(a[i])
	}

	return cooleyTukeyIterativeNTT(a64, omega, mod)
}

func (c cooleyTukeyIterative) INTT(a []uint64, omega uint64, mod uint64) []uint16 {
	return inttCooleyTukeyIterative(a, omega, mod)
}

func NewCooleyTukeyIterative() ntt.NumberTheoreticTransforms {
	return &cooleyTukeyIterative{}
}

func inttCooleyTukeyIterative(a []uint64, omegaInv uint64, mod uint64) []uint16 {
	// Compute the NTT with omegaInv (inverse of omega).
	nttResult := cooleyTukeyIterativeNTT(a, omegaInv, mod)

	n := uint64(len(nttResult))

	// Scale the result by nInv (modular inverse of n).
	nInv := ModExp(n, mod-2, mod) // nInv = n^(-1) mod mod
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

func cooleyTukeyIterativeNTT(a []uint64, omega uint64, mod uint64) []uint64 {
	n := misc.NextPowerOf2a(len(a))

	// Bit-reversal permutation
	result := array.BitReverseCopyN(a, uint64(n))

	// Cooley-Tukey NTT (iterative)
	for s := 1; s < n; s <<= 1 {
		m := s << 1
		wm := ModExp(omega, uint64(n/m), mod) // Twiddle factor for this stage

		for k := 0; k < n; k += m {
			w := uint64(1)
			for j := 0; j < s; j++ {
				// Butterfly operation
				t := ModMul(w, result[k+j+s], mod)
				u := result[k+j]
				result[k+j] = ModAdd(u, t, mod)
				result[k+j+s] = ModSub(u, t, mod)
				w = ModMul(w, wm, mod) // Update twiddle factor
			}
		}
	}

	return result
}
