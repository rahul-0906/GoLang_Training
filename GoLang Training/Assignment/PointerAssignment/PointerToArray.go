package main

import "fmt"

/**
4. Pointer to an Array

Write a function that takes a pointer to an array of integers and doubles every element in the array.

func doubleElements(arr *[5]int)
*/

func doubleElements(array *[5]int) {

	for i := 0; i < 5; i++ {
		array[i] = array[i] * 2
	}
}
func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Elements Before Double:", array)
	doubleElements(&array)
	fmt.Println("Elements After Double:", array)
}
