package array

func Split64To16(arr []uint64) []uint16 {
	result := make([]uint16, len(arr)*4)
	for i, x := range arr {
		result[4*i] = uint16(x & 0xFFFF)           // Least significant 16 bits
		result[4*i+1] = uint16((x >> 16) & 0xFFFF) // Next 16 bits
		result[4*i+2] = uint16((x >> 32) & 0xFFFF) // Next 16 bits
		result[4*i+3] = uint16((x >> 48) & 0xFFFF) // Most significant 16 bits
	}
	return result
}

func Split32To16(arr []uint32) []uint16 {
	result := make([]uint16, len(arr)*2)
	for i, x := range arr {
		result[2*i] = uint16(x & 0xFFFF)           // Least significant 16 bits
		result[2*i+1] = uint16((x >> 16) & 0xFFFF) // Next 16 bits
	}
	return result
}
