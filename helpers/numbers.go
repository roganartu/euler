package helpers

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
