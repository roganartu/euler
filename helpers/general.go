package helpers

import (
	"math"
)

// Factorise calculates all factors of the given target
func Factorise(target uint64) []uint64 {
	var i uint64
	factors := make([]uint64, 0)
	for i = 2; i <= uint64(math.Sqrt(float64(target)))+uint64(1); i++ {
		if target%i == 0 && target/i != 1 {
			factors = AppendUnique(factors, target/i, i)
		}
	}
	return factors
}