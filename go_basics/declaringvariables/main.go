package main

import "fmt"

func main() {
// Initialize age variable with value 30 using var keyword
var age int = 30
fmt.Println("Age:", age)

// Initialize name variable with value "Dan" using var keyword
var name = "Dan"
fmt.Println("Your name is:", name)

// Initialize a string variable s using shortend declaration with value "learning golang" and print it
s := "learning golang"
fmt.Println(s)

// Reassign the name variable to "Moffat"
name = "Moffat"

// Initialize a new string variable name1 shortend declaration with value "John"
name1 := "John"

// Suppress unused variable warning for name1
_ = name1

}