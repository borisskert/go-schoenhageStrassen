package main

import (
	"fmt"
)

func checkNttInttExample() {
	A := []int64{1, 2, 3, 4, 5, 6, 7, 9, 10} // Input array
	n := int64(len(A))
	n = NextPowerOf2(n) // Ensure n is a power of 2

	A = append(A, make([]int64, n-int64(len(A)))...) // Pad with zeros

	fmt.Println("Original A:", A)

	mod, _, omega, omegaInv, err := findModulus(A)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Compute NTT
	A_ntt := NTT(A, int64(omega), int64(mod))
	fmt.Println("After NTT :", A_ntt)

	// Compute INTT
	A_recovered := INTT(A_ntt, int64(omegaInv), int64(mod))

	// Verify recovery
	fmt.Println("After INTT:", A_recovered) // Should match original A
}
