package main

import "fmt"

// Method 1
func string_concat(s1 string, s2 string) string {
	return s1 + s2
}

// Method 2
func string_concat2(s1, s2 string) string {
	return s1 + s2
}

// Function as parameter | Nested function
type myFunc func(string, string) string

func string_concat3(fn myFunc) string {
	return fn("", "")
}

func main() {
	var name string = "John Doe"
	var message string = "Hello, welcome Mr./Ms. "

	fmt.Println(string_concat(message, name))
	fmt.Println(string_concat2(message, name))

	fmt.Println(string_concat3(string_concat))

	test_func := func(s1, s2 string) string {
		return "s1, s2 is a parameters as string for test_func which return as string and string_concat3 will return as string. Also the message change to this."
	}
	result := string_concat3(test_func)
	fmt.Println(result)
}
