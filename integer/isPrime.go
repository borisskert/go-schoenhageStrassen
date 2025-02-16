package integer

func IsPrime(n uint64) bool {
	return millerRabin(n, 10)
}

func isPrimeForSure(n uint64) bool {
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

func FindNextPrime(M uint64) uint64 {
	if M < 2 {
		return 2
	}

	if M < 3 {
		return 3
	}

	if M%2 == 0 {
		M++
	} else {
		M += 2
	}

	for n := M; ; n += 2 {
		if IsPrime(n) {
			return n
		}
	}
}
