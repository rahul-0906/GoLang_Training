package main

import (
	"Class_5_Packages/myMath"
	"fmt"
)

func main() {
	sum := myMath.Add(2, 2)
	product := myMath.Multiply(2, 2)
	subtract := myMath.Subtract(4, 9)
	divisionValue, divisionError := myMath.Divide(8, 0)

	fmt.Println("Sum: ", sum)
	fmt.Println("Product: ", product)
	fmt.Println("Substract: ", subtract)
	fmt.Printf("Division Value: %d, Division Error: %w\n", divisionValue, divisionError)
}
