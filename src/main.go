package main // For executable files

import (
	"fmt" // Standard library for formatting strings and printing
	"sort"
	"strings"
)

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

	// Arrays
	// var ages [3]int = [3]int{20, 25, 30} // ages will be an array of length 3 of all ints
	var ages = [3]int{20, 25, 30}
	names := [4]string{"yoshi", "bowser", "mario", "wario"}
	fmt.Println(ages, names, len(ages)) // len() grabs the length of the array

	// Slices (use arrays under the hood)
	scores := []int{100, 50, 60}
	scores[2] = 25              // Set index 2 to be value 25
	scores = append(scores, 85) // append returns a new slice
	fmt.Println(scores, len(scores))

	// Slice ranges
	rangeOne := names[1:3]  // Includes from 1 to but not including 3
	rangeTwo := names[2:]   // From index 2 to the end, including end
	rangeThree := names[:3] // From start to index 3 but not including 3
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	// Packages
	// Strings pack
	greeting := "hello there world"
	fmt.Println(strings.Contains(greeting, "hello"))         // returns true or false if strings contain value
	fmt.Println(strings.ReplaceAll(greeting, "hello", "Hi")) // Does not mutate original variable
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll")) // Get index of where the value starts
	fmt.Println(strings.Split(greeting, " "))  // Split string into a splice by delimitter

	// Sort pack
	newAges := []int{45, 20, 35, 30, 75, 60, 50, 25}
	sort.Ints(newAges) // Alters the original variable
	fmt.Println(newAges)
	fmt.Println(sort.SearchInts(newAges, 30)) // Searches slice and returns index

	newStrings := []string{"one", "two", "five", "three"}
	sort.Strings(newStrings)
	fmt.Println(newStrings)
	fmt.Println(sort.SearchStrings(newStrings, "five"))

	// Loops
	x := 0
	for x < 5 {
		fmt.Println("value of x is: ", x)
		x++
	}

	for i := 0; i < len(newStrings); i++ {
		fmt.Println(newStrings[i])
	}

	// Loop through slice and get each index and value
	// If we don't want to use the index or value, replace it with underscore "_"
	for index, value := range newStrings {
		fmt.Printf("the position at index %v is %v \n", index, value)
	}

	// Booleans and Conditions
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else if x == 0 {
		fmt.Println("x is 0")
	} else {
		fmt.Println("x is negative")
	}

	sayGreeting("mario")
	sayBye("luigi")

	cycle([]string{"cloud", "tifa"}, sayGreeting) // passing the function name
	result := cycle([]string{"cloud", "tifa"}, sayBye)
	fmt.Println(result)

	firstname, secondname := getInitials("chris bumstead")
	fmt.Println(firstname, secondname)
}

// function cycle takes a function as the 2nd parameter
func cycle(n []string, f func(string)) string {
	for _, v := range n {
		f(v)
	}
	return "returns a string"
}

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

// Function to return multiple values
func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1]) // Gets first letter of each string
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}
