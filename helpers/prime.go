package helpers

import (
	"math"
)

func PrimeFactorise(target uint64) []uint64 {
	var i uint64

	// Calculate all factors for target
	factors := make([]uint64, 0)
	for i = 2; i < uint64(math.Sqrt(float64(target))); i++ {
		if target%i == 0 {
			factors = append(factors, target/i, i)
		}
	}

	// Detect primes and recurse if necessary
	if len(factors) == 0 {
		return []uint64{target}
	} else {
		f := make([]uint64, 0)
		for j := range factors {
			children := PrimeFactorise(factors[j])
			for k := range children {
				f = append(f, children[k])
			}
		}
		return f
	}
}
