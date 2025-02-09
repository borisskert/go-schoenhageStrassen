package main

import "fmt"

// Modular exponentiation: (base^exp) % mod
func modExp2(base, exp, mod int64) int64 {
	result := int64(1)
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exp /= 2
	}
	return result
}

func checkOmega() {
	mod := int64(4294967297) // M = 2^32 + 1
	n := int64(4)            // Example FFT size
	//g := int64(3)            // Primitive root candidate

	var g int64 // Primitive root candidate

	for g = int64(1); g < 4294967297; g++ {
		omega := modExp2(g, (mod-1)/n, mod) // Compute nth root of unity

		// Verify ω^n mod M == 1
		check := modExp2(omega, n, mod)

		if check == 1 {
			break
		}
	}

	omega := modExp2(g, (mod-1)/n, mod) // Compute nth root of unity
	fmt.Println("Computed ω:", omega)

	// Verify ω^n mod M == 1
	check := modExp2(omega, n, mod)
	fmt.Println("Check ω^n mod M:", check) // Should print 1
}
