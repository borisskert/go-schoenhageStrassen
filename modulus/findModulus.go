package modulus

import (
	"errors"
	. "go-schoenhageStrassen/arithmetic"
	"go-schoenhageStrassen/integer"
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

func FindModulus16(A []uint16) (uint64, uint64, uint64, error) {
	n := int64(len(A))

	// 2. Choose a minimum working modulus M such that 1≤n<M and every input value is in the range [0,M).
	m := slices.Max(A)

	return findModulusMN(uint64(m), uint64(n))
}

func FindModulusForTwo16(A, B []uint16) (uint64, uint64, uint64, error) {
	n := max(len(A), len(B))

	// 2. Choose a minimum working modulus M such that 1≤n<M and every input value is in the range [0,M).
	Ma := uint64(slices.Max(A))
	Mb := uint64(slices.Max(B))

	m := Ma * Mb

	return findModulusMN(m, uint64(n))
}

func findModulusMN(m uint64, n uint64) (uint64, uint64, uint64, error) {
	// 1. Compute minimum modulus M to prevent overflow
	N, k := findWorkingModulus(m, n)

	omega, omegaInv, err := findOmega(N, n, k)
	for err != nil {
		// If ω is invalid, find the next working modulus
		N, k = findNextWorkingModulus(N+1, n)
		omega, omegaInv, err = findOmega(N, n, k)
	}

	return N, *omega, *omegaInv, nil
}

func findOmega(N uint64, n, k uint64) (*uint64, *uint64, error) {
	// 3. Find a primitive root `g`
	g, err := findPrimitiveRoot(N)
	if err != nil {
		return nil, nil, errors.New("primitive root not found")
	}

	// 4. Compute `ω = g^k mod N`
	omega := ModExp(*g, k, N)

	// 5. Validate `ω` is a primitive nth root of unity
	if omega == 0 || omega == 1 {
		return nil, nil, errors.New("invalid ω")
	}

	if ModExp(omega, n, N) != 1 {
		return nil, nil, errors.New("invalid ω")
	}

	if ModExp(omega, k, N) == 1 {
		return nil, nil, errors.New("invalid ω")
	}

	// 6. Compute modular inverse of ω
	omegaInv := ModInverse(omega, N)

	return &omega, &omegaInv, nil
}

func findWorkingModulus(m uint64, n uint64) (uint64, uint64) {
	M := m*n + 1
	return findNextWorkingModulus(M, n)
}

func findNextWorkingModulus(M uint64, n uint64) (uint64, uint64) {
	// 2. Find the smallest prime `N` such that N = kn + 1
	var N uint64
	if integer.IsPrime(M) {
		N = M
	} else {
		N = integer.FindNextPrime(M)
	}

	k := (N - 1) / n

	for k*n != N-1 {
		N = integer.FindNextPrime(N)
		k = (N - 1) / n
	}

	return N, k
}
