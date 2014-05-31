package helpers

func AppendUnique(slice []uint64, i ...uint64) []uint64 {
	for _, ele := range i {
		slice = appendUnique(slice, ele)
	}
	return slice
}

// From http://stackoverflow.com/a/9561388/943833
func appendUnique(slice []uint64, i uint64) []uint64 {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}
