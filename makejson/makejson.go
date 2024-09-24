/*
Write a program which prompts the user to first enter a name, and then enter
an address. Your program should create a map and add the name and address to
the map using the keys “name” and “address”, respectively. Your program should use
Marshal() to create a JSON object from the map, and then your program should print
the JSON object.

Submit your source code for the program,
“makejson.go”.

3 points will be given if a program is written.

2 points will be given if the program compiles correctly.

5 points will be given if the program correctly prints a JSON object with keys ("name", "address")
and they are associated with the name and address that was entered.
*/

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//Write a program which prompts the user to first enter a name, and then enter an address.
	var scanedName, scanedAddress string

	fmt.Println("Please enter the name: ")
	_, err := fmt.Scan(&scanedName)

	if err != nil {
		fmt.Println("Invalid name. Please enter a valid string.")
	}

	fmt.Println("\nPlease enter the address: ")

	_, err2 := fmt.Scan(&scanedAddress)

	if err2 != nil {
		fmt.Println("Invalid address. Please enter a valid string.")
	}

	//fmt.Println(scanedName)
	//fmt.Println(scanedAddress)

	//Your program should create a map
	// and add the name and address to the map using the keys “name” and “address”, respectively
	userMap := map[string]string{"name": scanedName, "address": scanedAddress}

	//fmt.Println(userMap)

	//Your program should use Marshal() to create a JSON object from the map
	barrJsonMarshall, err3 := json.Marshal(userMap)

	if err3 != nil {
		fmt.Println("Error in the json marshal creation")
	}

	//and then your program should print the JSON object
	fmt.Println("\nThe Json object is: ")
	fmt.Println(string(barrJsonMarshall))
}
