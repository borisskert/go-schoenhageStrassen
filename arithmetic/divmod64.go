package arithmetic

import "math/bits"

// Efficient 128-bit division and modulo
func divmod64(aHi, aLo, b uint64) (uint64, uint64, uint64) {
	qHi, r := bits.Div64(0, aHi, b) // Divide the high 64 bits
	qLo, r := bits.Div64(r, aLo, b) // Divide the lower 64 bits
	return qHi, qLo, r              // Quotient (ignored), remainder
}
