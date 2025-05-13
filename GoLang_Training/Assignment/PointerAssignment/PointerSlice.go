package main

import "fmt"

func appendElement(slice *[]int, value int) {

	*slice = append(*slice, value)
}

func main() {

	slice := []int{1, 2, 3}
	fmt.Println("Elements Before Double:", slice)
	appendElement(&slice, 4)
	fmt.Println("Elements After Double:", slice)
}
