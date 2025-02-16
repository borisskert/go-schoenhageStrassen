package arithmetic_test

import (
	. "github.com/borisskert/go-schoenhageStrassen/arithmetic"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("modAdd", func() {
	It("(3 + 4) mod 11", func() {
		mod := uint64(11)
		a := uint64(3)
		b := uint64(4)

		result := ModAdd(a, b, mod)
		Expect(result).To(BeNumerically("==", 7))
	})

	It("(6 + 4) mod 11", func() {
		mod := uint64(11)
		a := uint64(6)
		b := uint64(4)

		result := ModAdd(a, b, mod)
		Expect(result).To(BeNumerically("==", 10))
	})

	It("(6 + 7) mod 11", func() {
		mod := uint64(11)
		a := uint64(6)
		b := uint64(7)

		result := ModAdd(a, b, mod)
		Expect(result).To(BeNumerically("==", 2))
	})

	It("(10 + 10) mod 11", func() {
		mod := uint64(11)
		a := uint64(10)
		b := uint64(10)

		result := ModAdd(a, b, mod)
		Expect(result).To(BeNumerically("==", 9))
	})

	It("(11 + 10) mod 11", func() {
		mod := uint64(11)
		a := uint64(11)
		b := uint64(10)

		result := ModAdd(a, b, mod)
		Expect(result).To(BeNumerically("==", 10))
	})

	It("(11 + 10) mod 11", func() {
		mod := uint64(11)
		a := uint64(11)
		b := uint64(11)

		result := ModAdd(a, b, mod)
		Expect(result).To(BeNumerically("==", 0))
	})

	It("(4294967295 + 4294967295) mod 4294967311", func() {
		mod := uint64(4294967311)
		a := uint64(4294967295)
		b := uint64(4294967295)

		result := ModAdd(a, b, mod)
		Expect(result).To(BeNumerically("==", 4294967279))
	})
})
