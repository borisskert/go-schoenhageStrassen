package schoenhageStrassen

import (
	"fmt"
	. "go-schoenhageStrassen/arithmetic"
	"go-schoenhageStrassen/array"
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
	result = array.TrimLeadingZeros16(result)

	fmt.Println(result)

	return result
}

// padSlice pads the input slice with zeros to the specified length.
func padSlice(a []uint16, length int) []uint16 {
	padded := make([]uint16, length)
	copy(padded, a)
	return padded
}

func Multiply32(a, b []uint32) []uint32 {
	a32 := array.Split32To16(a)
	b32 := array.Split32To16(b)

	result16 := Multiply16(a32, b32)
	result32 := array.Combine16To32(result16)

	return array.TrimLeadingZeros32(result32)
}

func Multiply64(a, b []uint64) []uint64 {
	a16 := array.Split64To16(a)
	b16 := array.Split64To16(b)

	result16 := Multiply16(a16, b16)
	result64 := array.Combine16To64(result16)

	return array.TrimLeadingZeros64(result64)
}
