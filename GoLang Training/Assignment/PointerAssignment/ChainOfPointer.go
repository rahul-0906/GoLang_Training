package main

import "fmt"

func main() {

	var value int = 100
	var ptr1 *int = &value
	var ptr2 **int = &ptr1

	fmt.Println("Value of Value: ", value)
	fmt.Println("Value of ptr1: ", *ptr1)
	fmt.Println("value of **ptr2: ", **ptr2)
}
