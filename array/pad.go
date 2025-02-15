package array

func Pad16(a []uint16, length int) []uint16 {
	padded := make([]uint16, length)
	copy(padded, a)
	return padded
}
