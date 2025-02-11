package main

func log2(n int64) int64 {
	if n == 0 {
		return 0
	}

	var result int64
	for n > 1 {
		n /= 2
		result++
	}

	return result
}
