package modulus

import (
	"fmt"
	"sync"
)

type modulusWithK struct {
	modulus uint64
	k       uint64
}

func precomputeWorkingModulus() map[uint64]modulusWithK {
	uint64max := ^uint64(0)

	workingModulusCache := make(map[uint64]modulusWithK)

	m := uint64(^uint16(0)) * uint64(^uint16(0))

	for n := uint64(8388608); n < uint64max; n *= 2 {
		modulus, k := findNextWorkingModulus(m, n)
		fmt.Println(n, modulus)
		workingModulusCache[n] = modulusWithK{modulus, k}
	}

	return workingModulusCache
}

// precomputeWorkingModulus computes modulus values in parallel using goroutines
func precomputeWorkingModulusParallel() map[uint64]modulusWithK {
	//uint64max := ^uint64(0)
	workingModulusCache := make(map[uint64]modulusWithK)
	var mu sync.Mutex
	var wg sync.WaitGroup

	m := uint64(^uint16(0)) * uint64(^uint16(0))

	for i := uint64(2); i < 64; i++ {
		n := uint64(1) << i

		wg.Add(1)
		go func(n uint64) {
			defer wg.Done()
			modulus, k := findNextWorkingModulus(m, n)
			fmt.Println(n, modulus, k)
			mu.Lock()
			workingModulusCache[n] = modulusWithK{modulus, k}
			mu.Unlock()
		}(n)
	}

	wg.Wait()
	return workingModulusCache
}
