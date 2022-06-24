package main

import (
	"fmt"
)

func main() {
	name := createNameWithExclamation("Damla")
	greet(name)

	// when the program runs, it takes the arguments provided via os.Args
	// To run this code block, please edit run configuration to take program argument or run this command on terminal "go run greet.go Damla"
	/* nameFromArgs := os.Args[1]
	greet(nameFromArgs)
	*/

	// input can also be received from the user using Scanln
	var nameFromInput string
	fmt.Println("Please enter your name:")
	_, err := fmt.Scanln(&nameFromInput)
	if err != nil {
		return
	}
	greet(nameFromInput)

}

// greet takes a string parameter and with this name that takes as a parameter, prints with format specifier -> "Hello ..."
func greet(name string) {
	fmt.Printf("Hello %s \n", name)
}

// createNameWithExclamation returns the parameter it received, concatenated with another string, as a string.
func createNameWithExclamation(name string) string {
	nameWithExclamationMark := name + "!"
	return nameWithExclamationMark
}
