package main

import "fmt"

func NTT(a []int64, omega int64, mod int64) []int64 {
	n := len(a)
	res := make([]int64, n)
	copy(res, a) // Work on a copy

	fmt.Println("NTT Input:", res)

	step := 1
	for lenBlock := n / 2; lenBlock > 0; lenBlock /= 2 {
		w := int64(1)
		for j := 0; j < n; j += 2 * lenBlock {
			for k := 0; k < lenBlock; k++ {
				u := res[j+k]
				v := (w * res[j+k+lenBlock]) % mod

				res[j+k] = (u + v) % mod
				res[j+k+lenBlock] = (u - v + mod) % mod
			}
			w = (w * omega) % mod
		}
		fmt.Println("Step", step, ":", res)
		step++
	}

	fmt.Println("Final NTT:", res)
	return res
}
