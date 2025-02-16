package integer

import (
	"github.com/borisskert/go-schoenhageStrassen/arithmetic"
	"math/rand/v2"
)

// MillerRabin checks if a number is likely prime using the Miller-Rabin test.
// n: The number to test for primality.
// k: The number of iterations (higher k increases accuracy).
func millerRabin(n uint64, k int) bool {
	// Handle small numbers and even numbers
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	// Decompose n-1 into d * 2^s
	s := 0
	d := n - 1
	for d%2 == 0 {
		d /= 2
		s++
	}

	// Witness loop
	for i := 0; i < k; i++ {
		// Pick a random integer a in [2, n-2]
		a := rand.Uint64()%(n-3) + 2

		// Compute x = a^d mod n
		x := arithmetic.ModExp(a, d, n)

		// If x == 1 or x == n-1, continue to the next iteration
		if x == 1 || x == n-1 {
			continue
		}

		// Repeat squaring x up to s-1 times
		for j := 0; j < s-1; j++ {
			x = arithmetic.ModExp(x, 2, n)
			if x == n-1 {
				break
			}
		}

		// If x != n-1, n is composite
		if x != n-1 {
			return false
		}
	}

	// If no witness found, n is likely prime
	return true
}
