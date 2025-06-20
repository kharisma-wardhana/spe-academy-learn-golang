package main

import "fmt"

func main() {
	// Declare a constant
	const pi float64 = 3.14159 // Declare a constant named 'pi' of type float64
	fmt.Println(pi)            // Print the value of 'pi' to the console

	// Declare a constant using iota
	const (
		first  = iota // iota starts at 0
		second        // iota increments by 1 for each constant in the block
		third         // iota is now 2
	)
	fmt.Println(first, second, third) // Print the values of first, second, and third, which will be 0, 1, and 2 respectively

	// Declare a constant with a string value
	const greeting = "Hello, Go!" // Declare a constant named 'greeting' with a string value
	fmt.Println(greeting)         // Print the value of 'greeting' to the console

	// Declare multiple constants in a single line
	const (
		apple  = "Apple"  // Declare a constant named 'apple' with a string value
		banana = "Banana" // Declare a constant named 'banana' with a string value
		orange = "Orange" // Declare a constant named 'orange' with a string value
	)
	fmt.Println(apple, banana, orange) // Print the values of apple, banana, and orange
	// which will be "Apple Banana Orange"

	// Declare constants using a custom type
	type Status int // Define a custom type named 'Status' of type int
	const (
		Active    Status = iota // Declare a constant named 'Active' of type Status with value 0
		Inactive                // Declare a constant named 'Inactive' of type Status with value 1
		Suspended               // Declare a constant named 'Suspended' of type Status with value 2
	)
	fmt.Println(Active, Inactive, Suspended) // Print the values of Active, Inactive, and Suspended, which will be 0, 1, and 2 respectively

	// Declare constant using a bitwise operation
	const (
		Read    = 1 << iota // 1 << 0 = 1
		Write               // 1 << 1 = 2
		Execute             // 1 << 2 = 4
	)
	fmt.Println(Read, Write, Execute) // Print the values of Read, Write, and Execute
	// which will be 1, 2, and 4 respectively

	// Declare a constant using a rune
	const char rune = 'A' // Declare a constant named 'char' of type rune with the value 'A'
	fmt.Println(char)     // Print the value of 'char' to the console, which will be 'A'

}
