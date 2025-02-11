package main

func padToPowerOfTwo(A []int64) []int64 {
	n := len(A)
	nextPower := 1
	for nextPower < n {
		nextPower *= 2
	}
	padded := make([]int64, nextPower)
	copy(padded, A) // Copy original values
	return padded
}
