package main

// Fermat’s Little Theorem (When mod is Prime)
func modInverseFermat(n, mod int64) int64 {
	return modExp(n, mod-2, mod) // Uses modExp function from before
}

// Compute modular inverse using Extended Euclidean Algorithm
func modInverse(n, mod int64) int64 {
	g, x, _ := extendedGCD(n, mod)
	if g != 1 {
		panic("Inverse doesn't exist") // Only exists if gcd(n, mod) = 1
	}
	return (x%mod + mod) % mod // Ensure non-negative result
}

func extendedGCD(a, b int64) (int64, int64, int64) {
	if a == 0 {
		return b, 0, 1
	}
	g, x1, y1 := extendedGCD(b%a, a)
	x := y1 - (b/a)*x1
	y := x1
	return g, x, y
}
