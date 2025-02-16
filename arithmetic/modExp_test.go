package arithmetic_test

import (
	. "github.com/borisskert/go-schoenhageStrassen/arithmetic"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("modExp", func() {
	It("(3884671324 pow 4294967355) mod 4294967357", func() {
		a := uint64(3884671324)
		b := uint64(4294967355)
		mod := uint64(4294967357)

		result := ModExp(a, b, mod)
		Expect(result).To(BeNumerically("==", 410296033))
	})

	It("(4294967295 pow 4294967295) mod 4294967311", func() {
		a := uint64(4294967295)
		b := uint64(4294967295)
		mod := uint64(4294967311)

		result := ModExp(a, b, mod)
		Expect(result).To(BeNumerically("==", 2023406822))
	})
})
