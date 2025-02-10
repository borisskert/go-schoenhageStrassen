package main

import (
	"fmt"
)

// Find a suitable omega (2^m-th root of unity) for Schönhage-Strassen
func findOmega(m int, mod int64) (int64, error) {
	// Compute (M-1)/2^m
	//exp := (mod - 1) >> m // Equivalent to (M-1)/2^m
	exp := (mod - 1) / modExp(2, int64(m), mod)

	// Try small base numbers like 3, 5, 7, ...
	for g := int64(2); g < mod; g++ {
		omega := modExp(g, exp, mod)

		if isValidOmega(omega, m, mod) {
			return omega, nil
		}
	}

	return -1, fmt.Errorf("No valid omega found")
}

func isValidOmega(omega int64, m int, mod int64) bool {
	if omega != 1 {
		// Verify omega^2^m == 1 mod M
		if modExp(omega, 1<<m, mod) == 1 {
			return true
		}
	}

	return false
}

func findOmegaExample4294967297() {
	// Schönhage-Strassen modulus
	M := int64(4294967297) // 2^32 + 1
	m := 32                // 2^m

	omega, err := findOmega(m, M)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Omega for Schönhage-Strassen FFT:", omega)
	}
}

func findOmegaExample7681() {
	// Schönhage-Strassen modulus
	M := int64(7681) // 2^32 + 1
	m := 13          // 2^m

	omega, err := findOmega(m, M)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Omega for Schönhage-Strassen FFT:", omega)
	}
}
