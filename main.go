package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/roganartu/euler/problems"
)

func main() {
	var i, arg uint

	flag.Parse()
	if len(flag.Args()) > 0 {
		x, _ := strconv.Atoi(flag.Arg(0))
		arg = uint(x)
	}

	for {
		if arg == 0 {
			// Prompt user for problem to run
			fmt.Print("Which problem number would you like to solve?: ")
			_, _ = fmt.Scanf("%d", &i)
			fmt.Print("\n")
		} else {
			i = arg
		}

		// TODO refactor this (possibly using reflection?)
		switch i {
		case 3:
			problems.Problem_3(uint64(0))
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
		case 13:
			problems.Problem_13()
		case 14:
			problems.Problem_14()
		case 15:
			problems.Problem_15()
		case 16:
			problems.Problem_16()
		case 17:
			problems.Problem_17()
		case 18:
			problems.Problem_18()
		}

		if arg > 0 {
			break
		}

		fmt.Print("\n")
	}
}
