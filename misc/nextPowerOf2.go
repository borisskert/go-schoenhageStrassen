package misc

func NextPowerOf2(n uint64) uint64 {
	p := uint64(1)
	for p < n {
		p *= 2
	}
	return p
}

func NextPowerOf2a(n int) int {
	p := 1
	for p < n {
		p *= 2
	}
	return p
}
