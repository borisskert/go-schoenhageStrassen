package cooleyTukey

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go-schoenhageStrassen/modular"
	"go-schoenhageStrassen/ntt"
)

var _ = Describe("CooleyTukey (recursive)", func() {
	var nttAlgorithm ntt.NumberTheoreticTransforms

	BeforeEach(func() {
		nttAlgorithm = RecursiveAlgorithm()
	})

	It("[6, 0, 10, 7, 2] -> [3, 8, 1, 6, 0, 4, 6, 9]", func() {
		A := []uint16{6, 0, 10, 7, 2}

		omega := uint64(3)
		mod := uint64(11)

		result := nttAlgorithm.NTT(A, omega, mod)

		Expect(result).To(Equal([]uint64{3, 8, 1, 6, 0, 4, 6, 9}))
	})

	It("[6, 0, 10, 7, 2, 0, 0, 0] -> [3, 8, 1, 6, 0, 4, 6, 9]", func() {
		A := []uint16{6, 0, 10, 7, 2, 0, 0, 0}

		omega := uint64(3)
		mod := uint64(11)

		result := nttAlgorithm.NTT(A, omega, mod)

		Expect(result).To(Equal([]uint64{3, 8, 1, 6, 0, 4, 6, 9}))
	})

	It("[4,1,4,2,1,3,5,6] -> [26,338,228,115,2,457,437,448]", func() {
		A := []uint16{4, 1, 4, 2, 1, 3, 5, 6}

		omega := uint64(326)
		mod := uint64(673)

		result := nttAlgorithm.NTT(A, omega, mod)

		Expect(result).To(Equal([]uint64{26, 338, 228, 115, 2, 457, 437, 448}))
	})

	It("[6,1,8,0,3,3,9,8] -> [38,594,224,157,14,201,433,406]", func() {
		A := []uint16{6, 1, 8, 0, 3, 3, 9, 8}

		omega := uint64(326)
		mod := uint64(673)

		result := nttAlgorithm.NTT(A, omega, mod)

		Expect(result).To(Equal([]uint64{38, 594, 224, 157, 14, 201, 433, 406}))
	})

	It("[4, 3, 2, 1, 0, 0, 0, 0] =(85, 337)=> [10, 329, 298, 126, 2, 271, 43, 301]", func() {
		A := []uint16{4, 3, 2, 1, 0, 0, 0, 0}

		omega := uint64(85)
		mod := uint64(337)

		result := nttAlgorithm.NTT(A, omega, mod)

		Expect(result).To(Equal([]uint64{10, 329, 298, 126, 2, 271, 43, 301}))
	})

	It("[8, 7, 6, 5, 0, 0, 0, 0] =(85, 337)=> [26, 24, 298, 322, 2, 83, 43, 277]", func() {
		A := []uint16{8, 7, 6, 5, 0, 0, 0, 0}

		omega := uint64(85)
		mod := uint64(337)

		result := nttAlgorithm.NTT(A, omega, mod)

		Expect(result).To(Equal([]uint64{26, 24, 298, 322, 2, 83, 43, 277}))
	})

	It("[4,1,4,2,1,3,5,6] -> [26,338,228,115,2,457,437,448]", func() {
		A := []uint16{4, 1, 4, 2, 1, 3, 5, 6}

		mod, omega, omegaInv, err := modular.FindModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		result := nttAlgorithm.INTT(nttAlgorithm.NTT(A, uint64(omega), uint64(mod)), uint64(omegaInv), uint64(mod))

		Expect(result).To(Equal([]uint16{4, 1, 4, 2, 1, 3, 5, 6}))
	})

	It("INTT(NTT([6, 0, 10, 7, 2]))", func() {
		A := []uint16{6, 0, 10, 7, 2}

		mod, omega, omegaInv, err := modular.FindModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		result := nttAlgorithm.INTT(nttAlgorithm.NTT(A, uint64(omega), uint64(mod)), uint64(omegaInv), uint64(mod))

		Expect(result).To(Equal([]uint16{6, 0, 10, 7, 2, 0, 0, 0}))
	})

	It("INTT(NTT([6, 0, 10, 7, 2, 0, 0, 0]))", func() {
		A := []uint16{6, 0, 10, 7, 2, 0, 0, 0}

		mod, omega, omegaInv, err := modular.FindModulus16(A)

		Expect(err).NotTo(HaveOccurred())

		result := nttAlgorithm.INTT(nttAlgorithm.NTT(A, uint64(omega), uint64(mod)), uint64(omegaInv), uint64(mod))

		Expect(result).To(Equal([]uint16{6, 0, 10, 7, 2, 0, 0, 0}))
	})
})
