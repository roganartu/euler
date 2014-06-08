package helpers

import (
	"math"
	"math/big"
	"strconv"
)

func SumRange(lower uint64, upper uint64) uint64 {
	var i uint64
	var sum uint64 = 0
	for i = lower; i <= upper; i++ {
		sum += i
	}
	return sum
}

// SumDigits returns the sum of the digits contained in the given string.
func SumDigits(digits string) uint64 {
	var n int
	var sum uint64 = 0
	for _, i := range digits {
		n, _ = strconv.Atoi(string(i))
		sum += uint64(n)
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

// Collatz recursively calculates the length of the Collatz reduction chain.
func Collatz(n int, cache map[int]int) (map[int]int, int) {
	if cache != nil && cache[n] != 0 {
		return cache, cache[n]
	}

	if n%2 == 0 {
		n /= 2
	} else {
		n = 3*n + 1
	}

	if n == 1 {
		return cache, 1
	}

	var result int
	cache, result = Collatz(n, cache)
	return cache, result + 1
}

// BigFactorial calculates large factorial numbers using math/big.
//
// For ease of type representation, a string is returned.
func BigFactorial(target int) string {
	z := big.NewInt(int64(target))
	for i := target; i > 0; i-- {
		x := big.NewInt(int64(i))
		z.Mul(z, x)
	}
	return z.String()
}
