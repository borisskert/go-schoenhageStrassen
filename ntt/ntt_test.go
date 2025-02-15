package ntt

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ntt Package Test Suite")
}

var _ = Describe("NTT", func() {
	It("[6, 0, 10, 7, 2] -> [3, 7, 0, 5, 4]", func() {
		A := []uint16{6, 0, 10, 7, 2}

		omega := uint64(3)
		mod := uint64(11)

		result := NaiveNTT(A, omega, mod)

		Expect(result).To(Equal([]uint64{3, 7, 0, 5, 4}))
	})

	It("[4,1,4,2,1,3,5,6] -> [26,338,228,115,2,457,437,448]", func() {
		A := []uint16{4, 1, 4, 2, 1, 3, 5, 6}

		omega := uint64(326)
		mod := uint64(673)

		result := NaiveNTT(A, omega, mod)

		Expect(result).To(Equal([]uint64{26, 338, 228, 115, 2, 457, 437, 448}))
	})

	It("[6,1,8,0,3,3,9,8] -> [38,594,224,157,14,201,433,406]", func() {
		A := []uint16{6, 1, 8, 0, 3, 3, 9, 8}

		omega := uint64(326)
		mod := uint64(673)

		result := NaiveNTT(A, omega, mod)

		Expect(result).To(Equal([]uint64{38, 594, 224, 157, 14, 201, 433, 406}))
	})

	It("[4, 3, 2, 1, 0, 0, 0, 0] =(85, 337)=> [10, 329, 298, 126, 2, 271, 43, 301]", func() {
		A := []uint16{4, 3, 2, 1, 0, 0, 0, 0}

		omega := uint64(85)
		mod := uint64(337)

		result := NaiveNTT(A, omega, mod)

		Expect(result).To(Equal([]uint64{10, 329, 298, 126, 2, 271, 43, 301}))
	})

	It("[8, 7, 6, 5, 0, 0, 0, 0] =(85, 337)=> [26, 24, 298, 322, 2, 83, 43, 277]", func() {
		A := []uint16{8, 7, 6, 5, 0, 0, 0, 0}

		omega := uint64(85)
		mod := uint64(337)

		result := NaiveNTT(A, omega, mod)

		Expect(result).To(Equal([]uint64{26, 24, 298, 322, 2, 83, 43, 277}))
	})
})
