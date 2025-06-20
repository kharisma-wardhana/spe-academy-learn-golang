// expected nya package main hanya ada satu dalam satu folder
package main

import "fmt"

func main() {
	var name string   // Declare a variable named 'name' of type string
	name = "John Doe" // Assign the value "John Doe" to the variable 'name'
	fmt.Println(name) // Print the value of 'name' to the console

	// Declare a variable using the short variable declaration syntax
	age := 30
	fmt.Println(age) // Print the value of 'age' to the console

	// Declare multiple variables in a single line
	var city, country string = "New York", "USA"
	fmt.Println(city, country) // Print the values of 'city' and 'country' to the console

	// Declare a variable with an initial value
	var isActive bool = true // Declare a variable named 'isActive' of type bool and assign it the value true
	fmt.Println(isActive)    // Print the value of 'isActive' to the console

	// Declare a variable without an initial value (zero value)
	var score int      // Declare a variable named 'score' of type int without an initial value
	fmt.Println(score) // Print the value of 'score' to the console, which will be 0 (zero value)

	// Declare a variable using the var keyword with type inference
	var temperature = 36.6 // Declare a variable named 'temperature' with type inference (
	// the type will be inferred as float64)
	fmt.Println(temperature) // Print the value of 'temperature' to the console

	// Declare a variable type any
	var data any = "This can be anything" // Declare a variable named 'data' of type 'any'
	fmt.Println(data)                     // Print the value of 'data' to the console

	// Declare a variable using the new keyword
	var pointer *int = new(int) // Declare a pointer to an int variable using the new keyword
	*pointer = 42               // Assign the value 42 to the int variable pointed to by 'pointer'
	fmt.Println(*pointer)       // Print the value pointed to by 'pointer', which will be 42

	// Declare a variable using the make function
	slice := make([]int, 5) // Declare a slice of int with length
	slice[0] = 10           // Assign a value to the first element of the slice
	fmt.Println(slice)      // Print the slice, which will show [10 0 0 0 0 0]

	// Declare a variable using the blank identifier
	_ = "This value is ignored"          // The blank identifier _ is used to ignore a value
	fmt.Println("This value is ignored") // Print a message indicating that the value is ignored

	// Declare a variable using a type alias
	type MyString = string        // Create a type alias named MyString for the string type
	var myName MyString = "Alice" // Declare a variable named 'myName' of type MyString
	fmt.Println(myName)           // Print the value of 'myName' to the console

	// Declare a variable using a struct
	type Person struct { // Define a struct named 'Person'
		Name    string // Field for the person's name
		Age     int    // Field for the person's age
		Country string // Field for the person's country
	}
	var person Person      // Declare a variable named 'person' of type Person
	person.Name = "Alice"  // Assign a value to the Name field of the person struct
	person.Age = 30        // Assign a value to the Age field of the person struct
	person.Country = "USA" // Assign a value to the Country field of the person struct
	fmt.Println(person)    // Print the entire person struct, which will show the values of its fields

	// Declare a variable using a map
	var scores map[string]int     // Declare a variable named 'scores' of type map with string keys
	scores = make(map[string]int) // Initialize the map using make
	scores["Alice"] = 90          // Assign a score to the key "Alice"
	scores["Bob"] = 85            // Assign a score to the key "Bob"
	fmt.Println(scores)           // Print the map, which will show the scores for Alice and Bob
}
