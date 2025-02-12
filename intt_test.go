package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("INTT", func() {
	It("INTT(NTT([1, 2, 3, 4]))", func() {
		A := []int32{1, 2, 3, 4}

		mod, _, omega, omegaInv, err := findModulus32(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, int64(omega), int64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, int64(omegaInv), int64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 2, 3, 4, 5, 6, 7, 8]))", func() {
		A := []int32{1, 2, 3, 4, 5, 6, 7, 8}

		mod, _, omega, omegaInv, err := findModulus32(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, int64(omega), int64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, int64(omegaInv), int64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16]))", func() {
		A := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

		mod, _, omega, omegaInv, err := findModulus32(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, int64(omega), int64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, int64(omegaInv), int64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 7, 3, 5]))", func() {
		A := []int32{1, 7, 3, 5}

		mod, _, omega, omegaInv, err := findModulus32(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, int64(omega), int64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, int64(omegaInv), int64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([3, 2, 1]))", func() {
		A := []int32{3, 2, 1}

		mod, _, omega, omegaInv, err := findModulus32(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, int64(omega), int64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, int64(omegaInv), int64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([10000001, 10000002, 10000003], 10000004))", func() {
		A := []int32{10000001, 10000002, 10000003, 10000004}

		mod, _, omega, omegaInv, err := findModulus32(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, int64(omega), int64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, int64(omegaInv), int64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 2, 3, 4, 5, 6]))", func() {
		A := []int32{1, 2, 3, 4, 5, 6}

		mod, _, omega, omegaInv, err := findModulus32(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, int64(omega), int64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, int64(omegaInv), int64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 10, 100, 1000]))", func() {
		A := []int32{1, 10, 100, 1000}

		mod, _, omega, omegaInv, err := findModulus32(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, int64(omega), int64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, int64(omegaInv), int64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 10, 100, 1000, 10000]))", func() {
		A := []int32{1, 10, 100, 1000, 10000}

		mod, _, omega, omegaInv, err := findModulus32(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, int64(omega), int64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, int64(omegaInv), int64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 10, 100, 1000, 10000, 100000]))", func() {
		A := []int32{1, 10, 100, 1000, 10000, 100000}

		mod, _, omega, omegaInv, err := findModulus32(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, int64(omega), int64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, int64(omegaInv), int64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 10, 100, 1000, 10000, 100000, 1000000]))", func() {
		A := []int32{1, 10, 100, 1000, 10000, 100000, 1000000}

		mod, _, omega, omegaInv, err := findModulus32(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, int64(omega), int64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, int64(omegaInv), int64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 10, 100, 1000, 10000, 100000, 1000000, 10000000]))", func() {
		A := []int32{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000}

		mod, _, omega, omegaInv, err := findModulus32(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, int64(omega), int64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, int64(omegaInv), int64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000]))", func() {
		A := []int32{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000}

		mod, _, omega, omegaInv, err := findModulus32(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, int64(omega), int64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, int64(omegaInv), int64(mod))
		Expect(A_recovered).To(Equal(A))
	})
})
