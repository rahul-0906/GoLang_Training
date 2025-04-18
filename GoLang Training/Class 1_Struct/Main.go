/*
Struct is the collection of Various data type or collecton of different set of Data types
*/
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Define base struct
type Address struct {
	Street string
	City   string
	Zip    string
}

// Define another struct to embed
type ContactInfo struct {
	Email string
	Phone string
}

// Embed both Address and ContactInfo in the main struct
type Employee struct {
	Name       string
	Age        int
	Department string
	Salary     float64
	Location   Address
	Contact    ContactInfo
}

func Basic_Struct() {

	var p Person

	p.Name = "Rahul"

	fmt.Println("Printing Person (p)---", p)

	// 1. int
	var age int

	age = 10

	fmt.Println("Age:", age)

	age = 15

	fmt.Println("Age:", age)

	// 2. float64
	var price float64 = 99.99
	fmt.Println("Price:", price)

	// 3. string
	var name string = "Golang"
	fmt.Println("Name:", name)

	// 4. bool
	var isActive bool
	fmt.Println("Is Active:", isActive)

	// 5. array
	numbers := []int{10, 20, 30, 40}
	fmt.Println("Numbers Array:", numbers)

	// 6. struct
	person := Person{Name: "Alice", Age: 25}
	fmt.Println("Person Struct:", person)
}

func Embededtruct() {

	emp := Employee{
		Name:       "John Doe",
		Age:        30,
		Department: "Engineering",
		Salary:     75000.50,
		Location: Address{
			Street: "123 Main St",
			City:   "New York",
			Zip:    "10001",
		},
		Contact: ContactInfo{
			Email: "john.doe@example.com",
			Phone: "123-456-7890",
		},
	}

	// Accessing fields directly (because of embedding)
	fmt.Println("Name:", emp.Name)
	fmt.Println("Age:", emp.Age)
	fmt.Println("Department:", emp.Department)
	fmt.Println("Salary:", emp.Salary)

	fmt.Println("Address:", emp.Location.Street, emp.Location.City, emp.Location.Zip)
	fmt.Println("Contact:", emp.Contact.Email, emp.Contact.Phone)

	fmt.Println()
	fmt.Println("Printing Entire struct...............")
	fmt.Println(emp)

}

func Slice_indexing() {
	// 1. Declare a slice of strings
	var fruits []string

	fmt.Println(fruits)
	fmt.Println()

	// 2. Append values to the slice
	fruits = append(fruits, "Apple")
	fmt.Println(fruits)
	fmt.Println()
	fruits = append(fruits, "Banana")
	fmt.Println(fruits)
	fmt.Println()
	fruits = append(fruits, "Mango")
	fmt.Println(fruits)
	fmt.Println()

	// 3. Loop through the slice and print values
	for index, value := range fruits {
		fmt.Println("Index:", index, "Value:", value)
	}
}

func Append_in_loop() {

	// Create an empty slice of integers
	var numbers []int

	// Append 10 numbers (0 to 9) using a loop
	for i := 0; i <= 10; i++ {
		numbers = append(numbers, i)
	}
	fmt.Printf("Numbers Array ----- %v \n", numbers[4])

	// Print values using a standard for loop (no range)
	for i := 1; i < len(numbers); i++ {
		fmt.Println("Index:", i, "Value:", numbers[i])
	}
}

func main() {

	fmt.Println("Running Slice_indexing struct")
	Append_in_loop()

}
