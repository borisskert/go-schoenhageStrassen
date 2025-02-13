package arithmetic_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "go-schoenhageStrassen/arithmetic"
)

var _ = Describe("modInverse", func() {
	It("inv 3 mod 11", func() {
		mod := uint64(11)
		n := uint64(3)

		inverse := ModInverse(n, mod)
		Expect(inverse * n % mod).To(BeNumerically("==", 1))
	})

	It("inv 8 mod 17", func() {
		mod := uint64(17)
		n := uint64(8)

		inverse := ModInverse(n, mod)
		Expect(inverse * n % mod).To(BeNumerically("==", 1))
	})

	It("inv 8 mod 4294967297", func() {
		mod := uint64(4294967297)
		n := uint64(8)

		inverse := ModInverse(n, mod)
		Expect(inverse * n % mod).To(BeNumerically("==", 1))
	})

	It("inv 4 mod 4294967297", func() {
		mod := uint64(4294967297)
		n := uint64(4)

		inverse := ModInverse(n, mod)
		Expect(ModMul(inverse, n, mod)).To(BeNumerically("==", 1))
	})

	It("inv 3 mod 4294967297", func() {
		mod := uint64(4294967297)
		n := uint64(3)

		inverse := ModInverse(n, mod)
		Expect(inverse * n % mod).To(BeNumerically("==", 1))
	})

	It("check 3221225473 * 4 mod 4294967297", func() {
		mod := uint64(4294967297)
		n := uint64(4)

		inverse := uint64(3221225473)
		Expect(inverse * n % mod).To(BeNumerically("==", 1))
	})

	It("check 1431655766 * 3 mod 4294967297", func() {
		mod := uint64(4294967297)
		n := uint64(3)

		inverse := uint64(1431655766)
		Expect(inverse * n % mod).To(BeNumerically("==", 1))
	})
})
