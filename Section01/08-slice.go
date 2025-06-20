package main

func main() {
	// Declare slice with initial values
	numbers := []int{1, 2, 3, 4, 5} // Declare a slice named 'numbers' with initial values
	// Print the slice
	for _, num := range numbers {
		println("Number:", num) // Print each number in the slice
	}

	// Declare an empty slice
	var emptySlice []int // Declare an empty slice named 'emptySlice'
	if len(emptySlice) == 0 {
		println("The slice is empty.") // Print a message if the slice is empty
	} else {
		println("The slice is not empty.") // Print a message if the slice is not empty
	}

	// Append values to the slice
	numbers = append(numbers, 6, 7, 8)   // Append values to the 'numbers' slice
	println("After appending:", numbers) // Print the slice after appending values

	// Slice a slice
	subSlice := numbers[2:5]        // Create a sub-slice from index 2 to 4 (exclusive)
	println("Sub-slice:", subSlice) // Print the sub-slice

	// Copy a slice
	copiedSlice := make([]int, len(subSlice)) // Create a new slice with the same length as 'subSlice'
	copy(copiedSlice, subSlice)               // Copy the contents of 'subSlice'
	println("Copied slice:", copiedSlice)     // Print the copied slice

	// Check if a value exists in the slice
	valueToFind := 3 // Value to search for in the slice
	found := false   // Initialize a boolean variable to track if the value is found
	for _, num := range numbers {
		if num == valueToFind {
			found = true // Set found to true if the value is found
			break        // Exit the loop once the value is found
		}
	}
	if found {
		println("Value", valueToFind, "found in the slice.") // Print a message if the value is found
	} else {
		println("Value", valueToFind, "not found in the slice.") // Print a message if the value is not found
	}

	// Iterate over a slice with index and value
	for i, num := range numbers {
		println("Index:", i, "Value:", num) // Print the index and value of each element in the slice
	}

	// Declare a slice of strings
	cities := []string{"New York", "Los Angeles", "Chicago"} // Declare a slice of strings named 'cities'
	for i, city := range cities {
		println("City at index", i, "is", city) // Print the index and value of each city in the slice
	}

	// Modify a slice
	for i := range numbers {
		numbers[i] *= 2 // Double each value in the 'numbers' slice
	}
	println("After modification:", numbers) // Print the slice after modification

	// Declare a slice with make
	sliceWithMake := make([]int, 5) // Declare a slice with length
	for i := range sliceWithMake {
		sliceWithMake[i] = i + 1 // Assign values to the slice using a loop
	}
	println("Slice created with make:", sliceWithMake) // Print the slice created with make

	// Declare a slice with a capacity
	sliceWithCapacity := make([]int, 0, 10) // Declare a slice with initial length 0 and capacity 10
	sliceWithCapacity = append(sliceWithCapacity, 1, 2, 3)
	println("Slice with capacity:", sliceWithCapacity) // Print the slice with capacity, which allows for efficient appending without reallocating memory

	// Declare a slice of structs
	type Person struct { // Define a struct named 'Person'
		Name string // Field for the person's name
		Age  int    // Field for the person's age
	}
	people := []Person{ // Declare a slice of 'Person' structs
		{Name: "Alice", Age: 30},   // Add a person named Alice
		{Name: "Bob", Age: 25},     // Add a person named Bob
		{Name: "Charlie", Age: 35}, // Add a person named Charlie
	}
	for _, person := range people {
		println("Name:", person.Name, "Age:", person.Age) // Print the name and age of each person in the slice
	}

	// Declare a slice of interfaces
	var interfaceSlice []interface{}                           // Declare a slice of empty interface type
	interfaceSlice = append(interfaceSlice, "Hello", 42, 3.14) // Append different types of values to the slice
	for _, item := range interfaceSlice {
		println("Item:", item) // Print each item in the interface slice
	}

	// Declare a slice of any type
	var anySlice []any                           // Declare a slice of type 'any'
	anySlice = append(anySlice, "Go", 100, true) // Append different types of values to the 'any' slice
	for _, item := range anySlice {
		println("Any item:", item) // Print each item in the 'any' slice
	}

	// Declare a slice with a nil value
	var nilSlice []int // Declare a slice of int with a nil value
	if nilSlice == nil {
		println("The nilSlice is nil.") // Print a message if the slice is nil
	} else {
		println("The nilSlice is not nil.") // Print a message if the slice is not nil
	}
}
