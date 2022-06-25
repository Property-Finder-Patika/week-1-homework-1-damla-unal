package main

import "fmt"

// package-level declaration
// Package-level variables are initialized before main begins
// The lifetime of a package-level variable is the entire execution of the program.
// this variable is visible not only throughout the source file that contains  its declaration, but throughout all the files of the package
const boilingF = 212.0

func main() {
	// local variables have dynamic lifetimes
	// local declarations are visible only within the function in which they are declared
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gF or %gC\n", f, c)

	// each cal l to new returns a distinct variable with a unique address
	p := newInt()
	q := newInt2()
	fmt.Println(p == q) // false

	// two variables whose type carries no information and is therefore of size zero,
	// such as struct{} or [0]int, may, depending on the implementation, have the same address.
	var a [0]int
	var b [0]int
	fmt.Println(a == b)
}

// newInt and newInt2 funcs have identical behaviors
func newInt() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	return &dummy

}
