package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func scanTextFromConsole() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text(), nil
	}

	return "", errors.New("Cannot read the text.")
}

func readFile(file *os.File) []Person {
	slice := make([]Person, 0)
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		splitted := strings.Split(fileScanner.Text(), " ")
		slice = append(slice, Person{fname: splitted[0], lname: splitted[1]})
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	return slice
}

func main() {
	fmt.Printf("Please enter path to the file: ")
	path, err := scanTextFromConsole()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	slice := readFile(file)

	for index, element := range slice {
		fmt.Printf("%d. %s %s.\n", index+1, element.fname, element.lname)
	}
}
