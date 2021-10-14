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

	// Print
	fmt.Print("Print, ")  // Does not add a new line
	fmt.Print("world \n") // Escape character for new line
	fmt.Println("my age is", ageOne, "and the score is", scoreOne)

	// Formatting string (Printf)
	fmt.Printf("my age is %v and the score is %v \n", ageOne, scoreOne)  // Order of variables matters
	fmt.Printf("my name is %q and the score is %q \n", nameOne, nameTwo) // Order of variables matters
	fmt.Printf("age is of type %T \n", ageOne)                           // Outputs the type of the variable
	fmt.Printf("you scored %0.1f \n", 22.55)                             // Order of variables matters

	// Save Printf
	var str = fmt.Sprintf("my age is %v and the score is %v \n", ageOne, scoreOne)
	fmt.Println("The save string is: ", str)
}
