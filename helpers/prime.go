package helpers

import (
	"math"
)

func PrimeFactorise(target uint64, cache map[uint64][]uint64) (map[uint64][]uint64, []uint64) {
	var i uint64

	if cache != nil && cache[target] != nil {
		return cache, cache[target]
	}

	// Calculate all factors for target
	factors := make([]uint64, 0)
	for i = 2; i <= uint64(math.Sqrt(float64(target)))+uint64(1); i++ {
		if target%i == 0 && target/i != 1 {
			factors = AppendUnique(factors, target/i, i)
		}
	}

	// Detect primes and recurse if necessary
	if len(factors) == 0 {
		var result []uint64
		if target < 2 {
			// Special cases of 0 and 1
			result = []uint64{}
		} else {
			result = []uint64{1, target}
		}

		// Cache result
		if cache != nil {
			cache[target] = result
			return cache, cache[target]
		} else {
			return nil, result
		}
	} else {
		f := make([]uint64, 0)
		for j := range factors {
			var children []uint64
			cache, children = PrimeFactorise(factors[j], cache)
			for k := range children {
				f = AppendUnique(f, children[k])
			}
		}
		f = AppendUnique(f, target)
		if cache != nil {
			cache[target] = f
			return cache, cache[target]
		} else {
			return nil, f
		}
	}
}
