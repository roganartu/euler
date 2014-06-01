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

func Problem_8() {
	var number string = "73167176531330624919225119674426574742355349194934" +
		"96983520312774506326239578318016984801869478851843" +
		"85861560789112949495459501737958331952853208805511" +
		"12540698747158523863050715693290963295227443043557" +
		"66896648950445244523161731856403098711121722383113" +
		"62229893423380308135336276614282806444486645238749" +
		"30358907296290491560440772390713810515859307960866" +
		"70172427121883998797908792274921901699720888093776" +
		"65727333001053367881220235421809751254540594752243" +
		"52584907711670556013604839586446706324415722155397" +
		"53697817977846174064955149290862569321978468622482" +
		"83972241375657056057490261407972968652414535100474" +
		"82166370484403199890008895243450658541227588666881" +
		"16427171479924442928230863465674813919123162824586" +
		"17866458359124566529476545682848912883142607690042" +
		"24219022671055626321111109370544217506941658960408" +
		"07198403850962455444362981230987879927244284909188" +
		"84580156166097919133875499200524063689912560717606" +
		"05886116467109405077541002256983155200055935729725" +
		"71636269561882670428252483600823257530420752963450"

	digits, product := helpers.GreatestAdjacentProduct(number, 13)
	fmt.Printf("The largest product of 13 adjacent digits is %d composed of %d\n",
		product, digits)
}

func Problem_9() {
	fmt.Println("Finding largest Pythagorean Triplet that satisfies a + b + c = 1000")

	triplets := helpers.NaturalNSumToY(3, 1000)

	for _, t := range triplets {
		if helpers.PythagoreanTriplet(t[0], t[1], t[2]) {
			product := t[0] * t[1] * t[2]
			fmt.Printf("Triplet %d with product %d\n", t, product)
			return
		}
	}

	fmt.Println("No Pythagorean Triplet satisfies a + b + c = 1000")
}

func Problem_10() {
	fmt.Print("Finding sum of all primes < 2,000,000")

	var i uint64
	var result bool
	var sum uint64 = 0
	cache := make(map[uint64][]uint64)
	for i = 0; i < 2000000; i++ {
		if cache, result = helpers.IsPrime(i, cache); result {
			sum += i
		}

		// Display progress dots
		if i%250000 == 0 {
			fmt.Print(".")
		}
	}
	fmt.Printf("Sum is %d", sum)
}
