package arithmetic_test

import (
	. "github.com/borisskert/go-schoenhageStrassen/arithmetic"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("modSub", func() {
	It("(4 - 1) mod 11", func() {
		mod := uint64(11)
		a := uint64(4)
		b := uint64(1)

		result := ModSub(a, b, mod)
		Expect(result).To(BeNumerically("==", 3))
	})

	It("(4 - 3) mod 11", func() {
		mod := uint64(11)
		a := uint64(4)
		b := uint64(3)

		result := ModSub(a, b, mod)
		Expect(result).To(BeNumerically("==", 1))
	})

	It("(4 - 4) mod 11", func() {
		mod := uint64(11)
		a := uint64(4)
		b := uint64(4)

		result := ModSub(a, b, mod)
		Expect(result).To(BeNumerically("==", 0))
	})

	It("(4 - 5) mod 11", func() {
		mod := uint64(11)
		a := uint64(4)
		b := uint64(5)

		result := ModSub(a, b, mod)
		Expect(result).To(BeNumerically("==", 10))
	})

	It("(4 - 7) mod 11", func() {
		mod := uint64(11)
		a := uint64(4)
		b := uint64(7)

		result := ModSub(a, b, mod)
		Expect(result).To(BeNumerically("==", 8))
	})

	It("(4 - 11) mod 11", func() {
		mod := uint64(11)
		a := uint64(4)
		b := uint64(11)

		result := ModSub(a, b, mod)
		Expect(result).To(BeNumerically("==", 4))
	})

	It("(4 - 21) mod 11", func() {
		mod := uint64(11)
		a := uint64(4)
		b := uint64(21)

		result := ModSub(a, b, mod)
		Expect(result).To(BeNumerically("==", 5))
	})

	It("(23 - 2) mod 11", func() {
		mod := uint64(11)
		a := uint64(23)
		b := uint64(2)

		result := ModSub(a, b, mod)
		Expect(result).To(BeNumerically("==", 10))
	})

	It("(1 - 4294967295) mod 4294967311", func() {
		mod := uint64(4294967311)
		a := uint64(1)
		b := uint64(4294967295)

		result := ModSub(a, b, mod)
		Expect(result).To(BeNumerically("==", 17))
	})
})
