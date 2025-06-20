package main

import "fmt"

func main() {
	// Looping through a range of numbers using a for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i) // Print the current iteration number
	}

	// Looping through a slice using a for loop with range
	numbers := []int{1, 2, 3, 4, 5} // Declare a slice of integers
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value) // Print the index and value of each element in the slice
	}

	// Looping through a map using a for loop with range
	countries := map[string]string{
		"US": "United States",
		"CA": "Canada",
		"MX": "Mexico",
	} // Declare a map with country codes as keys and country names as values
	for code, name := range countries {
		fmt.Printf("Code: %s, Name: %s\n", code, name) // Print the country code and name for each entry in the map
	}

	// Looping through a string using a for loop with range
	text := "Hello, Go!" // Declare a string
	for index, char := range text {
		fmt.Printf("Index: %d, Character: %c\n", index, char) // Print the index and character for each character in the string
	}

	// Using a while-like loop with a condition
	count := 0 // Initialize a counter variable
	for count < 5 {
		fmt.Println("Count:", count) // Print the current count
		count++                      // Increment the counter
	}

	// Using a for loop with a break statement
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Breaking the loop at i =", i) // Print a message when breaking the loop
			break                                      // Exit the loop when i equals 5
		}
		fmt.Println("i =", i) // Print the current value of i
	}

	// Using a for loop with a continue statement
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // Skip the even numbers
		}
		fmt.Println("Odd number:", i) // Print the odd numbers
	}

	// Using a nested for loop
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i: %d, j: %d\n", i, j) // Print the values of i and j for each iteration
		}
	}

}
