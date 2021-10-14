package main // For executable files

import "fmt" // Standard library for formatting strings and printing

func main() { // Entry point of the application
	fmt.Println("Hello world")

	// Strings
	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string
	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "bowser"
	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "yoshi"
	fmt.Println(nameFour)

	// Ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40
	fmt.Println(ageOne, ageTwo, ageThree)

	// Bits and memory
	var numOne int8 = 25 // 8 bits
	var numTwo int8 = -128
	var numThree uint8 = 25
	fmt.Println(numOne, numTwo, numThree)

	// Floats
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 3845934578734859.234234234
	scoreThree := 1.5 // Shorthand inferred as float64
	fmt.Println(scoreOne, scoreTwo, scoreThree)
}
