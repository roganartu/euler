package helpers

import (
	"bufio"
	"fmt"
	"github.com/bmizerany/assert"
	"os"
	"strconv"
	"testing"
)

// Test the first 10,000 primes
func TestIsPrime(t *testing.T) {
	var x int
	var i uint64
	var result bool
	var invalid uint64 = 0

	file, _ := os.Open("files/primelist")
	scanner := bufio.NewScanner(file)
	cache := make(map[uint64][]uint64)
	for scanner.Scan() {
		x, _ = strconv.Atoi(scanner.Text())
		i = uint64(x)

		// Non-prime numbers should return more than two factors
		for ; invalid < i; invalid++ {
			_, result = IsPrime(invalid, cache)
			assert.Equal(t, false, result, invalid, " should not return true")
		}

		// Primes should be the same as in the provided file
		_, result = IsPrime(i, cache)
		assert.Equal(t, true, result)
		invalid++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
