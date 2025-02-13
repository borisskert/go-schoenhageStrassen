package main

func copy32to64N(left []uint32, n int) []uint64 {
	size := max(len(left), n)
	right := make([]uint64, size)

	for i := 0; i < len(left); i++ {
		right[i] = uint64(left[i])
	}

	return right
}

func copy16to64N(left []uint16, n int) []uint64 {
	size := max(len(left), n)
	right := make([]uint64, size)

	for i := 0; i < len(left); i++ {
		right[i] = uint64(left[i])
	}

	return right
}

func copy64to16(left []uint64) []uint16 {
	size := len(left)
	right := make([]uint16, size)

	for i := 0; i < len(left); i++ {
		if left[i] > 0xFFFF {
			panic("value too large to fit in 16 bits")
		}

		right[i] = uint16(left[i])
	}

	return right
}

func copy32to64(left []uint32) []uint64 {
	right := make([]uint64, len(left))

	for i := 0; i < len(left); i++ {
		right[i] = uint64(left[i])
	}

	return right
}
