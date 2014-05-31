package helpers

func SumRange(lower uint64, upper uint64) uint64 {
	var i uint64
	var sum uint64 = 0
	for i = lower; i <= upper; i++ {
		sum += i
	}
	return sum
}

func SumRangeSquares(lower uint64, upper uint64) uint64 {
	var i uint64
	var sum uint64 = 0
	for i = 0; i <= upper; i++ {
		sum += i * i
	}
	return sum
}
