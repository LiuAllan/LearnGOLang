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
- https://pkg.go.dev/fmt@go1.17.2
