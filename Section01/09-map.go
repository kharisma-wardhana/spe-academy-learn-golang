package main

import "fmt"

func main() {
	// Declare a map to hold country codes and their corresponding names
	countries := map[string]string{
		"US": "United States",
		"CA": "Canada",
		"MX": "Mexico",
	} // Initialize the map with country codes as keys and country names as values

	// Loop through the map using a for loop with range
	for code, name := range countries {
		fmt.Printf("Code: %s, Name: %s\n", code, name) // Print each country code and its corresponding name
	}

	// Declare a map with initial values
	ages := map[string]int{
		"Alice":   30,
		"Bob":     25,
		"Charlie": 35,
	} // Initialize a map with names as keys and ages as values

	// Add a new entry to the map
	ages["David"] = 28 // Add a new key-value pair to the ages map
	// Print the updated map
	fmt.Println("Updated ages map:", ages) // Print the entire ages map to see the new entry

	// Check if a key exists in the map
	if age, exists := ages["Alice"]; exists {
		fmt.Printf("Alice's age is %d\n", age) // Print Alice's age if the key exists
	} else {
		fmt.Println("Alice not found in the ages map") // Print a message if Alice is not found
	}

	// Delete an entry from the map
	delete(ages, "Bob")                               // Remove the entry with key "Bob" from the ages map
	fmt.Println("Ages map after deleting Bob:", ages) // Print the ages map after deletion

	// Iterate over the map and print each key-value pair
	for name, age := range ages {
		fmt.Printf("Name: %s, Age: %d\n", name, age) // Print each name and its corresponding age
	}

	// Declare a map with a custom type as keys
	type Person struct {
		Name string
		Age  int
	} // Define a struct type named Person with Name and Age fields

	// Declare a map with Person as keys and string as values
	// This map will hold job titles for each person
	people := make(map[Person]string)

	// Add entries to the map
	people[Person{Name: "Alice", Age: 30}] = "Engineer" // Add a Person entry with a job title
	people[Person{Name: "Bob", Age: 25}] = "Designer"   // Add another Person entry with a job title
	// Print the map with custom type keys
	for person, job := range people {
		fmt.Printf("Person: %s, Age: %d, Job: %s\n", person.Name, person.Age, job) // Print each person's details and job
	}

	// Declare a map with a function as a value
	functions := make(map[string]func(int, int) int)                // Declare a map to hold function values
	functions["add"] = func(a int, b int) int { return a + b }      // Add a function to the map that adds two integers
	functions["multiply"] = func(a int, b int) int { return a * b } // Add another function to the map that multiplies two integers
	// Call the functions from the map and print the results
	for name, fn := range functions {
		result := fn(2, 3)                                   // Call each function with arguments 2 and 3
		fmt.Printf("Function %s result: %d\n", name, result) // Print the name of the function and its result
	}

	// Declare slice of maps
	sliceOfMaps := []map[string]int{ // Declare a slice that holds maps with string keys and int values
		{"Alice": 30, "Bob": 25},     // First map with names and ages
		{"Charlie": 35, "David": 28}, // Second map with names and ages
	} // Initialize the slice with two maps
	// Iterate over the slice of maps and print each map
	for i, m := range sliceOfMaps {
		fmt.Printf("Map %d:\n", i) // Print the index of the map in the slice
		for name, age := range m {
			fmt.Printf("  Name: %s, Age: %d\n", name, age) // Print each name and age in the map
		}
	}

	// Declare a map with a slice as a value
	mapWithSlice := make(map[string][]int)        // Declare a map with string keys and slice of int values
	mapWithSlice["even"] = []int{2, 4, 6, 8}      // Add a slice of even numbers to the map under the key "even"
	mapWithSlice["odd"] = []int{1, 3, 5, 7}       // Add a slice of odd numbers to the map under the key "odd"
	fmt.Println("Map with slices:", mapWithSlice) // Print the map with slices to see its contents

	// Declare a map with a struct as a value
	type Address struct { // Define a struct type named Address with fields for street and city
		Street string
		City   string
	} // Define a struct type named Address with fields for street and city
	mapWithStruct := make(map[string]Address)                                   // Declare a map with string keys and Address struct values
	mapWithStruct["home"] = Address{Street: "123 Main St", City: "Springfield"} // Add an entry with a home address
	mapWithStruct["work"] = Address{Street: "456 Elm St", City: "Shelbyville"}  // Add an entry with a work address
	fmt.Println("Map with structs:")                                            // Print a header for the map with structs
	for key, address := range mapWithStruct {
		fmt.Printf("  %s: %s, %s\n", key, address.Street, address.City) // Print each key and its corresponding address
	}

	// Delete an entry from the map with a struct value
	delete(mapWithStruct, "work")                                  // Remove the entry with key "work" from the map
	fmt.Println("Map with structs after deletion:", mapWithStruct) // Print the map after deletion to see the remaining entries

}
