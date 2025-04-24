package main

import "fmt"

func StringReverse(inputString string) string {
	val := []rune(inputString)

	i, j := 0, len(val)-1
	for i < j {
		val[i], val[j] = val[j], val[i]
		i++
		j--
	}
	return string(val)
}

func main() {
	fmt.Println(StringReverse("Rahul"))
}
