package helpers

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

// Factorise calculates all factors of the given target
func Factorise(target uint64, full bool) []uint64 {
	var i uint64
	factors := make([]uint64, 0)
	for i = 2; i <= uint64(math.Sqrt(float64(target)))+uint64(1); i++ {
		if target%i == 0 && target/i != 1 {
			factors = AppendUnique(factors, target/i, i)
			if !full {
				break
			}
		}
	}
	return factors
}

// ProperDivisors returns all the proper divisors of the given target
func ProperDivisors(target uint64) []uint64 {
	factors := Factorise(target, true)
	return AppendUnique(factors, uint64(1))
}

// SumFactors returns the sum of the factors of target.
func SumFactors(target uint64) uint64 {
	factors := Factorise(target, true)
	sum := uint64(0)
	for _, x := range factors {
		sum += x
	}
	return sum
}

// LexographicPermutations recursively finds the ordered permutations of
// the given range.
func LexographicPermutations(prefix string, start int, end int) []string {
	perms := make([]string, 0)
	for i := start; i <= end; i++ {
		stri := strconv.Itoa(i)

		// Avoid dupes
		if strings.Contains(prefix, stri) {
			continue
		}

		if len(prefix) < end {
			for _, a := range LexographicPermutations(prefix+stri, start, end) {
				perms = append(perms, a)
			}
		} else {
			perms = append(perms, prefix+stri)
		}
	}
	return perms
}

// ShortRepeatedString returns the longest repeated substring in str.
//
// If uses an inefficient implementation (ie: not suffix-tree) and should only
// be used for sufficiently small inputs.
//
// Returns repeated substring if found, otherwise empty string.
func ShortRepeatedString(str string) string {
	for i := len(str) / 2; i < len(str)-1; i++ {
		suffix := str[i:len(str)]
		regex := regexp.MustCompile("" + suffix + suffix + "$")
		if regex.MatchString(str) {
			return suffix
		}
	}
	return ""
}

// CoinPermutations calculates all possible permutations of coin values totalling
// given sum.
func CoinPermutations(target int, coins ...int) []int {
	perms := make([]int, target+1)
	perms[0] = 1
	for _, c := range coins {
		for i := c; i <= target; i++ {
			perms[i] += perms[i-c]
		}
	}
	return perms
}
