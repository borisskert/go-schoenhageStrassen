package main

import (
	"fmt"
)

func findOmega(m int64, mod int64) (int64, error) {
	factors := factorize(mod - 1)

	gs, err := findGenerators(factors, mod)
	if err != nil {
		return -1, err
	}

	// Try each generator to find a suitable omega
	for _, g := range gs {
		omega := modExp(g, factors[0], mod)

		if isValidOmega(omega, m, mod) {
			return omega, nil
		}
	}

	return -1, fmt.Errorf("No valid omega found")
}

//// Find a 2^m-th root of unity efficiently
//func findOmega(m int64, mod int64) (int64, error) {
//	exp := (mod - 1) >> m // (M-1) / 2^m
//
//	for g := int64(2); g < mod; g++ {
//		omega := modExp(g, exp, mod) // Candidate ω
//
//		// Check if ω is a primitive root of unity
//		if omega != 1 &&
//			modExp(omega, 1<<m, mod) == 1 && // ω^(2^m) ≡ 1 (mod M)
//			modExp(omega, 1<<(m-1), mod) != 1 { // ω^(2^(m-1)) ≠ 1 (mod M)
//			return omega, nil
//		}
//	}
//
//	return -1, fmt.Errorf("No valid omega found")
//}

//// Find a 2^m-th root of unity (fast method)
//func findOmega(m int64, mod int64) (int64, error) {
//	exp := (mod - 1) >> m           // (M-1) / 2^m
//	return modExp(2, exp, mod), nil // Use 2 as base (primitive root for Fermat primes)
//}

//// Find a 2^m-th root of unity for Schönhage-Strassen (optimized)
//func findOmega(m int64, mod int64) (int64, error) {
//	exp := (mod - 1) >> m // (M-1) / 2^m
//
//	for g := int64(2); g < mod; g++ {
//		omega := modExp(g, exp, mod) // Candidate ω
//
//		// Check both conditions for a 2^m-th root of unity
//		if omega != 1 &&
//			modExp(omega, 1<<m, mod) == 1 && // ω^(2^m) ≡ 1 (mod M)
//			modExp(omega, 1<<(m-1), mod) != 1 { // ω^(2^(m-1)) ≠ 1 (mod M)
//			return omega, nil
//		}
//	}
//
//	return -1, fmt.Errorf("No valid omega found")
//}

//// Find a suitable omega (2^m-th root of unity) for Schönhage-Strassen
//func findOmega(m int64, mod int64) (int64, error) {
//	// Compute (M-1)/2^m
//	exp := (mod - 1) >> m // Equivalent to (M-1)/2^m
//	//exp := (mod - 1) / modExp(2, m, mod)
//
//	// Try small base numbers like 3, 5, 7, ...
//	for g := int64(2); g < mod; g++ {
//		omega := modExp(g, exp, mod)
//
//		if isValidOmega(omega, m, mod) {
//			return omega, nil
//		}
//	}
//
//	return -1, fmt.Errorf("No valid omega found")
//}

func findOmega2(m int64, mod int64) (int64, error) {
	exp := (mod - 1) >> m // Faster computation of (M-1)/2^m

	// Test small prime bases (faster than checking every g)
	smallPrimes := []int64{3, 5, 7, 11, 13, 17, 19, 23, 29, 31}

	for _, g := range smallPrimes {
		omega := modExp(g, exp, mod) // Compute candidate root

		// Verify omega is a valid 2^m-th root of unity
		if omega != 1 && modExp(omega, 1<<m, mod) == 1 {
			return omega, nil
		}
	}

	return -1, fmt.Errorf("No valid omega found")
}

func isValidOmega(omega int64, m int64, mod int64) bool {
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

	omega, err := findOmega(int64(m), M)
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

	omega, err := findOmega(int64(m), M)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Omega for Schönhage-Strassen FFT:", omega)
	}
}
