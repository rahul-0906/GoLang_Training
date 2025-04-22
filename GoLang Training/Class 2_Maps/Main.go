package main

import "fmt"

func main() {
	students := map[string]int{
		"Alice": 90,
		"Bob":   85,
		"Carol": 92,
	}

	// Add a new entry
	students["Dave"] = 88

	// Update value
	students["Alice"] = 95

	// Delete an entry
	delete(students, "Bob")

	// Check for a key
	if val, ok := students["Carol"]; ok {
		fmt.Println("Carol's score is:", val)
	}

	// Iterate over map
	fmt.Println("\nAll students:")
	for name, score := range students {
		fmt.Printf("%s: %d\n", name, score)
	}
}
