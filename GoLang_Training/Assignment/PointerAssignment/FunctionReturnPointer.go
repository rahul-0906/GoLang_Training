package main

import "fmt"

func createPointer(val int) *int {
	return &val
}

func main() {

	p := createPointer(10)
	fmt.Println("Value of P: ", *p)
}
