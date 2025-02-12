package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("modExp", func() {
	It("(3884671324 pow 4294967355) mod 4294967357", func() {
		a := int64(3884671324)
		b := int64(4294967355)
		mod := int64(4294967357)

		result := modExp(a, b, mod)
		Expect(result).To(BeNumerically("==", 410296033))
	})
})
