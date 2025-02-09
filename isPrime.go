package main

import "fmt"

func isPrime(n int64) bool {
	if n < 2 {
		return false
	}
	if n < 4 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := int64(5); i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

func findLargePrime() int64 {
	for n := int64(4294967297); ; n++ {
		if isPrime(n) {
			return n
		}
	}
}

func checkPrimeExample() {
	n := int64(4294967297) // M = 2^32 + 1
	fmt.Println("Is Prime:", isPrime(n))
}

func findLargePrimeExample() {
	n := findLargePrime()
	fmt.Println("Largest Prime:", n)
}
