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

func GreatestAdjacentProduct(series string, size int) ([]int, uint64) {
	var i int
	var max uint64 = 0
	list := make([]int, size)
	maxList := make([]int, size)
	for _, n := range series {
		i = int(n) - '0'

		product := uint64(1)
		for _, num := range list {
			product *= uint64(num)
		}
		if product > max {
			copy(maxList, list)
			max = product
		}
		list, _ = Shift(list)
		list[len(list)-1] = i
	}
	return maxList, max
}
