package main

import "fmt"

func main() {
	var a int = 42  // normal variable
	var p *int = &a // pointer to a

	*p = 100

	fmt.Println("Value of a:", a)          // prints 42
	fmt.Println("Address of a:", &a)       // prints address (like 0xc000012080)
	fmt.Println("Value of p:", p)          // prints address of a
	fmt.Println("Value pointed by p:", *p) // dereference: prints 42

	// Changing value using pointer

	fmt.Println("New value of a:", a) // prints 100
}

type Person struct {
	name string
}

func updatePerson(p *Person) {
	p.name = p.name + " Updated"
	fmt.Printf("P.Name: %v \n", p)
}

func main1() {
	person := Person{name: "Alice"}
	fmt.Println("Before:", person.name)

	updatePerson(&person)

	fmt.Println("After:", person.name)
}

func main2() {
	str := "Hello, World!" // Define a string variable

	ptr := &str // Take the address of str

	fmt.Println("Access directly:", str) // Direct access

	*ptr = "Rahul"

	fmt.Println("Access directly:", str)         // Direct access
	fmt.Println("Access through pointer:", *ptr) // Access through pointer (dereferencing)
}

func replaceSlice(s *[]int) {
	*s = append(*s, 10, 20, 30)
}

func main3() {
	nums := []int{1, 2}
	replaceSlice(&nums)
	fmt.Println("After replaceSlice:", nums) // [1 2 10 20 30]
}
