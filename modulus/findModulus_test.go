package modulus

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("findModulusMN", func() {
	It("Should fail for n = 1", func() {
		modulus, omega, omegaInv, err := findModulusMN(65535, 1)

		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("n must be greater than 1"))
		Expect(modulus).To(BeNil())
		Expect(omega).To(BeNil())
		Expect(omegaInv).To(BeNil())
	})

	It("for m = 65535, n = 2", func() {
		modulus, omega, omegaInv, err := findModulusMN(65535, 2)

		Expect(err).NotTo(HaveOccurred())
		Expect(*modulus).To(BeNumerically("==", 131071))
		Expect(*omega).To(BeNumerically("==", 131070))
		Expect(*omegaInv).To(BeNumerically("==", 131070))
	})

	It("for m = 65535, n = 3", func() {
		modulus, omega, omegaInv, err := findModulusMN(65535, 3)

		Expect(err).NotTo(HaveOccurred())
		Expect(*modulus).To(BeNumerically("==", 196657))
		Expect(*omega).To(BeNumerically("==", 70289))
		Expect(*omegaInv).To(BeNumerically("==", 126367))
	})

	It("for m = 65535 * 65535, n = 2", func() {
		modulus, omega, omegaInv, err := findModulusMN(65535*65535, 2)

		Expect(err).NotTo(HaveOccurred())
		Expect(*modulus).To(BeNumerically("==", 8589672499))
		Expect(*omega).To(BeNumerically("==", 8589672498))
		Expect(*omegaInv).To(BeNumerically("==", 8589672498))
	})

	It("for m = 65535 * 65535, n = 10", func() {
		modulus, omega, omegaInv, err := findModulusMN(65535*65535, 10)

		Expect(err).NotTo(HaveOccurred())
		Expect(*modulus).To(BeNumerically("==", 42948362251))
		Expect(*omega).To(BeNumerically("==", 26456647984))
		Expect(*omegaInv).To(BeNumerically("==", 24741851202))
	})

	It("for m = 65535 * 65535, n = 100", func() {
		modulus, omega, omegaInv, err := findModulusMN(65535*65535, 100)

		Expect(err).NotTo(HaveOccurred())
		Expect(*modulus).To(BeNumerically("==", 429483625201))
		Expect(*omega).To(BeNumerically("==", 74977269637))
		Expect(*omegaInv).To(BeNumerically("==", 97607101009))
	})

	It("for m = 65535 * 65535, n = 1000", func() {
		modulus, omega, omegaInv, err := findModulusMN(65535*65535, 1000)

		Expect(err).NotTo(HaveOccurred())
		Expect(*modulus).To(BeNumerically("==", 4294836241001))
		Expect(*omega).To(BeNumerically("==", 1968671582207))
		Expect(*omegaInv).To(BeNumerically("==", 4069846287939))
	})

	It("for m = 65535 * 65535, n = 10000", func() {
		modulus, omega, omegaInv, err := findModulusMN(65535*65535, 10000)

		Expect(err).NotTo(HaveOccurred())
		Expect(*modulus).To(BeNumerically("==", 42948362280001))
		Expect(*omega).To(BeNumerically("==", 5094882205187))
		Expect(*omegaInv).To(BeNumerically("==", 7328359974866))
	})

	It("for m = 65535 * 65535, n = 100000", func() {
		modulus, omega, omegaInv, err := findModulusMN(65535*65535, 100000)

		Expect(err).NotTo(HaveOccurred())
		Expect(*modulus).To(BeNumerically("==", 429483622900001))
		Expect(*omega).To(BeNumerically("==", 196664381113895))
		Expect(*omegaInv).To(BeNumerically("==", 45480248106017))
	})

	It("for m = 65535 * 65535, n = 1000000", func() {
		modulus, omega, omegaInv, err := findModulusMN(65535*65535, 1000000)

		Expect(err).NotTo(HaveOccurred())
		Expect(*modulus).To(BeNumerically("==", 4294836228000001))
		Expect(*omega).To(BeNumerically("==", 2798360728920501))
		Expect(*omegaInv).To(BeNumerically("==", 912339468719032))
	})
})
