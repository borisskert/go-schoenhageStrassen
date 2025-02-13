package modular

// Factorize returns the prime factors of a given positive int64 number
func factorize(n uint64) []uint64 {
	var factors []uint64

	// Handle factor 2 separately
	for n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}

	for n%3 == 0 {
		factors = append(factors, 3)
		n /= 3
	}

	// Check odd factors from 3 to sqrt(n)
	for i := uint64(5); i*i <= n; i += 6 {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
		for n%(i+2) == 0 {
			factors = append(factors, i+2)
			n /= i + 2
		}
	}

	// If n is still greater than 1, it must be a prime number
	if n > 1 {
		factors = append(factors, n)
	}

	return factors
}
