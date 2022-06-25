package main

import (
	"fmt"
)

// this class shows basically how OOP is achieved in Go

// Person is a struct type, struct is a type for data abstraction in Go
type Person struct {
	name string
}

// greet is method that belong to Person struct,
func (p Person) greet() string {
	return "Hello " + p.name + "!"
}

func main() {
	greeter := Person{"Damla"}
	// after creating instance of Person, greet method can be called.
	var greeting = greeter.greet()
	fmt.Printf("%s\n", greeting)
}
