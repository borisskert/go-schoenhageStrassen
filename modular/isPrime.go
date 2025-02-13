package modular

func isPrime(n uint64) bool {
	if n < 2 {
		return false
	}
	if n < 4 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := uint64(5); i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

func findNextPrime(M uint64) uint64 {
	for n := M + 1; ; n++ {
		if isPrime(n) {
			return n
		}
	}
}
