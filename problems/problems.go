package problems

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"

	"../helpers"
)

func Problem_3() {
	var target, max uint64

	// Get user input
	fmt.Print("Enter number [600851475143]: ")
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

func Problem_11() {
	grid := make([][]int, 0)
	grid = append(grid, []int{8, 2, 22, 97, 38, 15, 0, 40, 0, 75, 4, 5, 7, 78, 52, 12, 50, 77, 91, 8})
	grid = append(grid, []int{49, 49, 99, 40, 17, 81, 18, 57, 60, 87, 17, 40, 98, 43, 69, 48, 4, 56, 62, 0})
	grid = append(grid, []int{81, 49, 31, 73, 55, 79, 14, 29, 93, 71, 40, 67, 53, 88, 30, 3, 49, 13, 36, 65})
	grid = append(grid, []int{52, 70, 95, 23, 4, 60, 11, 42, 69, 24, 68, 56, 1, 32, 56, 71, 37, 2, 36, 91})
	grid = append(grid, []int{22, 31, 16, 71, 51, 67, 63, 89, 41, 92, 36, 54, 22, 40, 40, 28, 66, 33, 13, 80})
	grid = append(grid, []int{24, 47, 32, 60, 99, 3, 45, 2, 44, 75, 33, 53, 78, 36, 84, 20, 35, 17, 12, 50})
	grid = append(grid, []int{32, 98, 81, 28, 64, 23, 67, 10, 26, 38, 40, 67, 59, 54, 70, 66, 18, 38, 64, 70})
	grid = append(grid, []int{67, 26, 20, 68, 2, 62, 12, 20, 95, 63, 94, 39, 63, 8, 40, 91, 66, 49, 94, 21})
	grid = append(grid, []int{24, 55, 58, 5, 66, 73, 99, 26, 97, 17, 78, 78, 96, 83, 14, 88, 34, 89, 63, 72})
	grid = append(grid, []int{21, 36, 23, 9, 75, 0, 76, 44, 20, 45, 35, 14, 0, 61, 33, 97, 34, 31, 33, 95})
	grid = append(grid, []int{78, 17, 53, 28, 22, 75, 31, 67, 15, 94, 3, 80, 4, 62, 16, 14, 9, 53, 56, 92})
	grid = append(grid, []int{16, 39, 5, 42, 96, 35, 31, 47, 55, 58, 88, 24, 0, 17, 54, 24, 36, 29, 85, 57})
	grid = append(grid, []int{86, 56, 0, 48, 35, 71, 89, 7, 5, 44, 44, 37, 44, 60, 21, 58, 51, 54, 17, 58})
	grid = append(grid, []int{19, 80, 81, 68, 5, 94, 47, 69, 28, 73, 92, 13, 86, 52, 17, 77, 4, 89, 55, 40})
	grid = append(grid, []int{04, 52, 8, 83, 97, 35, 99, 16, 7, 97, 57, 32, 16, 26, 26, 79, 33, 27, 98, 66})
	grid = append(grid, []int{88, 36, 68, 87, 57, 62, 20, 72, 3, 46, 33, 67, 46, 55, 12, 32, 63, 93, 53, 69})
	grid = append(grid, []int{04, 42, 16, 73, 38, 25, 39, 11, 24, 94, 72, 18, 8, 46, 29, 32, 40, 62, 76, 36})
	grid = append(grid, []int{20, 69, 36, 41, 72, 30, 23, 88, 34, 62, 99, 69, 82, 67, 59, 85, 74, 4, 36, 16})
	grid = append(grid, []int{20, 73, 35, 29, 78, 31, 90, 1, 74, 31, 49, 71, 48, 86, 81, 16, 23, 57, 5, 54})
	grid = append(grid, []int{01, 70, 54, 71, 83, 51, 54, 69, 16, 92, 33, 48, 61, 43, 52, 1, 89, 19, 67, 48})

	var max int = 0
	var combo []int
	// Horizontal combination products
	for _, combo = range helpers.HorizontalGridCombos(grid, 4) {
		max = int(math.Max(float64(helpers.ProductInt(combo)), float64(max)))
	}
	// Vertical combination products
	for _, combo = range helpers.VerticalGridCombos(grid, 4) {
		max = int(math.Max(float64(helpers.ProductInt(combo)), float64(max)))
	}
	// Diagonal combination products
	for _, combo = range helpers.DiagonalGridCombos(grid, 4) {
		max = int(math.Max(float64(helpers.ProductInt(combo)), float64(max)))
	}

	fmt.Printf("Maximum number in grid is %d\n", max)
}

func Problem_12() {
	fmt.Println("Finding first triangle number with > 500 factors")
	var i, current uint64 = 0, 0
	var factors []uint64

	for i = 1; ; i++ {
		current += i
		factors = helpers.Factorise(current)
		if len(factors) > 500 {
			fmt.Printf("Found %d with factors %d\n", current, factors)
			break
		}
	}
}

func Problem_13() {
	total := big.NewInt(int64(0))
	each := big.NewInt(int64(0))

	file, _ := os.Open("problems/files/problem13")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		each.SetString(scanner.Text(), 10)
		total.Add(total, each)
	}
	fmt.Printf("The sum of the provided 100x 50 digit numbers is %s\n", total.String())
}

func Problem_14() {
	var result int
	cache := make(map[int]int)
	max := 0
	maxStart := 0
	for i := 1; i < 1000000; i++ {
		cache, result = helpers.Collatz(i, cache)
		if result > max {
			max = result
			maxStart = i
		}
	}
	fmt.Printf("The longest Collatz chain for a starting number under one million "+
		"is %d started from %d\n", max, maxStart)
}

func Problem_15() {
	var target int
	// Get user input
	fmt.Print("Enter grid size [20]: ")
	_, _ = fmt.Scanf("%d", &target)
	if target == 0 {
		target = 20
	}

	fmt.Printf("There are %d possible paths through a %dx%d lattice\n",
		helpers.LatticePathCount(target), target, target)
}
