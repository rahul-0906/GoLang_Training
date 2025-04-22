package main

import "fmt"

func ReverseArray(inputArray []int) []int {
	var result []int
	for i := len(inputArray) - 1; i >= 0; i-- {
		result = append(result, inputArray[i])
	}
	return result
}

func main() {
	fmt.Println(ReverseArray([]int{10, 20, 30, 40, 50}))
}
