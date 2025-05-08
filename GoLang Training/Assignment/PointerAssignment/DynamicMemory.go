package main

import "fmt"

func main() {

	var a = new(int)
	*a = 67

	fmt.Println("Value of *a: ", *a)
	fmt.Println("Value of &a: ", &a)

}
