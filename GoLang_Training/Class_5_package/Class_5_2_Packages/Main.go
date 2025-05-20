package main

import (
	"Class_5_2_Packages/stringutils"
	"fmt"
)

func main() {

	s := "Rahul"
	fmt.Println("Reverse: ", stringutils.Reverse(s))
	fmt.Println("UpperCase: ", stringutils.ToUpperCase(s))
	fmt.Println("Count Of Vowels: ", stringutils.CountVowels(s))
}
