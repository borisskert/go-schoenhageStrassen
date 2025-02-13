package schoenhageStrassen

import (
	"fmt"
	. "go-schoenhageStrassen/arithmetic"
	"go-schoenhageStrassen/modular"
	"go-schoenhageStrassen/ntt"
)

func Multiply16(a, b []uint16) []uint16 {
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	// Ensure the lengths of a and b are the same and a power of two
	//n := nextPowerOfTwo(len(a) + len(b) - 1)
	n := len(a) + len(b)
	aPadded := padSlice(a, n)
	bPadded := padSlice(b, n)

	// Choose a suitable modulus and primitive root of unity
	mod, omega, omegaInv, err := modular.FindModulus16Two(aPadded, bPadded)
	if err != nil {
		panic(err)
	}

	// Compute NTT of a and b
	nttA := ntt.NttNaive(aPadded, uint64(omega), uint64(mod))
	nttB := ntt.NttNaive(bPadded, uint64(omega), uint64(mod))

	// Pointwise multiplication in NTT domain
	productNTT := make([]uint64, n)
	for i := 0; i < n; i++ {
		//productNTT[i] = (nttA[i] * nttB[i]) % mod
		productNTT[i] = ModMul(nttA[i], nttB[i], uint64(mod))
	}

	// Compute INTT to get the result in the original domain
	//omegaInv := ModExp(omega, mod-2, mod) // Inverse of omega
	result := ntt.InttNaive(productNTT, uint64(omegaInv), uint64(mod))

	// Trim leading zeros and return
	result = trimLeadingZeros(result)

	fmt.Println(result)

	return result
}

// padSlice pads the input slice with zeros to the specified length.
func padSlice(a []uint16, length int) []uint16 {
	padded := make([]uint16, length)
	copy(padded, a)
	return padded
}

// nextPowerOfTwo returns the smallest power of two >= n.
func nextPowerOfTwo(n int) int {
	if n == 0 {
		return 1
	}
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n++
	return n
}

// trimLeadingZeros removes leading zeros from the result.
func trimLeadingZeros(a []uint16) []uint16 {
	i := len(a) - 1
	for i >= 0 && a[i] == 0 {
		i--
	}
	if i < 0 {
		return []uint16{0}
	}
	return a[:i+1]
}
