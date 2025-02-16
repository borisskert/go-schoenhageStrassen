package ntt

type NumberTheoreticTransforms interface {
	NTT(a []uint16, omega uint64, mod uint64) []uint64
	INTT(a []uint64, omega uint64, mod uint64) []uint16
}
