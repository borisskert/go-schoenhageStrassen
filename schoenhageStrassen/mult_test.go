package schoenhageStrassen

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "main Package Test Suite")
}

var _ = Describe("Multiply16", func() {
	It("[4, 3, 2, 1] * [8, 7, 6, 5]", func() {
		A := []uint16{4, 3, 2, 1}
		B := []uint16{8, 7, 6, 5}

		result := Multiply16(A, B)

		Expect(result).To(Equal([]uint16{32, 52, 61, 60, 34, 16, 5}))
	})

	It("[0, 1] * [0, 1]", func() {
		A := []uint16{0, 1}
		B := []uint16{0, 1}

		result := Multiply16(A, B)

		Expect(result).To(Equal([]uint16{0, 0, 1}))
	})

	It("[0, 0, 1] * [0, 0, 1]", func() {
		A := []uint16{0, 0, 1}
		B := []uint16{0, 0, 1}

		result := Multiply16(A, B)

		Expect(result).To(Equal([]uint16{0, 0, 0, 0, 1}))
	})

	It("[0, 0, 0, 1] * [0, 0, 0, 1]", func() {
		A := []uint16{0, 0, 0, 1}
		B := []uint16{0, 0, 0, 1}

		result := Multiply16(A, B)

		Expect(result).To(Equal([]uint16{0, 0, 0, 0, 0, 0, 1}))
	})

	It("[0, 0, 0, 1] * [0, 0, 1]", func() {
		A := []uint16{0, 0, 0, 1}
		B := []uint16{0, 0, 1}

		result := Multiply16(A, B)

		Expect(result).To(Equal([]uint16{0, 0, 0, 0, 0, 1}))
	})

	It("[0, 0, 0, 4] * [1, 0, 0, 1]", func() {
		A := []uint16{0, 0, 0, 4}
		B := []uint16{1, 0, 0, 1}

		result := Multiply16(A, B)

		Expect(result).To(Equal([]uint16{0, 0, 0, 4, 0, 0, 4}))
	})

	It("[1, 2, 3, 4] * [1, 2, 3, 4]", func() {
		A := []uint16{1, 2, 3, 4}
		B := []uint16{1, 2, 3, 4}

		result := Multiply16(A, B)

		Expect(result).To(Equal([]uint16{1, 4, 10, 20, 25, 24, 16})) // TODO is this correct?
	})

	It("[0, 32767] * [0, 2]", func() {
		A := []uint16{0, 32767}
		B := []uint16{0, 2}

		result := Multiply16(A, B)

		Expect(result).To(Equal([]uint16{0, 0, 65534}))
	})

	It("[0, 65535] * [0, 2]", func() {
		A := []uint16{0, 65535}
		B := []uint16{0, 2}

		result := Multiply16(A, B)

		Expect(result).To(Equal([]uint16{0, 0, 65534, 1}))
	})

	It("[0, 65535] * [0, 65535]", func() {
		A := []uint16{0, 65535}
		B := []uint16{0, 65535}

		result := Multiply16(A, B)

		Expect(result).To(Equal([]uint16{0, 0, 1, 65534}))
	})

	It("[65535, 65535] * [0, 65535]", func() {
		A := []uint16{65535, 65535}
		B := []uint16{0, 65535}

		result := Multiply16(A, B)

		Expect(result).To(Equal([]uint16{0, 1, 65535, 65534}))
	})

	It("[65535, 65535] * [65535, 65535]", func() {
		A := []uint16{65535, 65535}
		B := []uint16{65535, 65535}

		result := Multiply16(A, B)

		Expect(result).To(Equal([]uint16{1, 0, 65534, 65535}))
	})

	It("[65535, 65535, 65535] * [65535, 65535]", func() {
		A := []uint16{65535, 65535, 65535}
		B := []uint16{65535, 65535}

		result := Multiply16(A, B)

		Expect(result).To(Equal([]uint16{1, 0, 65535, 65534, 65535}))
	})
})
