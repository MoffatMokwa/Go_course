package main

import "fmt"

// The main function is the entry point of the program.
func main() {
	// Declare an integer variable 'a' and assign it the value 4.
	var a = 4

	// Declare a floating-point variable 'b' and assign it the value 5.2.
	var b = 5.2

	// Convert the float value 'b' to an integer and assign it to 'a'.
	// Note: This conversion truncates the decimal part, so 'a' becomes 5.
	a = int(b) 

	// Print the values of 'a' and 'b'.
	// Output will be: "5 5.2"
	fmt.Println(a, b)

	// Declare variables of different types.
	var value int      // 'value' is an integer variable.
	var price float64  // 'price' is a floating-point variable.
	var name string    // 'name' is a string variable.
	var done bool      // 'done' is a boolean variable.

	// Print the default values of the declared variables.
	// Default values: 0 for int, 0.0 for float64, empty string for string, and false for bool.
	// Output will be: "0 0  false"
	fmt.Println(value, price, name, done)
}
