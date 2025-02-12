package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("modSub", func() {
	It("(4 - 1) mod 11", func() {
		mod := int64(11)
		a := int64(4)
		b := int64(1)

		result := modSub(a, b, mod)
		Expect(result).To(BeNumerically("==", 3))
	})

	It("(4 - 3) mod 11", func() {
		mod := int64(11)
		a := int64(4)
		b := int64(3)

		result := modSub(a, b, mod)
		Expect(result).To(BeNumerically("==", 1))
	})

	It("(4 - 4) mod 11", func() {
		mod := int64(11)
		a := int64(4)
		b := int64(4)

		result := modSub(a, b, mod)
		Expect(result).To(BeNumerically("==", 0))
	})

	It("(4 - 5) mod 11", func() {
		mod := int64(11)
		a := int64(4)
		b := int64(5)

		result := modSub(a, b, mod)
		Expect(result).To(BeNumerically("==", 10))
	})

	It("(4 - 7) mod 11", func() {
		mod := int64(11)
		a := int64(4)
		b := int64(7)

		result := modSub(a, b, mod)
		Expect(result).To(BeNumerically("==", 8))
	})

	It("(4 - 11) mod 11", func() {
		mod := int64(11)
		a := int64(4)
		b := int64(11)

		result := modSub(a, b, mod)
		Expect(result).To(BeNumerically("==", 4))
	})

	It("(4 - 21) mod 11", func() {
		mod := int64(11)
		a := int64(4)
		b := int64(21)

		result := modSub(a, b, mod)
		Expect(result).To(BeNumerically("==", 5))
	})

	It("(23 - 2) mod 11", func() {
		mod := int64(11)
		a := int64(23)
		b := int64(2)

		result := modSub(a, b, mod)
		Expect(result).To(BeNumerically("==", 10))
	})
})
