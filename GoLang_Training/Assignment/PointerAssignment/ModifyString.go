package main

import "fmt"

func modifyFirstCharacter(b *[]byte) {
	(*b)[0] = 100
}

func main() {
	name := "Rahul"
	bytes := []byte(name)

	fmt.Println("Name before modification:", name)
	fmt.Println("Bytes before modification:", bytes)
	modifyFirstCharacter(&bytes)
	name = string(bytes)
	fmt.Println("Name before modification:", name)
	fmt.Println("Bytes before modification:", bytes)
}
