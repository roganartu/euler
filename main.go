package main

import (
	"fmt"
	"github.com/roganartu/euler/problems"
)

func main() {
	var i uint

	for {
		// Prompt user for problem to run
		fmt.Print("Which problem number would you like to solve?: ")
		_, _ = fmt.Scanf("%d", &i)

		// TODO refactor this (possibly using reflection?)
		switch i {
		case 3:
			problems.Problem_3()
		}
		fmt.Print("\n")
	}
}
