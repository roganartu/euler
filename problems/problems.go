package problems

import (
	"fmt"
	"math"

	"../helpers"
)

func Problem_3() {
	var target, max uint64

	// Get user input
	fmt.Print("Enter number: ")
	_, _ = fmt.Scanf("%d", &target)
	if target == 0 {
		target = 600851475143
	}

	_, factors := helpers.PrimeFactorise(target, nil)
	fmt.Printf("Prime factors of %d are %d\n", target, factors)
	max = 0
	for i := range factors {
		if factors[i] > max {
			max = factors[i]
		}
	}
	fmt.Printf("Largest prime factor of %d is %d\n", target, max)
}

func Problem_5() {
	var i uint64 = 20

	fmt.Println("Finding smallest number divisible by all numbers 1 through 20 (inclusive)")
	for {
		failed := false
		for j := 1; j <= 20; j++ {
			if i%uint64(j) != 0 {
				failed = true
				break
			}
		}
		if !failed {
			fmt.Printf("%d is divisible by 1 through 20 inclusive\n", i)
			break
		}
		i++
	}
}

func Problem_6() {
	fmt.Println("Finding difference between sum-of-squares and square-of-sum for range 1-100")
	fmt.Printf("Difference is %d\n", uint64(math.Pow(float64(helpers.SumRange(1, 100)), 2))-
		helpers.SumRangeSquares(1, 100))
}

func Problem_7() {
	var i uint64 = 0
	var count int = 0
	cache := make(map[uint64][]uint64)

	fmt.Println("Calculating the 10,001st prime number")
	for {
		var result bool
		if cache, result = helpers.IsPrime(i, cache); result {
			count++
			if count == 10001 {
				fmt.Printf("The 10,001st prime number is %d\n", i)
				return
			}
		}
		i++
	}
}
