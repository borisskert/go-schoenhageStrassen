package arithmetic_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "go-schoenhageStrassen/arithmetic"
)

var _ = Describe("modMul", func() {
	It("(4 * 1) mod 11", func() {
		mod := uint64(11)
		a := uint64(4)
		b := uint64(1)

		result := ModMul(a, b, mod)
		Expect(result).To(BeNumerically("==", 4))
	})

	It("(4 * 2) mod 11", func() {
		mod := uint64(11)
		a := uint64(4)
		b := uint64(2)

		result := ModMul(a, b, mod)
		Expect(result).To(BeNumerically("==", 8))
	})

	It("(4 * 3) mod 11", func() {
		mod := uint64(11)
		a := uint64(4)
		b := uint64(3)

		result := ModMul(a, b, mod)
		Expect(result).To(BeNumerically("==", 1))
	})

	It("(4 * 5) mod 11", func() {
		mod := uint64(11)
		a := uint64(4)
		b := uint64(5)

		result := ModMul(a, b, mod)
		Expect(result).To(BeNumerically("==", 9))
	})

	It("(9 * 5) mod 11", func() {
		mod := uint64(11)
		a := uint64(9)
		b := uint64(5)

		result := ModMul(a, b, mod)
		Expect(result).To(BeNumerically("==", 1))
	})

	It("(9 * 11) mod 11", func() {
		mod := uint64(11)
		a := uint64(9)
		b := uint64(11)

		result := ModMul(a, b, mod)
		Expect(result).To(BeNumerically("==", 0))
	})

	It("(14 * 23) mod 11", func() {
		mod := uint64(11)
		a := uint64(14)
		b := uint64(23)

		result := ModMul(a, b, mod)
		Expect(result).To(BeNumerically("==", 3))
	})

	It("(253 * 546) mod 11", func() {
		mod := uint64(11)
		a := uint64(253)
		b := uint64(546)

		result := ModMul(a, b, mod)
		Expect(result).To(BeNumerically("==", 0))
	})

	It("(3451 * 3446) mod 11", func() {
		mod := uint64(11)
		a := uint64(3451)
		b := uint64(3446)

		result := ModMul(a, b, mod)
		Expect(result).To(BeNumerically("==", 2))
	})

	It("(4294967295 * 4294967295) mod 4294967311", func() {
		mod := uint64(4294967311)
		a := uint64(4294967295)
		b := uint64(4294967295)

		result := ModMul(a, b, mod)
		Expect(result).To(BeNumerically("==", 256))
	})
})
