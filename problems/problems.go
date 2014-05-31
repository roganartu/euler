package problems

import (
	"fmt"
	"github.com/roganartu/euler/helpers"
)

func Problem_3() {
	var target, max uint64

	// Get user input
	fmt.Print("Enter number: ")
	_, _ = fmt.Scanf("%d", &target)
	if target == 0 {
		target = 600851475143
	}

	factors := helpers.PrimeFactorise(target)
	fmt.Printf("Prime factors of %d are %d\n", target, factors)
	max = 0
	for i := range factors {
		if factors[i] > max {
			max = factors[i]
		}
	}
	fmt.Printf("Largest prime factor of %d is %d\n", target, max)
}
