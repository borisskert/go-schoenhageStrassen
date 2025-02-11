package main

// Factorize returns the prime factors of a given positive int64 number
func factorize(n int64) []int64 {
	var factors []int64

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
	for i := int64(5); i*i <= n; i += 6 {
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

func factorize2(n int64) (int64, int64) {
	for i := sqrt(n) - 1; i > 1; i-- {
		if n%i == 0 {
			return i, n / i
		}
	}

	return 1, n
}
