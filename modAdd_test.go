package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("modAdd", func() {
	It("(3 + 4) mod 11", func() {
		mod := int64(11)
		a := int64(3)
		b := int64(4)

		result := modAdd(a, b, mod)
		Expect(result).To(BeNumerically("==", 7))
	})

	It("(6 + 4) mod 11", func() {
		mod := int64(11)
		a := int64(6)
		b := int64(4)

		result := modAdd(a, b, mod)
		Expect(result).To(BeNumerically("==", 10))
	})

	It("(6 + 7) mod 11", func() {
		mod := int64(11)
		a := int64(6)
		b := int64(7)

		result := modAdd(a, b, mod)
		Expect(result).To(BeNumerically("==", 2))
	})

	It("(10 + 10) mod 11", func() {
		mod := int64(11)
		a := int64(10)
		b := int64(10)

		result := modAdd(a, b, mod)
		Expect(result).To(BeNumerically("==", 9))
	})

	It("(11 + 10) mod 11", func() {
		mod := int64(11)
		a := int64(11)
		b := int64(10)

		result := modAdd(a, b, mod)
		Expect(result).To(BeNumerically("==", 10))
	})

	It("(11 + 10) mod 11", func() {
		mod := int64(11)
		a := int64(11)
		b := int64(11)

		result := modAdd(a, b, mod)
		Expect(result).To(BeNumerically("==", 0))
	})
})
