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
