package main

func trimZeros16(a []uint16) []uint16 {
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != 0 {
			return a[:i+1]
		}
	}

	return nil
}
