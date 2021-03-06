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
		case 19:
			problems.Problem_19()
		case 20:
			problems.Problem_20()
		case 21:
			problems.Problem_21()
		case 22:
			problems.Problem_22()
		case 23:
			problems.Problem_23()
		case 24:
			problems.Problem_24()
		case 25:
			problems.Problem_25()
		case 26:
			problems.Problem_26()
		case 27:
			problems.Problem_27()
		case 28:
			problems.Problem_28()
		case 29:
			problems.Problem_29()
		case 30:
			problems.Problem_30()
		case 31:
			problems.Problem_31()
		case 32:
			problems.Problem_32()
		case 33:
			problems.Problem_33()
		case 34:
			problems.Problem_34()
		case 35:
			problems.Problem_35()
		case 36:
			problems.Problem_36()
		case 37:
			problems.Problem_37()
		case 67:
			problems.Problem_67()
		}

		if arg > 0 {
			break
		}

		fmt.Print("\n")
	}
}
