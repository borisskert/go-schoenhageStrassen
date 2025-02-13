package main

func nextPowerOf2(n int64) int64 {
	p := int64(1)
	for p < n {
		p *= 2
	}
	return p
}
