package main

import "fmt"

func MapExample() {
	var capitals map[string]string

	capitals = map[string]string{
		"India":     "Delhi",
		"USA":       "Washington D.C",
		"France":    "Paris",
		"Australia": "Canberra",
		"Japan":     "Tokyo",
	}

	printOperation(capitals)
}

func printOperation(capitals map[string]string) {
	fmt.Println("Capital of France: ", capitals["France"])
	capitals["USA"] = "DC"
	delete(capitals, "Japan")

	if capital, exists := capitals["Germany"]; exists {
		fmt.Println("Germany is Exist in the Map. Capital: ", capital)
	} else {
		fmt.Println("Germany is not Exist in the Map")
	}

	fmt.Println("Map: ", capitals)
}

func main() {
	MapExample()
}
