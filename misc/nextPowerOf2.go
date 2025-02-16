package misc

func NextPowerOf2(n int) int {
	p := 1
	for p < n {
		p *= 2
	}
	return p
}
