package array

func TrimLeadingZeros16(a []uint16) []uint16 {
	i := len(a) - 1
	for i >= 0 && a[i] == 0 {
		i--
	}
	if i < 0 {
		return []uint16{0}
	}
	return a[:i+1]
}

func TrimLeadingZeros32(a []uint32) []uint32 {
	i := len(a) - 1
	for i >= 0 && a[i] == 0 {
		i--
	}
	if i < 0 {
		return []uint32{0}
	}
	return a[:i+1]
}

func TrimLeadingZeros64(a []uint64) []uint64 {
	i := len(a) - 1
	for i >= 0 && a[i] == 0 {
		i--
	}
	if i < 0 {
		return []uint64{0}
	}
	return a[:i+1]
}
