package main

import "fmt"

/**
2. Find Maximum Using Pointers

Write a function max that takes two integer pointers and returns the larger number.

func max(a, b *int) int

Test it with some examples.
*/

func FindLargestNumber(a, b *int) int {
	if *a > *b {
		return *a
	}
	return *b
}

func main() {
	a, b := 20, 50
	fmt.Printf("Before Swap: A: %d, and B: %d\n", a, b)

	fmt.Printf("Maximum No is:  %d", FindLargestNumber(&a, &b))
}
