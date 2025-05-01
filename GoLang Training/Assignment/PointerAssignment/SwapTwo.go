package main

import "fmt"

/*
*
1. Swap Two Numbers Using Pointers

Write a function swap that takes two integer pointers and swaps their values.

func swap(a, b *int)

Test it by swapping two numbers.
*/
func swapTwoNumber(x, y *int) {

	temp := *x
	*x = *y
	*y = temp
}

func swapTwoNumberWithoutThirdVariable(x, y *int) {

	*x = *x + *y
	*y = *x - *y
	*x = *x - *y
}

func main() {
	a, b := 10, 20
	fmt.Printf("Before Swap: A: %d, and B: %d\n", a, b)
	//swapTwoNumber(&a, &b)
	swapTwoNumberWithoutThirdVariable(&a, &b)
	fmt.Printf("After Swap: A: %d, and B: %d", a, b)
}
