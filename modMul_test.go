package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("modMul", func() {
	It("(4 * 1) mod 11", func() {
		mod := int64(11)
		a := int64(4)
		b := int64(1)

		result := modMul(a, b, mod)
		Expect(result).To(BeNumerically("==", 4))
	})

	It("(4 * 2) mod 11", func() {
		mod := int64(11)
		a := int64(4)
		b := int64(2)

		result := modMul(a, b, mod)
		Expect(result).To(BeNumerically("==", 8))
	})

	It("(4 * 3) mod 11", func() {
		mod := int64(11)
		a := int64(4)
		b := int64(3)

		result := modMul(a, b, mod)
		Expect(result).To(BeNumerically("==", 1))
	})

	It("(4 * 5) mod 11", func() {
		mod := int64(11)
		a := int64(4)
		b := int64(5)

		result := modMul(a, b, mod)
		Expect(result).To(BeNumerically("==", 9))
	})

	It("(9 * 5) mod 11", func() {
		mod := int64(11)
		a := int64(9)
		b := int64(5)

		result := modMul(a, b, mod)
		Expect(result).To(BeNumerically("==", 1))
	})

	It("(9 * 11) mod 11", func() {
		mod := int64(11)
		a := int64(9)
		b := int64(11)

		result := modMul(a, b, mod)
		Expect(result).To(BeNumerically("==", 0))
	})

	It("(14 * 23) mod 11", func() {
		mod := int64(11)
		a := int64(14)
		b := int64(23)

		result := modMul(a, b, mod)
		Expect(result).To(BeNumerically("==", 3))
	})

	It("(253 * 546) mod 11", func() {
		mod := int64(11)
		a := int64(253)
		b := int64(546)

		result := modMul(a, b, mod)
		Expect(result).To(BeNumerically("==", 0))
	})

	It("(3451 * 3446) mod 11", func() {
		mod := int64(11)
		a := int64(3451)
		b := int64(3446)

		result := modMul(a, b, mod)
		Expect(result).To(BeNumerically("==", 2))
	})
})
