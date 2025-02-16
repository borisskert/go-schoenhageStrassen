package naive

import (
	"fmt"
	. "go-schoenhageStrassen/arithmetic"
)

// inttNaive computes the Inverse Number Theoretic Transform.
func inttNaive(a []uint64, omegaInv uint64, mod uint64) []uint16 {
	fmt.Println("INTT input:", a)

	n := len(a)
	result := make([]uint64, n)

	for i := uint64(0); i < uint64(n); i += 1 {
		yi := uint64(0)

		for j := uint64(0); j < uint64(n); j += 1 {
			u := ModExp(omegaInv, i*j, mod)
			v := ModMul(a[j], u, mod)
			yi = ModAdd(yi, v, mod)
		}

		result[i] = yi % mod
	}

	fmt.Println("INTT result:", result)

	nInv := ModExp(uint64(n), mod-2, mod)

	if ModMul(uint64(n), nInv, mod) != 1 {
		fmt.Println("ERROR: n * nInv % mod != 1")
		return nil
	}

	fmt.Println("nInv:", nInv)

	output64 := make([]uint64, n)
	for i := range result {
		output64[i] = ModMul(result[i], nInv, mod)
	}

	fmt.Println("INTT output64:", output64)

	output := make([]uint16, n)
	carry := uint64(0)
	for i := range output64 {
		mul := output64[i] + carry
		output[i] = uint16(mul & 0xFFFF) // Store the lower 16 bits
		carry = mul >> 16                // Shift to get the carry (if any)
	}

	// If there's any carry left after the loop
	if carry > 0 {
		panic("ERROR: carry > 0")
	}

	fmt.Println("INTT output16:", output)

	return output
}
