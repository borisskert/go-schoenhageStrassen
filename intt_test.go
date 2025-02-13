package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("INTT", func() {
	It("INTT(NTT([1, 1, 1, 1]))", func() {
		A := []uint64{1, 1, 1, 1}

		mod, omega, omegaInv, err := findModulus64(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([10, 10, 10, 10]))", func() {
		A := []uint64{10, 10, 10, 10}

		mod, omega, omegaInv, err := findModulus64(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([100, 100, 100, 100]))", func() {
		A := []uint64{100, 100, 100, 100}

		mod, omega, omegaInv, err := findModulus64(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1000, 1000, 1000, 1000]))", func() {
		A := []uint64{1000, 1000, 1000, 1000}

		mod, omega, omegaInv, err := findModulus64(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([10000, 10000, 10000, 10000]))", func() {
		A := []uint64{10000, 10000, 10000, 10000}

		mod, omega, omegaInv, err := findModulus64(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([100000, 100000, 100000, 100000]))", func() {
		A := []uint64{100000, 100000, 100000, 100000}

		mod, omega, omegaInv, err := findModulus64(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1000000, 1000000, 1000000, 1000000]))", func() {
		A := []uint64{1000000, 1000000, 1000000, 1000000}

		mod, omega, omegaInv, err := findModulus64(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([10000000, 10000000, 10000000, 10000000]))", func() {
		A := []uint64{10000000, 10000000, 10000000, 10000000}

		mod, omega, omegaInv, err := findModulus64(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([100000000, 100000000, 100000000, 100000000]))", func() {
		A := []uint64{100000000, 100000000, 100000000, 100000000}

		mod, omega, omegaInv, err := findModulus64(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 2, 3, 4]))", func() {
		A := []uint64{1, 2, 3, 4}

		mod, omega, omegaInv, err := findModulus64(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 2, 3, 4, 5, 6, 7, 8]))", func() {
		A := []uint64{1, 2, 3, 4, 5, 6, 7, 8}

		mod, omega, omegaInv, err := findModulus64(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16]))", func() {
		A := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

		mod, omega, omegaInv, err := findModulus64(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 7, 3, 5]))", func() {
		A := []uint64{1, 7, 3, 5}

		mod, omega, omegaInv, err := findModulus64(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([3, 2, 1]))", func() {
		A := []uint64{3, 2, 1}

		mod, omega, omegaInv, err := findModulus64(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([10001, 10002, 10003], 10004))", func() {
		A := []uint64{10001, 10002, 10003, 10004}

		mod, omega, omegaInv, err := findModulus64(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([100001, 100002, 100003], 100004))", func() {
		A := []uint64{100001, 100002, 100003, 100004}

		mod, omega, omegaInv, err := findModulus64(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1000001, 1000002, 1000003], 1000004))", func() {
		A := []uint64{1000001, 1000002, 1000003, 1000004}

		mod, omega, omegaInv, err := findModulus64(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([10000001, 10000002, 10000003], 10000004))", func() {
		A := []uint64{10000001, 10000002, 10000003, 10000004}

		mod, omega, omegaInv, err := findModulus64(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 2, 3, 4, 5, 6]))", func() {
		A := []uint64{1, 2, 3, 4, 5, 6}

		mod, omega, omegaInv, err := findModulus64(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 10, 100, 1000]))", func() {
		A := []uint64{1, 10, 100, 1000}

		mod, omega, omegaInv, err := findModulus64(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 10, 100, 1000, 10000]))", func() {
		A := []uint64{1, 10, 100, 1000, 10000}

		mod, omega, omegaInv, err := findModulus64(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 10, 100, 1000, 10000, 100000]))", func() {
		A := []uint64{1, 10, 100, 1000, 10000, 100000}

		mod, omega, omegaInv, err := findModulus64(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 10, 100, 1000, 10000, 100000, 1000000]))", func() {
		A := []uint64{1, 10, 100, 1000, 10000, 100000, 1000000}

		mod, omega, omegaInv, err := findModulus64(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 10, 100, 1000, 10000, 100000, 1000000, 10000000]))", func() {
		A := []uint64{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000}

		mod, omega, omegaInv, err := findModulus64(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000]))", func() {
		A := []uint64{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000}

		mod, omega, omegaInv, err := findModulus64(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([2147483647, 2147483647, 2147483647, 2147483647]))", func() {
		A := []uint64{2147483647, 2147483647, 2147483647, 2147483647}

		mod, omega, omegaInv, err := findModulus64(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})
})

var _ = Describe("INTT 16bit", func() {

	It("INTT(NTT([0, 1])) (16bit)", func() {
		A := []uint16{0, 1}

		mod, omega, omegaInv, err := findModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT16(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT16(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([0, 1, 0, 0])) (16bit)", func() {
		A := []uint16{0, 1, 0, 0}

		mod, omega, omegaInv, err := findModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT16(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := trimZeros16(INTT16(A_ntt, uint64(omegaInv), uint64(mod)))
		Expect(A_recovered).To(Equal(trimZeros16(A)))
	})

	It("INTT(NTT([1, 1, 1, 1])) (16bit)", func() {
		A := []uint16{1, 1, 1, 1}

		mod, omega, omegaInv, err := findModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT16(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT16(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 1, 1])) (16bit)", func() {
		A := []uint16{1, 1, 1}

		mod, omega, omegaInv, err := findModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT16(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT16(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([10, 10, 10, 10]))", func() {
		A := []uint16{10, 10, 10, 10}

		mod, omega, omegaInv, err := findModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT16(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT16(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([100, 100, 100, 100]))", func() {
		A := []uint16{100, 100, 100, 100}

		mod, omega, omegaInv, err := findModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT16(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT16(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1000, 1000, 1000, 1000]))", func() {
		A := []uint16{1000, 1000, 1000, 1000}

		mod, omega, omegaInv, err := findModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT16(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT16(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([10000, 10000, 10000, 10000]))", func() {
		A := []uint16{10000, 10000, 10000, 10000}

		mod, omega, omegaInv, err := findModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT16(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT16(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 10, 100, 1000, 10000]))", func() {
		A := []uint16{1, 10, 100, 1000, 10000}

		mod, omega, omegaInv, err := findModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT16(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT16(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([65535, 65535, 65535, 65535]))", func() {
		A := []uint16{65535, 65535, 65535, 65535}

		mod, omega, omegaInv, err := findModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT16(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT16(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([0, 1, 3, 15, 31, 63, 127, 255, 511, 1023, 2047, 4095, 8191, 16383, 32767, 65535]))", func() {
		A := []uint16{0, 1, 3, 15, 31, 63, 127, 255, 511, 1023, 2047, 4095, 8191, 16383, 32767, 65535}

		mod, omega, omegaInv, err := findModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT16(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT16(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

})
