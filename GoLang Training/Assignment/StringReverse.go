package main

import "fmt"

func StringReverse(inputString string) string {
	val := []rune(inputString)
	fmt.Println("Val: ", val)
	i, j := 0, len(val)-1
	for i < j {
		val[i], val[j] = val[j], val[i]
		i++
		j--
	}
	fmt.Println("Val: ", val)
	return string(val)
}

func main() {
	fmt.Println(string(82))
	fmt.Println(StringReverse("Rahul"))
}
