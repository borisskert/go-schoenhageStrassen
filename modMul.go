package main

func modMul(a, b, mod int64) int64 {
	return ((a % mod) * (b % mod)) % mod
}
