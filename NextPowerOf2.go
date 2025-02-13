package main

func nextPowerOf2(n uint64) uint64 {
	p := uint64(1)
	for p < n {
		p *= 2
	}
	return p
}

func nextPowerOf2a(n int) int {
	p := int(1)
	for p < n {
		p *= 2
	}
	return p
}
