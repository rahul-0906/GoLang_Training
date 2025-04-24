package main

import "fmt"

// Interface
type Shape interface {
	area() int
	perimeter() int
}

// Rectangle struct and methods
type Rectangle struct {
	width, height int
}

// Square struct and methods
type Square struct {
	side int
}

func (r Rectangle) area() int {
	return r.width * r.height
}

func (r Rectangle) perimeter() int {
	return 2 * (r.width + r.height)
}

func (s Square) area() int {
	return s.side * s.side
}

func (s Square) perimeter() int {
	return 4 * s.side
}

// Function that accepts any Shape
func printShapeDetails(s Shape) {
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter:", s.perimeter())
}

func main() {
	rect := Rectangle{width: 10, height: 5}
	sq := Square{side: 6}

	fmt.Println("Rectangle:")
	printShapeDetails(rect)

	fmt.Println("\nSquare:")
	printShapeDetails(sq)
}
