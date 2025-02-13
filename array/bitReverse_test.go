package array

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("bitReverseInPlace", func() {
	It("[0, 1]", func() {
		A := []uint64{0, 1}

		bitwiseReverseInPlace64(A)

		Expect(A).To(Equal([]uint64{32768, 0}))
	})

	It("[0, 2]", func() {
		A := []uint64{0, 2}

		bitwiseReverseInPlace64(A)

		Expect(A).To(Equal([]uint64{16384, 0}))
	})

	It("[0, 6]", func() {
		A := []uint64{0, 6}

		bitwiseReverseInPlace64(A)

		Expect(A).To(Equal([]uint64{24576, 0}))
	})

	It("[0, 32768]", func() {
		A := []uint64{0, 32768}

		bitwiseReverseInPlace64(A)

		Expect(A).To(Equal([]uint64{1, 0}))
	})

	It("[0, 16384]", func() {
		A := []uint64{0, 16384}

		bitwiseReverseInPlace64(A)

		Expect(A).To(Equal([]uint64{2, 0}))
	})
})
