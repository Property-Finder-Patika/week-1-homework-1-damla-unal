package main

import "fmt"

// Celsius and Fahrenheit type declaration defines a new named type that has the same underlying type as an existing type.
type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	FreezingC Celsius = 0
	BoilingC  Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
	// fmt.Printf("%g\n", boilingF-FreezingC) // compile error: type mismatch
	fmt.Println(FToC(boilingF))

	boilingK := CToK(BoilingC)
	fmt.Println(boilingK)

	diff := KToC(Kelvin(1)) - Celsius(2)
	fmt.Println(diff)
}

// CToF converts between the two scales C -> F
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// KToC coverts a Kelvin to Celsius
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

// CToK coverts a Celsius to Kelvin
func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}
