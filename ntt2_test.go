package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go-schoenhageStrassen/modular"
)

var _ = Describe("INTT 16bit", func() {

	It("[6, 0, 10, 7, 2] -> [3, 7, 0, 5, 4]", func() {
		A := []uint16{6, 0, 10, 7, 2}

		omega := uint64(3)
		mod := uint64(11)

		result := NTT2(A, omega, mod)

		Expect(result).To(Equal([]uint64{3, 7, 0, 5, 4}))
	})

	It("NTT()", func() {
		A := []uint16{0, 0, 0, 1, 0, 0, 0, 0}
		expected := []uint64{1, 15, 4, 9, 16, 2, 13, 8}

		mod, omega, _, err := findModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT2(A, uint64(omega), uint64(mod))

		Expect(A_ntt).To(Equal(expected))
	})

	It("INTT(NTT([1, 1, 1, 1])) (16bit)", func() {
		A := []uint16{1, 1, 1, 1}

		mod, omega, omegaInv, err := findModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT2(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT2a(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 1, 1])) (16bit)", func() {
		A := []uint16{1, 1, 1}

		mod, omega, omegaInv, err := findModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT2(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := trimZeros16(INTT2a(A_ntt, uint64(omegaInv), uint64(mod)))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([10, 10, 10, 10]))", func() {
		A := []uint16{10, 10, 10, 10}

		mod, omega, omegaInv, err := findModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT2(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT2a(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([100, 100, 100, 100]))", func() {
		A := []uint16{100, 100, 100, 100}

		mod, omega, omegaInv, err := findModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT2(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT2a(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1000, 1000, 1000, 1000]))", func() {
		A := []uint16{1000, 1000, 1000, 1000}

		mod, omega, omegaInv, err := findModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT2(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT2a(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([10000, 10000, 10000, 10000]))", func() {
		A := []uint16{10000, 10000, 10000, 10000}

		mod, omega, omegaInv, err := findModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT2(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT2a(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 10, 100, 1000, 10000]))", func() {
		A := []uint16{1, 10, 100, 1000, 10000}

		mod, omega, omegaInv, err := findModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT2(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := trimZeros16(INTT2a(A_ntt, uint64(omegaInv), uint64(mod)))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([65535, 65535, 65535, 65535]))", func() {
		A := []uint16{65535, 65535, 65535, 65535}

		mod, omega, omegaInv, err := findModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT2(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT2a(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([0, 1, 3, 15, 31, 63, 127, 255, 511, 1023, 2047, 4095, 8191, 16383, 32767, 65535]))", func() {
		A := []uint16{0, 1, 3, 15, 31, 63, 127, 255, 511, 1023, 2047, 4095, 8191, 16383, 32767, 65535}

		mod, omega, omegaInv, err := findModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NTT2(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := INTT2a(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("[4,1,4,2,1,3,5,6] -> [26,338,228,115,2,457,437,448]", func() {
		A := []uint16{4, 1, 4, 2, 1, 3, 5, 6}

		mod, omega, omegaInv, err := modular.FindModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		result := INTT2a(NTT2(A, uint64(omega), uint64(mod)), uint64(omegaInv), uint64(mod))

		Expect(result).To(Equal([]uint16{4, 1, 4, 2, 1, 3, 5, 6}))
	})
})
