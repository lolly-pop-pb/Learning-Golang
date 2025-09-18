package main

import (
	"errors"
	"fmt"
)

func main() {
	var textinput string = "Print this too"
	printMe(textinput)

	var numerator int = 11
	var denominator int = 11
	// var result int = intDivision(numerator, denominator)
	// // fmt.Println(result)
	var result, remainder, err = intDivisionwithError(numerator, denominator)
	// // using if else if and els
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("The result of the integer division is %v", result)
	// } else {
	// 	fmt.Printf("The result of the integer division is %v and the remainder is %v", result, remainder)
	// }

	var q, r int = intDivision_rem(numerator, denominator)
	fmt.Printf("The result of the integer division is %v and the remainder is %v \n", q, r)
	// using switch
	// break is implied so from a logical perspective this is equal to the if else block
	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("The result of integer division is %v", result)
	default:
		fmt.Printf("The result of integer division is %v and remainder is %v", result, remainder)
	}

	// conditional switch statements
	// to print additional values based on remainder value
	switch remainder {
	case 0:
		fmt.Printf("The division was exact")
	case 1, 2:
		fmt.Printf("The division was close")
	default:
		fmt.Printf("The division was not close")
	}

	var x int = 1
	var y int = 2
	var z int = 3
	var w int = 4
	if x == y && z == w {
		fmt.Println("\nPassed the check")
	} else {
		fmt.Println("\nFailed")
	}

	if x == y || z == w {
		fmt.Println("Passed the check")
	} else {
		fmt.Println("Failed")
	}
}

func printMe(x string) {
	fmt.Println(x)
}

// func intDivision(a int, b int) int{
// 	var res int = a/b
// 	return res
// }

func intDivisionwithError(a int, b int) (int, int, error) {
	var err error
	if b == 0 {
		err = errors.New("Cannot divide by 0")
		return 0, 0, err
	}
	var quotient int = a / b
	var remainder int = a % b
	return quotient, remainder, err
}

func intDivision_rem(a int, b int) (int, int) {
	var quotient int = a / b
	var remainder int = a % b
	return quotient, remainder
}
