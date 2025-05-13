package main

import "fmt"

func Array() {
	arr := [5]int{75, 80, 68, 90, 85}
	slice := make([]int, len(arr))
	copy(slice, arr[:])
	fmt.Println("Length of Slice Before Appending: ", len(slice), " Capacity Of Slice Before Appending: ", cap(slice))
	slice = append(slice, 70)
	slice = append(slice, 80)
	slice = append(slice, 92)

	fmt.Println("Array: ", arr)
	fmt.Println("Slice:  ", slice)
	fmt.Println("Length of Slice After Appending: ", len(slice), " Capacity Of Slice After Appending: ", cap(slice))
}

func main() {
	Array()
}
