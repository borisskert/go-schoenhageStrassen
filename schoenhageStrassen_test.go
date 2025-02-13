package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("schoenhageStrassenMultiply", func() {
	It("[0, 1] * [0, 1]", func() {
		A := []int32{0, 1}
		B := []int32{0, 1}

		result := schoenhageStrassenMultiply(A, B)

		Expect(result).To(Equal([]int32{0, 0, 1}))
	})

	It("[0, 0, 1] * [0, 0, 1]", func() {
		A := []int32{0, 0, 1}
		B := []int32{0, 0, 1}

		result := schoenhageStrassenMultiply(A, B)

		Expect(result).To(Equal([]int32{0, 0, 0, 0, 1}))
	})

	It("[0, 0, 0, 1] * [0, 0, 0, 1]", func() {
		A := []int32{0, 0, 0, 1}
		B := []int32{0, 0, 0, 1}

		result := schoenhageStrassenMultiply(A, B)

		Expect(result).To(Equal([]int32{0, 0, 0, 0, 0, 0, 1}))
	})

	It("[0, 0, 0, 1] * [0, 0, 1]", func() {
		A := []int32{0, 0, 0, 1}
		B := []int32{0, 0, 1}

		result := schoenhageStrassenMultiply(A, B)

		Expect(result).To(Equal([]int32{0, 0, 0, 0, 0, 1}))
	})

	It("[0, 0, 0, 4] * [1, 0, 0, 1]", func() {
		A := []int32{0, 0, 0, 4}
		B := []int32{1, 0, 0, 1}

		result := schoenhageStrassenMultiply(A, B)

		Expect(result).To(Equal([]int32{0, 0, 0, 4, 0, 0, 4}))
	})

	It("[1, 2, 3, 4] * [1, 2, 3, 4]", func() {
		A := []int32{1, 2, 3, 4}
		B := []int32{1, 2, 3, 4}

		result := schoenhageStrassenMultiply(A, B)

		Expect(result).To(Equal([]int32{1, 4, 10, 20, 25, 24, 16})) // TODO is this correct?
	})

	It("[2147483647] * [2]", func() {
		A := []int32{0, 2147483647}
		B := []int32{0, 2}

		result := schoenhageStrassenMultiply(A, B)

		Expect(result).To(Equal([]int32{2147483646, 1}))
	})
})
