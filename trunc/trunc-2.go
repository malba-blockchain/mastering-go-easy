/*
Write a program which prompts the user to enter a floating point number and prints the integer
which is a truncated version of the floating point number that was entered.
Truncation is the process of removing the digits to the right of the decimal place.

Submit your source code for the program, “trunc.go”.
*/

package main

import (
	"fmt"
	"math"
)

func main2() {
	var floatingPointNumber float64

	fmt.Printf("Enter the floating point number: ")
	for {
		//Write a program which prompts the user to enter a floating point number
		_, err := fmt.Scan(&floatingPointNumber)

		if err != nil {
			fmt.Println("Invalid input. Please enter a valid floating point number.")
			continue
		}

		// Check if the number is within the range of int64
		if floatingPointNumber > math.MaxInt64 || floatingPointNumber < math.MinInt64 {
			fmt.Println("Number is too large or too small for int64 truncation.")
			continue
		}

		break
	}

	//truncated version of the floating point number that was entered.
	var truncatedNumber = int(floatingPointNumber)

	//prints the integer
	fmt.Println("The truncated number is: ", truncatedNumber)
}
