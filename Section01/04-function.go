package main

import "fmt"

func main() {
	sayHello() // Call the sayHello function to print a message

	// Declare a function named 'greet' that takes a string parameter and returns nothing
	greet := func(name string) {
		fmt.Println("Hello,", name) // Print a greeting message with the provided name
	}

	// Call the 'greet' function with the argument "Alice"
	greet("Alice") // Output: Hello, Alice

	// Declare a function that returns an integer value
	add := func(a int, b int) int {
		return a + b // Return the sum of the two integers
	}

	// Call the 'add' function and print the result
	result := add(5, 3) // Output: 8
	fmt.Println(result)

	// Declare a function that takes two parameters and returns a string
	concat := func(firstName string, lastName string) (result string) {
		result = firstName + " " + lastName // Concatenate the first and last names with a space in between
		return
	}
	// Call the 'concat' function and print the result
	fullName := concat("John", "Doe") // Output: John Doe
	fmt.Println(fullName)

	// Declare a function that takes a variable number of arguments
	sum := func(numbers ...int) int {
		total := 0                    // Initialize a variable to hold the sum
		for _, num := range numbers { // Iterate over the variable arguments
			total += num // Add each number to the total
		}
		return total // Return the total sum
	}
	// Call the 'sum' function with multiple arguments and print the result
	total := sum(1, 2, 3, 4, 5) // Output: 15
	fmt.Println(total)

	// Declare a function that returns multiple values
	multiplyAndDivide := func(a, b int) (int, float64) {
		multiplication := a * b             // Calculate the multiplication of a and b
		division := float64(a) / float64(b) // Calculate the division of a by b, converting to float64
		return multiplication, division     // Return both results
	}
	// Call the 'multiplyAndDivide' function and print the results
	mult, div := multiplyAndDivide(10, 2) // Output: 20, 5.0
	fmt.Println("Multiplication:", mult, "Division:", div)

	// Declare a function that takes a function as an argument
	applyFunction := func(fn func(int, int) int, a int, b int) int {
		return fn(a, b) // Call the provided function with the given arguments and return the result
	}
	// Call the 'applyFunction' with the 'add' function and print the result
	resultApply := applyFunction(add, 4, 6)                // Output: 10
	fmt.Println("Result from applyFunction:", resultApply) // Print the result from applyFunction

	// Declare a function that returns another function
	makeMultiplier := func(factor int) func(int) int {
		return func(x int) int {
			return x * factor // Return a function that multiplies its argument by the factor
		}
	}
	// Call the 'makeMultiplier' function to create a multiplier function
	multiplyByTwo := makeMultiplier(2) // Create a multiplier function that multiplies by 2

	// Call the returned function and print the result
	resultMultiplier := multiplyByTwo(5)                              // Output: 10
	fmt.Println("Result from multiplier function:", resultMultiplier) // Print the result from the multiplier function

	// Declare a function that uses a closure
	counter := func() func() int {
		count := 0 // Initialize a variable to hold the count
		return func() int {
			count++      // Increment the count each time the returned function is called
			return count // Return the current count
		}
	}
	// Call the 'counter' function to create a counter function
	countFunc := counter() // Create a counter function
	// Call the counter function multiple times and print the results
	fmt.Println("Counter:", countFunc()) // Output: Counter: 1
	fmt.Println("Counter:", countFunc()) // Output: Counter: 2

	// Call the 'factorial' function and print the result
	factResult := factorial(5)                 // Output: 120
	fmt.Println("Factorial of 5:", factResult) // Print the factorial result

}

func sayHello() {
	fmt.Println("Hello from a separate function!") // Print a message from a separate function
}

// Define a named function for recursive factorial calculation
func factorial(n int) int {
	if n == 0 { // Base case: factorial of 0 is 1
		return 1
	}
	return n * factorial(n-1) // Recursive case: n * factorial of (n-1)
}
