package array

func Combine16To64(arr []uint16) []uint64 {
	arr = append(arr, make([]uint16, 4-len(arr)%4)...)

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

func Combine16To32(arr []uint16) []uint32 {
	arr = append(arr, make([]uint16, 2-len(arr)%2)...)

	result := make([]uint32, len(arr)/2)
	for i := 0; i < len(arr); i += 2 {
		x0 := uint32(arr[i])
		x1 := uint32(arr[i+1])
		result[i/2] = x0 | (x1 << 16)
	}
	return result
}
