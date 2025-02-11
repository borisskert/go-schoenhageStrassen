package main

import (
	"fmt"
	"slices"
)

// https://www.nayuki.io/page/number-theoretic-transform-integer-dft
/*
1. Suppose the input vector is a sequence of n non-negative integers.
2. Choose a minimum working modulus M such that 1≤n<M and every input value is in the range [0,M).
3. Select some integer k≥1 and define N=kn+1 as the working modulus. We require N≥M, and N to be a prime number. Dirichlet’s theorem guarantees that for any n and M, there exists some choice of k to make N be prime.
4. Because N is prime, the multiplicative group of ℤN has size φ(N)=N−1=kn. Furthermore, the group must have at least one generator g, which is also a primitive (N−1)th root of unity.
5. Define ω≡gk mod N. We have ωn=gkn=gN−1=gφ(N)≡1 mod N due to Euler’s theorem. Furthermore because g is a generator, we know that ωi=gik≢1 for 1≤i<n, because ik<nk=N−1. Hence ω is a primitive nth root of unity, as required by the DFT of length n.
6. The rest of the procedure for the forward and inverse transforms is identical to the complex DFT. Moreover, the NTT can be modified to implement a fast Fourier transform algorithm such as Cooley–Tukey.
*/
func findModulus32(a []int32) modulus {
	// 2. Choose a minimum working modulus M such that 1≤n<M and every input value is in the range [0,M).
	M := int64(slices.Max(a)) + 1

	// 3. Select some integer k≥1 and define N=kn+1 as the working modulus. We require N≥M, and N to be a prime number.
	// Dirichlet’s theorem guarantees that for any n and M, there exists some choice of k to make N be prime.
	var N int64
	if isPrime(M) {
		N = M
	} else {
		N = findLargePrime(M)
	}

	var n, k int64

	for k = 2; N != k*n+1; k++ {
		n = (N - 1) / k
	}

	// 4. Because N is prime, the multiplicative group of ℤN has size φ(N)=N−1=kn. Furthermore, the group must have at
	// least one generator g, which is also a primitive (N−1)th root of unity.
	var g int64
	for g = 2; g < N; g++ {
		if (k == 1 || modExp(g, (N-1)/k, N) != 1) && (n == 1 || modExp(g, (N-1)/n, N) != 1) {
			break
		}
	}

	omega := modExp(g, k, N) // Compute nth root of unity

	fmt.Println("M:", M)
	fmt.Println("N:", N)
	fmt.Println("n:", n)
	fmt.Println("k:", k)
	fmt.Println("g:", g)
	fmt.Println("ω:", omega)
	if modExp(omega, n, N) != 1 {
		fmt.Println("validate ω:", modExp(omega, n, N)) // Should print 1
		panic("invalid ω")
	}

	panic("not implemented")
}

type modulus int64

func sqrt(n int64) int64 {
	var x int64
	for x = 1; x*x < n; x++ {
	}
	return x
}
