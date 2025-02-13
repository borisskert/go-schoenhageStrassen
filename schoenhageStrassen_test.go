package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("schoenhageStrassenMultiply (32bit)", func() {
	It("[0, 1] * [0, 1]", func() {
		A := []uint32{0, 1}
		B := []uint32{0, 1}

		result := schoenhageStrassenMultiply32(A, B)

		Expect(result).To(Equal([]uint32{0, 0, 1}))
	})

	It("[0, 0, 1] * [0, 0, 1]", func() {
		A := []uint32{0, 0, 1}
		B := []uint32{0, 0, 1}

		result := schoenhageStrassenMultiply32(A, B)

		Expect(result).To(Equal([]uint32{0, 0, 0, 0, 1}))
	})

	It("[0, 0, 0, 1] * [0, 0, 0, 1]", func() {
		A := []uint32{0, 0, 0, 1}
		B := []uint32{0, 0, 0, 1}

		result := schoenhageStrassenMultiply32(A, B)

		Expect(result).To(Equal([]uint32{0, 0, 0, 0, 0, 0, 1}))
	})

	It("[0, 0, 0, 1] * [0, 0, 1]", func() {
		A := []uint32{0, 0, 0, 1}
		B := []uint32{0, 0, 1}

		result := schoenhageStrassenMultiply32(A, B)

		Expect(result).To(Equal([]uint32{0, 0, 0, 0, 0, 1}))
	})

	It("[0, 0, 0, 4] * [1, 0, 0, 1]", func() {
		A := []uint32{0, 0, 0, 4}
		B := []uint32{1, 0, 0, 1}

		result := schoenhageStrassenMultiply32(A, B)

		Expect(result).To(Equal([]uint32{0, 0, 0, 4, 0, 0, 4}))
	})

	It("[1, 2, 3, 4] * [1, 2, 3, 4]", func() {
		A := []uint32{1, 2, 3, 4}
		B := []uint32{1, 2, 3, 4}

		result := schoenhageStrassenMultiply32(A, B)

		Expect(result).To(Equal([]uint32{1, 4, 10, 20, 25, 24, 16})) // TODO is this correct?
	})

	It("[0, 2147483647] * [0, 2]", func() {
		A := []uint32{0, 2147483647}
		B := []uint32{0, 2}

		result := schoenhageStrassenMultiply32(A, B)

		Expect(result).To(Equal([]uint32{0, 0, 4294967294}))
	})

	It("[0, 4294967295] * [0, 2]", func() {
		A := []uint32{0, 4294967295}
		B := []uint32{0, 2}

		result := schoenhageStrassenMultiply32(A, B)

		Expect(result).To(Equal([]uint32{0, 0, 4294967294, 1}))
	})

	It("[0, 4294967295] * [0, 4294967295]", func() {
		A := []uint32{0, 4294967295}
		B := []uint32{0, 4294967295}

		result := schoenhageStrassenMultiply32(A, B)

		Expect(result).To(Equal([]uint32{0, 0, 1, 4294967294}))
	})
})

