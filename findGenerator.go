package main

import "fmt"

func findGenerators(factors []int64, mod int64) ([]int64, error) {
	// Find the smallest generator of the group
	// A generator is a number g such that g^i != 1 for all i < mod
	// A generator must be coprime with mod
	// A generator must be a primitive root modulo mod
	// A generator must be a square root of 1 modulo mod
	// A generator must be a square root of -1 modulo mod
	// A generator must be a square root of any prime

	// Find a generators for each prime factor
	var generators []int64
	for _, factor := range factors {
		g, err := findGenerator(factor, mod)
		if err != nil {
			return nil, err
		}
		generators = append(generators, g)
	}

	return generators, nil
}

func findGenerator(factor int64, mod int64) (int64, error) {
	// Find a generator for a prime factor
	for g := int64(2); g < mod; g++ {
		if isGenerator(g, factor, mod) {
			return g, nil
		}
	}
	return -1, fmt.Errorf("No valid generator found")
}

func isGenerator(g int64, factor int64, mod int64) bool {
	// Verify g^2^m != 1 mod M
	return modExp(g, mod/factor, mod) != 1
}
