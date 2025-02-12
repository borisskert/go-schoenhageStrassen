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
6. The rest of the procedure for the forward and inverse transforms is identical to the complex DFT. Moreover, the ntt can be modified to implement a fast Fourier transform algorithm such as Cooley–Tukey.
*/
func findModulus32(A []int32) (Modulus, Generator, Omega, OmegaInverse, error) {
	n := int64(len(A))

	// 2. Choose a minimum working modulus M such that 1≤n<M and every input value is in the range [0,M).
	m := slices.Max(A)

	return findModulusMN(int64(m), n)
}

func findModulus(A []int64) (Modulus, Generator, Omega, OmegaInverse, error) {
	n := int64(len(A))

	// 2. Choose a minimum working modulus M such that 1≤n<M and every input value is in the range [0,M).
	m := slices.Max(A)

	return findModulusMN(m, n)
}

func findModulusMN(m int64, n int64) (Modulus, Generator, Omega, OmegaInverse, error) {
	fmt.Println("n:", n)

	// When convolving two vectors of length n where each input value is at most m, the upper bound on each output value
	// is (m^2)*n. Choosing a minimum working modulus of M=(m^2)*n+1 is sufficient to always avoid overflow in the worst case.
	M := m*m*n + 1
	fmt.Println("Minimum working modulus M:", M)

	// 3. Select some integer k≥1 and define N=kn+1 as the working modulus. We require N≥M, and N to be a prime number.
	// Dirichlet’s theorem guarantees that for any n and M, there exists some choice of k to make N be prime.
	var N int64
	if isPrime(M) {
		N = M
	} else {
		N = findLargePrime(M)
	}

	k := (N - 1) / n

	for k*n != N-1 {
		N = findLargePrime(N)
		k = (N - 1) / n
	}

	fmt.Println("Actual working modulus:", N)

	// 4. Because N is prime, the multiplicative group of ℤN has size φ(N)=N−1=kn. Furthermore, the group must have at
	// least one generator g, which is also a primitive (N−1)th root of unity.
	var g int64
	var omega int64
	for g = 2; g < N; g++ {

		// 5. Define ω≡g^k mod N. We have ω^n=g^(kn)=g^(N−1)=g^(φ(N))≡1 mod N due to Euler’s theorem. Furthermore because g is a
		// generator, we know that ω^i=g^(ik)≢1 for 1≤i<n, because ik<nk=N−1. Hence ω is a primitive nth root of unity, as
		// required by the DFT of length n.
		omega = modExp(g, k, N)

		if modExp(g, n, N) != 1 && omega != 1 {
			break
		}
	}

	fmt.Println("k:", k)
	fmt.Println("generator:", g)
	fmt.Println("ω:", omega)

	if modExp(omega, k, N) != 1 && modExp(omega, n, N) != 1 {
		return -1, -1, -1, -1, fmt.Errorf("invalid ω")
	}

	omegaInv := modInverse(omega, N)
	fmt.Println("ω:", omega, "ω(Inv):", omegaInv)
	fmt.Println("ω ^ n =", modExp(omega, n, N))           // should be 1 if ω is correct
	fmt.Println("ω * ω(Inv) mod M =", (omega*omegaInv)%N) // should be 1

	return Modulus(N), Generator(g), Omega(omega), OmegaInverse(omegaInv), nil
}

func findModulus2(a, b []int64) (Modulus, Generator, Omega, OmegaInverse, error) {
	n := max(int64(len(a)), int64(len(b))) // TODO padding

	// 2. Choose a minimum working modulus M such that 1≤n<M and every input value is in the range [0,M).
	Ma := slices.Max(a)
	Mb := slices.Max(b)

	m := max(Ma, Mb)

	return findModulusMN(m, n)
}

type Modulus int64

type Generator int64

type Omega int64

type OmegaInverse int64

func sqrt(n int64) int64 {
	var x int64
	for x = 1; x*x < n; x++ {
	}
	return x
}
