package schoenhageStrassen

import (
	. "github.com/borisskert/go-schoenhageStrassen/arithmetic"
	"github.com/borisskert/go-schoenhageStrassen/array"
	"github.com/borisskert/go-schoenhageStrassen/misc"
	"github.com/borisskert/go-schoenhageStrassen/modulus"
	NTT "github.com/borisskert/go-schoenhageStrassen/ntt"
)

type SchoenhageStrassen struct {
	ntt NTT.NumberTheoreticTransforms
}

func (s SchoenhageStrassen) Multiply16(a, b []uint16) []uint16 {
	return multiply16(a, b, s.ntt.NTT, s.ntt.INTT)
}

func (s SchoenhageStrassen) Multiply32(a, b []uint32) []uint32 {
	return multiply32(a, b, s.ntt.NTT, s.ntt.INTT)
}

func (s SchoenhageStrassen) Multiply64(a, b []uint64) []uint64 {
	return multiply64(a, b, s.ntt.NTT, s.ntt.INTT)
}

func NewSchoenhageStrassen(ntt NTT.NumberTheoreticTransforms) *SchoenhageStrassen {
	return &SchoenhageStrassen{
		ntt: ntt,
	}
}

func multiply16(
	a, b []uint16,
	ntt func([]uint16, uint64, uint64) []uint64,
	intt func([]uint64, uint64, uint64) []uint16,
) []uint16 {
	n := misc.NextPowerOf2(len(a) + len(b))
	aPadded := array.Pad16(a, n)
	bPadded := array.Pad16(b, n)

	// Choose a suitable modulus and primitive root of unity
	mod, omega, omegaInv, err := modulus.FindModulusForTwo16(aPadded, bPadded)
	if err != nil {
		panic(err)
	}

	nttA := ntt(aPadded, omega, mod)
	nttB := ntt(bPadded, omega, mod)

	// Pointwise multiplication in NTT domain
	productNTT := make([]uint64, n)
	for i := 0; i < n; i++ {
		productNTT[i] = ModMul(nttA[i], nttB[i], mod)
	}

	result := intt(productNTT, omegaInv, mod)

	return array.TrimLeadingZeros16(result)
}

func multiply32(
	a, b []uint32,
	ntt func([]uint16, uint64, uint64) []uint64,
	intt func([]uint64, uint64, uint64) []uint16,
) []uint32 {
	a32 := array.Split32To16(a)
	b32 := array.Split32To16(b)

	result16 := multiply16(a32, b32, ntt, intt)
	result32 := array.Combine16To32(result16)

	return array.TrimLeadingZeros32(result32)
}

func multiply64(
	a, b []uint64,
	ntt func([]uint16, uint64, uint64) []uint64,
	intt func([]uint64, uint64, uint64) []uint16,
) []uint64 {
	a16 := array.Split64To16(a)
	b16 := array.Split64To16(b)

	result16 := multiply16(a16, b16, ntt, intt)
	result64 := array.Combine16To64(result16)

	return array.TrimLeadingZeros64(result64)
}
