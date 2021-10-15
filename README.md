# LearnGOLang
Repository for learning GO lang.

## What is Go?
- Go is a fast, statically typed compiled language
- Compiled to machine level code, does not need to be interpretted one line at a time
- Static typed - picks up variable types
- General purpose language
- Built-in testing support
- Object-oriented langauge in it's own way different from other languages

## First Go file
- One main() function allowed in the application
- Methods on imported packages all start with a capital letter ie) Println
- Running the program: 
> go run filename.go

## Variables, Strings, Numbers
- Strings in GO must use double quotes: ""
- Variables can be typed or automatically infer it:
> var nameOne string = "mario"  
> var nameTwo = "luigi"  
> var nameThree string // Variable is initialized. If no value is given, the string is empty  
> nameFour := "yoshi" // Shorthand that is only used to initialize the variable the first time  

- ":=" Can't be used **outside** of a function

- Ints can be typed to use bits. Prevents certain numbers to be used within the scope.
> var ageOne int = 20  
> var ageTwo = 30  
> ageThree := 40  
> var numOne int8 = 25 // 8 bits  
> var numThree uint = 25 // Prevents the use of negative ints (0 - 255)

- Floats must specify the bit-size. Bit-size dictate the range of numbers we can use
> var scoreOne float32 = 25.98  
> var scoreTwo float64 = 3845934578734859.234234234  
> scoreThree := 1.5 // Shorthand inferred as float64  

## Printing & Formatting Strings
- Printing variables and strings
> fmt.Println("my age is", ageOne, "and the score is", scoreOne)
- Printing formatted strings using a format specifier "%v". The order of variables matters.
> fmt.Printf("my age is %v and the score is %v", ageOne, scoreOne)
- Specifier "%q" adds quotes around the string variables
- Specifier "%T" outputs the type of the variable
- Specifier "%f" outputs floats. This can be rounded to the nearest decimal point
> fmt.Printf("you scored %0.1f", 22.55) = "you scored 22.6"
- Sprintf = Save printf allows ability to save the string into a variable
> var str = fmt.Sprintf("my age is %v and the score is %v \n", ageOne, scoreOne)  
> fmt.Println("The save string is: ", str)  
https://pkg.go.dev/fmt@go1.17.2

## Arrays and Slices
- If the length of an array is declared, the array is a fixed and can't be changed
- Arrays are created using curley braces instead of square brackets
> var ages [3]int = [3]int{20, 25, 30} // ages will be an array of length 3 of all ints  
> var ages = [3]int{20, 25, 30} // Same as above  
> names := [4]string{"yoshi", "bowser", "mario", "wario"}  
> fmt.Println(ages, len(ages)) // len() grabs the length of the array

- Slices use arrays under the hood but can manipulate arrays more
- When no number is placed inside the square brackets, it specifies to use slices instead of arrays. They have no fixed length
- Can append items to a slice
> scores := []int{100, 50, 60}  
> scores[2] = 25              // Set index 2 to be value 25  
> scores = append(scores, 85) // append returns a new slice  
> fmt.Println(scores, len(scores))  
- Slice ranges:
> rangeOne := names[1:3] // Includes from 1 to but not including 3  
> rangeTwo := names[2:]   // From index 2 to the end, including end  
> rangeThree := names[:3] // From start to index 3 but not including 3  

## The Standard Library
- Strings package
> greeting := "hello there world"  
> fmt.Println(strings.Contains(greeting, "hello"))         // returns true or false if > strings contain value  
> fmt.Println(strings.ReplaceAll(greeting, "hello", "Hi")) // Does not mutate original > variable  
> fmt.Println(strings.ToUpper(greeting))  
> fmt.Println(strings.Index(greeting, "ll")) // Get index of where the value starts  
> fmt.Println(strings.Split(greeting, " "))  // Split string into a splice by delimitter  
- Sort package
> newAges := []int{45, 20, 35, 30, 75, 60, 50, 25}  
> sort.Ints(newAges) // Alters the original variable  
> fmt.Println(newAges)  
> fmt.Println(sort.SearchInts(newAges, 30)) // Searches slice and returns index  

> newStrings := []string{"one", "two", "five", "three"}  
> sort.Strings(newStrings)  
> fmt.Println(newStrings)  
> fmt.Println(sort.SearchStrings(newStrings, "five"))  
https://pkg.go.dev/std

## Loops
> x := 0  
> for x < 5 {  
>     fmt.Println("value of x is: ", x)  
>     x++  
> }  
> 
> for i := 0; i < len(newStrings); i++ {  
>     fmt.Println(newStrings[i])  
> }  
> 
> // Loop through slice and get each index and value  
> // If we don't want to use the index or value, replace it with underscore "_"  
> for index, value := range newStrings {  
>     fmt.Printf("the position at index %v is %v \n", index, value)  
> }  
## Booleans and Conditions
- If and else statements work the same. They do not require brackets
    if x > 10 {
        fmt.Println("x is greater than 10")
    } else if x == 0 {
        fmt.Println("x is 0")
    } else {
        fmt.Println("x is negative")
    }

## Functions