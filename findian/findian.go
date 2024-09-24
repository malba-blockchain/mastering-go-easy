package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Create a new reader for capturing the entire input including spaces
	reader := bufio.NewReader(os.Stdin)

	var stringToFind string

	fmt.Println("Enter the string:")

	// Read the full string including spaces
	stringToFind, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input, please enter a valid string")
		os.Exit(0)
	}

	// Remove the spaces of the string and convert to lowercase
	stringToFind = strings.TrimSpace(strings.ToLower(stringToFind))

	// Check if string starts with 'i', ends with 'n', and contains 'a'
	if strings.HasPrefix(stringToFind, "i") &&
		strings.HasSuffix(stringToFind, "n") &&
		strings.Contains(stringToFind, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
