package helpers

import (
	"math"
)

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

// PythagoreanTriplet checks if the provided integers are a Pythagorean Triplet.
//
// Returns true if they are, false otherwise
func PythagoreanTriplet(a int, b int, c int) bool {
	return math.Pow(float64(a), 2)+math.Pow(float64(b), 2) ==
		math.Pow(float64(c), 2)
}

// NaturalNSumToY calculates a list of n-size groups of Natural numbers that sum to y.
//
// For the purposes of this function, the set of Natural numbers excludes zero.
// The returned array contains no duplicates. [1, 1, 2] and [1, 2, 1] are
// considered the same.
func NaturalNSumToY(n int, y int) [][]int {
	result := make([][]int, 0)
	for i := 1; i <= y; i++ {
		for j := 1; j <= y; j++ {
			for k := 1; k <= y; k++ {
				if i+j+k == y {
					result = append(result, []int{i, j, k})
				}
			}
		}
	}
	return result
}

// GetTriangleNumber returns the nth Triangle Number
func GetTriangleNumber(n int) uint64 {
	return SumRange(uint64(1), uint64(n))
}
