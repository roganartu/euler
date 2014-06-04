package helpers

func PrimeFactorise(target uint64, cache map[uint64][]uint64) (map[uint64][]uint64, []uint64) {
	if cache != nil && cache[target] != nil {
		return cache, cache[target]
	}

	factors := Factorise(target)

	// Detect primes and recurse if necessary
	if len(factors) == 0 {
		var result []uint64
		if target < 2 {
			// Special cases of 0 and 1
			result = []uint64{}
		} else {
			result = []uint64{target}
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
		if cache != nil {
			cache[target] = f
			return cache, cache[target]
		} else {
			return nil, f
		}
	}
}

func IsPrime(n uint64, cache map[uint64][]uint64) (map[uint64][]uint64, bool) {
	cache, result := PrimeFactorise(n, cache)
	return cache, len(result) == 1 && result[0] == n
}
