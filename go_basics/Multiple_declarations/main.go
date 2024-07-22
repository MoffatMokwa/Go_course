package main

import "fmt"

func main() {
    // Declare and initialize car and cost variables
    car, cost := "Audi", 50000
    fmt.Println(car, cost) // Output: Audi 50000

    // Reassign car variable and declare a new year variable
    car, year := "BMW", 2018
    _ = year // Avoid unused variable warning

    // Initialize and reassign opened and file variables
    var opened = false
    opened, file := true, "a.txt"
    _, _ = opened, file // Avoid unused variable warnings

    // Declare multiple variables with their zero values
    var (
        salary    float64
        firstName string
        gender    bool
    )
    fmt.Println(salary, firstName, gender) // Output: 0  false

    // Declare multiple integer variables with zero values
    var a, b, c int
    fmt.Println(a, b, c) // Output: 0 0 0

    // Initialize integer variables i and j
    var i, j int
    i, j = 5, 8

    // Swap values of i and j
    j, i = i, j
    fmt.Println(i, j) // Output: 8 5

    // Perform and print the sum of an integer and a float value
    sum := 5 + 2.3
    fmt.Println(sum) // Output: 7.3
}
