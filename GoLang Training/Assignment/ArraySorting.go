package main

import "fmt"

func SortArray(inputArray []int) []int {
	//var result []int

	n := len(inputArray)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if inputArray[j] < inputArray[j+1] {
				inputArray[j], inputArray[j+1] = inputArray[j+1], inputArray[j]
			}
		}
	}
	return inputArray
}

func main() {
	fmt.Println(SortArray([]int{4, 6, 9, 1, 7, 0, 2, 6, 8}))
}
