package ntt

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go-schoenhageStrassen/modular"
)

var _ = Describe("INTT", func() {
	It("[3,7,0,5,4] -> [6,0,10,7,2]", func() {
		A := []uint64{3, 7, 0, 5, 4}

		omega := uint64(4)
		mod := uint64(11)

		result := InttNaive(A, omega, mod)

		Expect(result).To(Equal([]uint16{6, 0, 10, 7, 2}))
	})

	It("INTT(NTT([0, 1]))", func() {
		A := []uint16{0, 1}

		mod, omega, omegaInv, err := modular.FindModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NttNaive(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := InttNaive(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([0, 1, 0, 0]))", func() {
		A := []uint16{0, 1, 0, 0}

		mod, omega, omegaInv, err := modular.FindModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NttNaive(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := InttNaive(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 1, 1, 1]))", func() {
		A := []uint16{1, 1, 1, 1}

		mod, omega, omegaInv, err := modular.FindModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NttNaive(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := InttNaive(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 1, 1]))", func() {
		A := []uint16{1, 1, 1}

		mod, omega, omegaInv, err := modular.FindModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NttNaive(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := InttNaive(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([10, 10, 10, 10]))", func() {
		A := []uint16{10, 10, 10, 10}

		mod, omega, omegaInv, err := modular.FindModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NttNaive(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := InttNaive(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([100, 100, 100, 100]))", func() {
		A := []uint16{100, 100, 100, 100}

		mod, omega, omegaInv, err := modular.FindModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NttNaive(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := InttNaive(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1000, 1000, 1000, 1000]))", func() {
		A := []uint16{1000, 1000, 1000, 1000}

		mod, omega, omegaInv, err := modular.FindModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NttNaive(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := InttNaive(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([10000, 10000, 10000, 10000]))", func() {
		A := []uint16{10000, 10000, 10000, 10000}

		mod, omega, omegaInv, err := modular.FindModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NttNaive(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := InttNaive(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([1, 10, 100, 1000, 10000]))", func() {
		A := []uint16{1, 10, 100, 1000, 10000}

		mod, omega, omegaInv, err := modular.FindModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NttNaive(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := InttNaive(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([65535, 65535, 65535, 65535]))", func() {
		A := []uint16{65535, 65535, 65535, 65535}

		mod, omega, omegaInv, err := modular.FindModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NttNaive(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := InttNaive(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})

	It("INTT(NTT([0, 1, 3, 15, 31, 63, 127, 255, 511, 1023, 2047, 4095, 8191, 16383, 32767, 65535]))", func() {
		A := []uint16{0, 1, 3, 15, 31, 63, 127, 255, 511, 1023, 2047, 4095, 8191, 16383, 32767, 65535}

		mod, omega, omegaInv, err := modular.FindModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		A_ntt := NttNaive(A, uint64(omega), uint64(mod))
		Expect(A_ntt).NotTo(Equal(A))

		A_recovered := InttNaive(A_ntt, uint64(omegaInv), uint64(mod))
		Expect(A_recovered).To(Equal(A))
	})
})
