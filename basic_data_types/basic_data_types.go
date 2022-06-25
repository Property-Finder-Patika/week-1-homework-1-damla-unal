package main

import "fmt"

func main() {

	// Float to integer conversion discards any fractional part, truncating toward zero.
	f := 3.141 // a float64
	i := int(f)
	fmt.Println(f, i) // "3.141 3"
	f = 1.99
	fmt.Println(int(f)) // "1"

	// Rune literals are written as a character wit hin single quotes.
	// they are printed with %c, or with %q
	ascii := 'a'
	unicode := 'D'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 D 'D'"
	fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"

	// The built-in function complex creates a complex number from its real and imaginary components,
	// and the built-in real and imag functions extract those components
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Println(x * y)               // "(-5+10i)"
	fmt.Println(real(x * y))         // "-5"
	fmt.Println(imag(x * y))         // "10"

	// The built-in len function returns the number of bytes (not runes) in a string ,
	//and the index operation s[i] retrieves the i-th byte of string s, where 0 ≤ i < len(s).
	s := "hello, world"
	fmt.Println(len(s))     // "12"
	fmt.Println(s[0], s[7]) // "104 119" ('h' and 'w')
	//c := s[len(s)] // panic: index out of range
	fmt.Println(s[:5])             // "hello"
	fmt.Println(s[7:])             // "world"
	fmt.Println(s[:])              // "hello, world"
	fmt.Println("goodbye" + s[5:]) // "goodbye, world"

	// Since strings are immutable, constructions that try to modify a string’s data in place are not allowed.
	// s[0] = 'L' // compile error: cannot assign to s[0]
}
