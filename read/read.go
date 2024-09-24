/*
Write a program which reads information from a file and represents it in a slice of structs.

Assume that there is a text file which contains a series of names.

Each line of the text file has a first name and a last name, in that order,
separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name,
and lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file.

Your program will successively read each line of the text file and create a struct which contains
the first and last names found in the file.

Each struct created will be added to a slice,
and after all lines have been read from the file, your program will have a slice containing
one struct for each line in the file.

After reading all lines from the file, your program should
iterate through your slice of structs and print the first and last names found in each struct.

3 points will be given if a program is written.

2 points will be given if the program compiles correctly.

5 points will be given if test execution is successful and your program:

1. Opens a named text file.

2. Prints all first name/ last name pairs.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Your program will define a name struct which has two fields
//fname for the first name,and lname for the last name.

type Name struct {
	fname [20]string
	lname [20]string
}

func main() {

	var scannedTextFileName string

	//Your program should prompt the user for the name of the text file.
	fmt.Println("Enter the name of the text file:")

	_, err := fmt.Scan(&scannedTextFileName)

	if err != nil {
		fmt.Println("Invalid text file name. Please enter a valid string.")
	}

	//Your program will successively read each line of the text file
	//and create a struct which contains the first and last names found in the file.

	lines, err1 := readFile(scannedTextFileName)

	if err1 != nil {
		fmt.Println("Error reading the file")
	}

	//Each struct created will be added to a slice, and after all lines have been read
	//from the file, your program will have a slice containing one struct for each line in the file.
	sliceOfNames := make([]Name, 0)

	for i := range len(lines) {
		var nameStruct Name
		parts := strings.Split(lines[i], " ")

		nameStruct.fname[0] = parts[0]
		nameStruct.lname[0] = parts[1]

		sliceOfNames = append(sliceOfNames, nameStruct)
	}

	printNames(sliceOfNames)
}

func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var names []string

	//Each line of the text file has a first name and a last name, in that order,
	//separated by a single space on the line.
	for scanner.Scan() {
		line := scanner.Text()
		names = append(names, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning file: %w", err)
	}

	return names, nil
}

func printNames(sliceOfNames []Name) {

	//After reading all lines from the file, your program should
	//iterate through your slice of structs and print the first and last names found in each struct.

	for _, name := range sliceOfNames {
		fmt.Println("First name: ", name.fname[0], ". Last name: ", name.lname[0])
	}
}
