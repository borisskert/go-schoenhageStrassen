package main

func splitUint64ToUint16(arr []uint64) []uint16 {
	result := make([]uint16, len(arr)*4)
	for i, x := range arr {
		result[4*i] = uint16(x & 0xFFFF)           // Least significant 16 bits
		result[4*i+1] = uint16((x >> 16) & 0xFFFF) // Next 16 bits
		result[4*i+2] = uint16((x >> 32) & 0xFFFF) // Next 16 bits
		result[4*i+3] = uint16((x >> 48) & 0xFFFF) // Most significant 16 bits
	}
	return result
}

func splitUint32ToUint16(arr []uint32) []uint16 {
	result := make([]uint16, len(arr)*2)
	for i, x := range arr {
		result[2*i] = uint16(x & 0xFFFF)           // Least significant 16 bits
		result[2*i+1] = uint16((x >> 16) & 0xFFFF) // Next 16 bits
	}
	return result
}

func combineUint16ToUint64(arr []uint16) []uint64 {
	if len(arr)%4 != 0 {
		panic("length of arr must be a multiple of 4") // TODO this must be possible to handle
	}
	result := make([]uint64, len(arr)/4)
	for i := 0; i < len(arr); i += 4 {
		x0 := uint64(arr[i])
		x1 := uint64(arr[i+1])
		x2 := uint64(arr[i+2])
		x3 := uint64(arr[i+3])
		result[i/4] = x0 | (x1 << 16) | (x2 << 32) | (x3 << 48)
	}
	return result
}
