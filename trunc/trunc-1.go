/*
Write a program which prompts the user to enter a floating point number and prints the integer
which is a truncated version of the floating point number that was entered.
Truncation is the process of removing the digits to the right of the decimal place.

Submit your source code for the program, “trunc.go”.
*/

package main

import "fmt"

func main1() {
	var floatingPointNumber float64

	fmt.Printf("Enter the floating point number: ")
	//Write a program which prompts the user to enter a floating point number
	fmt.Scan(&floatingPointNumber)

	//truncated version of the floating point number that was entered.
	var truncatedNumber = int(floatingPointNumber)

	//prints the integer
	fmt.Println("The truncated number is: ", truncatedNumber)
}
