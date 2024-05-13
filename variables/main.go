package main

import "fmt"

func main() {

	// Constants can be character, string, boolean, or numeric values.
	// Constants cannot be declared using the := syntax.
	// https://gobyexample.com/constants
	const (
		Sunday    = iota // 0
		Monday           // 1
		Tuesday          // 2
		Wednesday        // 3
		Thursday         // 4
		Friday           // 5
		Saturday         // 6
	)

	const testconst = 1558

	// Declare a variable and initialize it (type strictly inferred)
	var name string
	var name2 string = "John Doe"

	// Assign a value to the variable using the assignment operator (shorthand notation)
	message := "Hello, welcome Mr./Ms. "

	// The %q verb shows them as quoted characters
	fmt.Printf("Name: %q\n", name)
	fmt.Printf("Type of name: %T\n", name)

	// %v show keys and values in their default formats.
	fmt.Printf("%v%v\n", message, name2)

	// Show type | From Golang example course
	i := 5
	f := 3.7
	b := true
	r := '界'

	fmt.Printf("int: %v\n", i)
	fmt.Printf("float64: %v\n", f)
	fmt.Printf("bool: %v\n", b)
	fmt.Printf("rune: %v\n", r)

	fmt.Println("Sunday:", Sunday)
	fmt.Println("Monday:", Monday)
	fmt.Println("Tuesday:", Tuesday)
	fmt.Println("Wednesday:", Wednesday)
	fmt.Println("Thursday:", Thursday)
	fmt.Println("Friday:", Friday)
	fmt.Println("Saturday:", Saturday)

	fmt.Printf("testconst: %v\n", testconst)
	fmt.Printf("Type of testconst: %T\n", testconst)
}
