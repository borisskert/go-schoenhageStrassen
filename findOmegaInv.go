package main

import "fmt"

// Compute omega^(-1) mod M
func computeOmegaInverse(omega, mod int64) int64 {
	return modExp(omega, mod-2, mod) // omega^(-1) = omega^(M-2) mod M
}

func findOmegaInvExample() {
	// Example Schönhage-Strassen modulus
	M := int64(4294967297) // 2^32 + 1
	omega, err := findOmega(32, M)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Omega Inverse:", omega)

	omegaInv := computeOmegaInverse(omega, M)
	fmt.Println("Omega Inverse:", omegaInv)

	// Verify correctness
	fmt.Println("Check:", (omega*omegaInv)%M) // Should be 1
}
