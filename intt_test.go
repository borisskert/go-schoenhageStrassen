package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("INTT", func() {
	It("INTT(NTT([1, 2, 3, 4]))", func() {
		A := []int64{1, 2, 3, 4}

		mod, _, omega, omegaInv, err := findModulus(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, 4, int64(omega), int64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, 4, int64(omegaInv), int64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 7, 3, 5]))", func() {
		A := []int64{1, 7, 3, 5}

		mod, _, omega, omegaInv, err := findModulus(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, 4, int64(omega), int64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, 4, int64(omegaInv), int64(mod))
		Expect(A_recovered).To(Equal(A))
	})
})
