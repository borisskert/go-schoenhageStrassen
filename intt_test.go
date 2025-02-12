package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("INTT", func() {
	It("INTT(NTT([1, 1, 1, 1]))", func() {
		A := []int32{1, 1, 1, 1}

		mod, _, omega, omegaInv, err := findModulus32(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, int64(omega), int64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, int64(omegaInv), int64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([10, 10, 10, 10]))", func() {
		A := []int32{10, 10, 10, 10}

		mod, _, omega, omegaInv, err := findModulus32(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, int64(omega), int64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, int64(omegaInv), int64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([100, 100, 100, 100]))", func() {
		A := []int32{100, 100, 100, 100}

		mod, _, omega, omegaInv, err := findModulus32(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, int64(omega), int64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, int64(omegaInv), int64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1000, 1000, 1000, 1000]))", func() {
		A := []int32{1000, 1000, 1000, 1000}

		mod, _, omega, omegaInv, err := findModulus32(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, int64(omega), int64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, int64(omegaInv), int64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([10000, 10000, 10000, 10000]))", func() {
		A := []int32{10000, 10000, 10000, 10000}

		mod, _, omega, omegaInv, err := findModulus32(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, int64(omega), int64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, int64(omegaInv), int64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([100000, 100000, 100000, 100000]))", func() {
		A := []int32{100000, 100000, 100000, 100000}

		mod, _, omega, omegaInv, err := findModulus32(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, int64(omega), int64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, int64(omegaInv), int64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1000000, 1000000, 1000000, 1000000]))", func() {
		A := []int32{1000000, 1000000, 1000000, 1000000}

		mod, _, omega, omegaInv, err := findModulus32(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, int64(omega), int64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, int64(omegaInv), int64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([10000000, 10000000, 10000000, 10000000]))", func() {
		A := []int32{10000000, 10000000, 10000000, 10000000}

		mod, _, omega, omegaInv, err := findModulus32(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, int64(omega), int64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, int64(omegaInv), int64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([100000000, 100000000, 100000000, 100000000]))", func() {
		A := []int32{100000000, 100000000, 100000000, 100000000}

		mod, _, omega, omegaInv, err := findModulus32(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, int64(omega), int64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, int64(omegaInv), int64(mod))
		Expect(A_recovered).To(Equal(A))
	})

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

	It("INTT(NTT([10001, 10002, 10003], 10004))", func() {
		A := []int32{10001, 10002, 10003, 10004}

		mod, _, omega, omegaInv, err := findModulus32(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, int64(omega), int64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, int64(omegaInv), int64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([100001, 100002, 100003], 100004))", func() {
		A := []int32{100001, 100002, 100003, 100004}

		mod, _, omega, omegaInv, err := findModulus32(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, int64(omega), int64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, int64(omegaInv), int64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1000001, 1000002, 1000003], 1000004))", func() {
		A := []int32{1000001, 1000002, 1000003, 1000004}

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
