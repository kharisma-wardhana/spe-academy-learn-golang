package main

import "fmt"

func main() {
	// Declare a variable to hold a number
	number := 10

	// Check if the number is positive, negative, or zero
	if number > 0 {
		fmt.Println("The number is positive.")
	} else if number < 0 {
		fmt.Println("The number is negative.")
	} else {
		fmt.Println("The number is zero.")
	}

	// Use a short variable declaration to declare a new variable
	if num := number; num > 0 {
		fmt.Println("The number is positive (using short variable declaration).")
	} else if num < 0 {
		fmt.Println("The number is negative (using short variable declaration).")
	}

	// Use a switch statement to check the value of the number
	switch {
	case number > 0:
		fmt.Println("The number is positive (using switch).")
	case number < 0:
		fmt.Println("The number is negative (using switch).")
	default:
		fmt.Println("The number is zero (using switch).")
	}

	// Use a switch statement with a specific value
	switch number {
	case 10:
		fmt.Println("The number is ten.")
	case 20:
		fmt.Println("The number is twenty.")
	default:
		fmt.Println("The number is neither ten nor twenty.")
	}

}
