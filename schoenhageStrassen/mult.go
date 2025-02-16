package schoenhageStrassen

import (
	"fmt"
	. "go-schoenhageStrassen/arithmetic"
	"go-schoenhageStrassen/array"
	"go-schoenhageStrassen/misc"
	"go-schoenhageStrassen/modular"
	"go-schoenhageStrassen/ntt/cooleyTukey"
)

// var ntt = cooleyTukey.NewCooleyTukeyIterative()
var ntt = cooleyTukey.NewCooleyTukeyRecursive()

func Multiply16(a, b []uint16) []uint16 {
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	n := misc.NextPowerOf2a(len(a) + len(b))
	aPadded := array.Pad16(a, n)
	bPadded := array.Pad16(b, n)

	// Choose a suitable modulus and primitive root of unity
	mod, omega, omegaInv, err := modular.FindModulusForTwo16(aPadded, bPadded)
	if err != nil {
		panic(err)
	}

	nttA := ntt.NTT(aPadded, uint64(omega), uint64(mod))
	nttB := ntt.NTT(bPadded, uint64(omega), uint64(mod))

	fmt.Println("nttA:", nttA)
	fmt.Println("nttB:", nttB)

	// Pointwise multiplication in NTT domain
	productNTT := make([]uint64, n)
	for i := 0; i < n; i++ {
		productNTT[i] = ModMul(nttA[i], nttB[i], uint64(mod))
	}

	fmt.Println("productNTT:", productNTT)

	result := ntt.INTT(productNTT, uint64(omegaInv), uint64(mod))

	fmt.Println("result (INTT):", result)

	return array.TrimLeadingZeros16(result)
}

func Multiply32(a, b []uint32) []uint32 {
	a32 := array.Split32To16(a)
	b32 := array.Split32To16(b)

	result16 := Multiply16(a32, b32)
	result32 := array.Combine16To32(result16)

	return array.TrimLeadingZeros32(result32)
}

func Multiply64(a, b []uint64) []uint64 {
	a16 := array.Split64To16(a)
	b16 := array.Split64To16(b)

	result16 := Multiply16(a16, b16)
	result64 := array.Combine16To64(result16)

	return array.TrimLeadingZeros64(result64)
}
