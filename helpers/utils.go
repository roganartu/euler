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

func AppendUniqueString(slice []string, i ...string) []string {
	for _, ele := range i {
		slice = appendUniqueString(slice, ele)
	}
	return slice
}

// From http://stackoverflow.com/a/9561388/943833
func appendUniqueString(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

// Unshift shifts all elements forwards, removing the first element.
//
// Both the new slice and the removed first element are returned.
// A zero is added to the end of the array
func Shift(slice []int) ([]int, int) {
	removed := slice[0]
	shifted := slice[1:len(slice)]
	shifted = append(shifted, 0)
	return shifted, removed
}

func ProductInt(a []int) int {
	product := 1
	for _, x := range a {
		product *= x
	}
	return product
}
