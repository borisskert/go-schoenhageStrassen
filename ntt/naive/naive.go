package naive

import "github.com/borisskert/go-schoenhageStrassen/ntt"

type naiveNTT struct{}

func (n naiveNTT) NTT(a []uint16, omega uint64, mod uint64) []uint64 {
	return nttNaive(a, omega, mod)
}

func (n naiveNTT) INTT(a []uint64, omega uint64, mod uint64) []uint16 {
	return inttNaive(a, omega, mod)
}

func NewNaiveNTT() ntt.NumberTheoreticTransforms {
	return &naiveNTT{}
}
