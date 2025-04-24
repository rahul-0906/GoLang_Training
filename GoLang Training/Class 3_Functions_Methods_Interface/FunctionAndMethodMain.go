package main

import "fmt"

func add(a int, b int) (int, int) {
	return a + b, a - b
}

// Define a struct
type Rectangle1 struct {
	width  int
	height int
}

// Define a method on Rectangle
func (r Rectangle1) area() int {
	return r.width * r.height
}

func main() {
	// calling func
	add, Sub := add(5, 3)
	fmt.Println("Sum:", add, Sub)

	// calling method
	rect := Rectangle1{width: 10, height: 5}
	fmt.Println("Area:", rect.area())
}
