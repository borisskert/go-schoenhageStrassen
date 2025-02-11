package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("modInverse", func() {
	It("inv 3 mod 11", func() {
		mod := int64(11)
		n := int64(3)

		inverse := modInverse(n, mod)
		Expect(inverse * n % mod).To(BeNumerically("==", 1))
	})

	It("inv 8 mod 17", func() {
		mod := int64(17)
		n := int64(8)

		inverse := modInverse(n, mod)
		Expect(inverse * n % mod).To(BeNumerically("==", 1))
	})

	It("inv 8 mod 4294967297", func() {
		mod := int64(4294967297)
		n := int64(8)

		inverse := modInverse(n, mod)
		Expect(inverse * n % mod).To(BeNumerically("==", 1))
	})

	It("inv 4 mod 4294967297", func() {
		mod := int64(4294967297)
		n := int64(4)

		inverse := modInverse(n, mod)
		Expect(inverse * n % mod).To(BeNumerically("==", 1))
	})

	It("inv 3 mod 4294967297", func() {
		mod := int64(4294967297)
		n := int64(3)

		inverse := modInverse(n, mod)
		Expect(inverse * n % mod).To(BeNumerically("==", 1))
	})

	It("check 3221225473 * 4 mod 4294967297", func() {
		mod := int64(4294967297)
		n := int64(4)

		inverse := int64(3221225473)
		Expect(inverse * n % mod).To(BeNumerically("==", 1))
	})

	It("check 1431655766 * 3 mod 4294967297", func() {
		mod := int64(4294967297)
		n := int64(3)

		inverse := int64(1431655766)
		Expect(inverse * n % mod).To(BeNumerically("==", 1))
	})
})
