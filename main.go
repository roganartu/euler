package main

import (
	"fmt"

	"./problems"
)

func main() {
	var i uint

	for {
		// Prompt user for problem to run
		fmt.Print("Which problem number would you like to solve?: ")
		_, _ = fmt.Scanf("%d", &i)
		fmt.Print("\n")

		// TODO refactor this (possibly using reflection?)
		switch i {
		case 3:
			problems.Problem_3()
		case 5:
			problems.Problem_5()
		case 6:
			problems.Problem_6()
		case 7:
			problems.Problem_7()
		case 8:
			problems.Problem_8()
		case 9:
			problems.Problem_9()
		case 10:
			problems.Problem_10()
		case 11:
			problems.Problem_11()
		case 12:
			problems.Problem_12()
		}
		fmt.Print("\n")
	}
}
