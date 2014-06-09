package helpers

import (
	"math"
	"math/big"
)

// Stringify converts n into a string representation of its value.
//
// The British usage of "and" is adhered to during generation.
// eg: 1234 would generate "one thousand two hundred and thirty four"
func Stringify(n int, cache map[int]string) (map[int]string, string) {
	if result, ok := cache[n]; ok {
		return cache, result
	}

	const (
		ONES     = iota
		TENS     = iota
		TEENS    = iota
		HUNDRED  = "hundred"
		THOUSAND = "thousand"
		MILLION  = "million"
	)

	var result, tmp string
	words := make([][]string, 3)
	words[ONES] = []string{"zero", "one", "two", "three", "four", "five",
		"six", "seven", "eight", "nine"}
	words[TENS] = []string{"ten", "twenty", "thirty", "forty", "fifty", "sixty",
		"seventy", "eighty", "ninety"}
	words[TEENS] = []string{"ten", "eleven", "twelve", "thirteen", "fourteen",
		"fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}

	if n < 10 {
		result = words[ONES][n]
	} else if n < 20 {
		result = words[TEENS][n%10]
	} else if n < 100 {
		cache, tmp = Stringify(n%10, cache)
		result = words[TENS][n/10-1]
	} else if n < 1000 {
		cache, tmp = Stringify(n%100, cache)
		result = words[ONES][n/100] + " hundred"
	} else if n < 10000 {
		cache, tmp = Stringify(n%1000, cache)
		result = words[ONES][n/1000] + " thousand"
	}

	if tmp != "" && tmp != "zero" {
		if n >= 100 && n < 1000 {
			result += " and"
		}
		result += " " + tmp
	}

	if cache != nil {
		cache[n] = result
	}
	return cache, result
}

// IsAmicable determines whether the given target is an Amicable Number.
//
// It returns the other number in the pair if found, otherwise 0.
func IsAmicable(target int) int {
	a := SumFactors(uint64(target)) + 1
	b := SumFactors(a) + 1
	if b == uint64(target) && b != a {
		return int(a)
	}
	return 0
}

// IsPerfectNum determines whether the given target is a Perfect Number.
//
// Specifically, whether the proper divisors of target sum to target.
//
// Returns 0 if target is a Perfect Number, -1 if deficient and +1 if abundant.
// Deficient means sum of target's factors is less than target
// Abundant means sum of target's factors is greater than target
func IsPerfectNum(target int) int {
	sum := int(SumFactors(uint64(target)) + 1)
	if sum < target {
		return -1
	} else if sum > target {
		return 1
	}
	return 0
}

// IsCyclic determines whether the given target is the denominator of a
// Cyclic Number.
//
// It first checks whether the target is a prime number, and if it is
// then determines the length of the period using the following formula:
// (10^(p-1)-1)/p
//
// Returns true if p is prime and the length of the period is equal to p-1
// otherwise false.
func IsCyclic(p int) bool {
	if _, prime := IsPrime(uint64(p), nil); prime {
		// Add 1 zero if greater than 9, 2 zeroes if greater than 99 etc
		multiplier := big.NewRat(int64(math.Pow10(int(math.Log10(float64(p))))), 1)

		a := big.NewInt(0)
		frac := big.NewRat(1, 1)
		a.Exp(big.NewInt(10), big.NewInt(int64(p-1)), nil)
		a.Sub(a, big.NewInt(1))
		frac.SetFrac(a, big.NewInt(int64(p)))
		frac.Mul(frac, multiplier)

		str := frac.FloatString(1)
		if str[len(str)-1] == '0' && ShortRepeatedString(str[:len(str)-2]) == "" {
			return true
		}
	}
	return false
}
