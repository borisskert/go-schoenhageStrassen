package arithmetic

// Fermat’s Little Theorem (When mod is Prime)
func modInverseFermat(n, mod uint64) uint64 {
	return ModExp(n, mod-2, mod) // Uses modExp function from before
}

// Compute modular inverse using Extended Euclidean Algorithm
func ModInverse(n, mod uint64) uint64 {
	g, x, _ := extendedGCD(n, mod)
	if g != 1 {
		panic("Inverse doesn't exist") // Only exists if gcd(n, mod) = 1
	}
	// Ensure x is positive
	return safeMod(x, mod)
}

func safeMod(x int64, mod uint64) uint64 {
	if x < 0 {
		_, _, r := divmod64(0, uint64(-x), mod)
		return mod - r
	}

	return uint64(x % int64(mod))
}

func extendedGCD(a, b uint64) (uint64, int64, int64) {
	if a == 0 {
		return b, 0, 1
	}
	g, x1, y1 := extendedGCD(b%a, a)
	x := y1 - int64(b/a)*x1
	y := x1
	return g, x, y
}
