package modulus

import (
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("precompute modulus", func() {
	It("Just do it", func() {
		precomputeWorkingModulusParallel()
	})
})
