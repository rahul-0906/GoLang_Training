package main

import "fmt"

func safeValue(p *int) int {
	if p == nil {
		return -1
	}
	return *p
}

func main() {

	a := 54
	//var c *int = &a
	var b *int = nil

	fmt.Println(safeValue(&a))
	fmt.Println(safeValue(b))
}
