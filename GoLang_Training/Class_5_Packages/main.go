package main

import (
	"Class_5_Packages/myMath"
	"fmt"
)

func main() {
	sum := myMath.Add(2, 2)
	product := myMath.Multiply(2, 2)

	fmt.Println("Sum: ", sum)
	fmt.Println("Product: ", product)
}
