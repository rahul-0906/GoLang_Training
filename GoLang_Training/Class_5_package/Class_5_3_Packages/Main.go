package main

import (
	"fmt"

	"CLass_5_3_Packages/geometry/circle"
	"CLass_5_3_Packages/geometry/rectangle"
)

func main() {
	radius := 5.0

	circleValue := circle.Area(radius)
	fmt.Printf("Value of Circle: %.2f\n", circleValue)

	length, width := 4.0, 5.5
	rectangleValue := rectangle.Area(length, width)
	fmt.Printf("Value of Rectangle: %.2f", rectangleValue)
}
