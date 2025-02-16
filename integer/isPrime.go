package integer

func IsPrime(n uint64) bool {
	return millerRabin(n, 100)
}

func isPrimeForSure(n uint64) bool { //nolint:unused
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
