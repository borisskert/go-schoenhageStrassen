package arithmetic

func ModSub(a, b, mod uint64) uint64 {
	a %= mod
	b %= mod

	if a < b {
		diff := safeSub(b, a, mod)
		return safeSub(mod, diff, mod)
	}

	return a - b
}

func safeSub(a, b, mod uint64) uint64 {
	b %= mod

	if a < b {
		b, a = a, b
	}

	return a - b
}