var _ = Describe("schoenhageStrassenMultiply (16bit)", func() {
	It("[0, 1] * [0, 1]", func() {
		A := []uint16{0, 1}
		B := []uint16{0, 1}

		result := schoenhageStrassenMultiply16(A, B)

		Expect(result).To(Equal([]uint16{0, 0, 1}))
	})

	It("[0, 1] * [0, 1] (a)", func() {
		A := []uint16{0, 1}
		B := []uint16{0, 1}

		result := schoenhageStrassenMultiply16a(A, B)

		Expect(result).To(Equal([]uint16{0, 0, 1}))
	})

	It("[0, 0, 1] * [0, 0, 1]", func() {
		A := []uint16{0, 0, 1}
		B := []uint16{0, 0, 1}

		result := schoenhageStrassenMultiply16(A, B)

		Expect(result).To(Equal([]uint16{0, 0, 0, 0, 1}))
	})

	It("[0, 0, 0, 1] * [0, 0, 0, 1]", func() {
		A := []uint16{0, 0, 0, 1}
		B := []uint16{0, 0, 0, 1}

		result := schoenhageStrassenMultiply16(A, B)

		Expect(result).To(Equal([]uint16{0, 0, 0, 0, 0, 0, 1}))
	})

	It("[0, 0, 0, 1] * [0, 0, 1]", func() {
		A := []uint16{0, 0, 0, 1}
		B := []uint16{0, 0, 1}

		result := schoenhageStrassenMultiply16(A, B)

		Expect(result).To(Equal([]uint16{0, 0, 0, 0, 0, 1}))
	})

	It("[0, 0, 0, 4] * [1, 0, 0, 1]", func() {
		A := []uint16{0, 0, 0, 4}
		B := []uint16{1, 0, 0, 1}

		result := schoenhageStrassenMultiply16(A, B)

		Expect(result).To(Equal([]uint16{0, 0, 0, 4, 0, 0, 4}))
	})

	It("[0, 0, 0, 4] * [1, 0, 0, 1] (a)", func() {
		A := []uint16{0, 0, 0, 4}
		B := []uint16{1, 0, 0, 1}

		result := schoenhageStrassenMultiply16a(A, B)

		Expect(result).To(Equal([]uint16{0, 0, 0, 4, 0, 0, 4}))
	})

	It("[0, 0, 0, 4] * [2, 0, 0, 1] (a)", func() {
		A := []uint16{0, 0, 0, 4}
		B := []uint16{2, 0, 0, 1}

		result := schoenhageStrassenMultiply16a(A, B)

		Expect(result).To(Equal([]uint16{0, 0, 0, 8, 0, 0, 4}))
	})

	It("[1, 2, 3, 4] * [1, 2, 3, 4]", func() {
		A := []uint16{1, 2, 3, 4}
		B := []uint16{1, 2, 3, 4}

		result := schoenhageStrassenMultiply16(A, B)

		Expect(result).To(Equal([]uint16{1, 4, 10, 20, 25, 24, 16})) // TODO is this correct?
	})

	It("[1, 2, 3, 4] * [1, 2, 3, 4] (a)", func() {
		A := []uint16{1, 2, 3, 4}
		B := []uint16{1, 2, 3, 4}

		result := schoenhageStrassenMultiply16a(A, B)

		Expect(result).To(Equal([]uint16{1, 4, 10, 20, 25, 24, 16})) // TODO is this correct?
	})

	It("[0, 32767] * [0, 2]", func() {
		A := []uint16{0, 32767}
		B := []uint16{0, 2}

		result := schoenhageStrassenMultiply16(A, B)

		Expect(result).To(Equal([]uint16{0, 0, 65534}))
	})

	It("[0, 65535] * [0, 2]", func() {
		A := []uint16{0, 65535}
		B := []uint16{0, 2}

		result := schoenhageStrassenMultiply16(A, B)

		Expect(result).To(Equal([]uint16{0, 0, 65534, 1}))
	})

	It("[0, 65535] * [0, 65535]", func() {
		A := []uint16{0, 65535}
		B := []uint16{0, 65535}

		result := schoenhageStrassenMultiply16(A, B)

		Expect(result).To(Equal([]uint16{0, 0, 1, 65534}))
	})

	It("[0, 65535] * [0, 65535]", func() {
		A := []uint16{65535, 65535}
		B := []uint16{0, 65535}

		result := schoenhageStrassenMultiply16(A, B)

		Expect(result).To(Equal([]uint16{0, 1, 65535, 65534}))
	})
})

var _ = Describe("schoenhageStrassenMultiply (64bit)", func() {
	It("[0, 1] * [0, 1]", func() {
		A := []uint64{0, 1}
		B := []uint64{0, 1}

		result := schoenhageStrassenMultiply64(A, B)

		Expect(result).To(Equal([]uint64{0, 0, 1}))
	})

	It("[0, 4294967295] * [0, 4294967295]", func() {
		A := []uint64{0, 65535}
		B := []uint64{0, 65535}

		result := schoenhageStrassenMultiply64(A, B)

		Expect(result).To(Equal([]uint64{0, 0, 1, 65534}))
	})
})
