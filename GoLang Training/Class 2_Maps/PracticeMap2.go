package main

import (
	"fmt"
	"strings"
)

func constructMap() {

	cities := map[string][]string{
		"India":  {"Chennai", "Delhi", "Mumbai"},
		"USA":    {"San Antonio", "New York", "Chicago"},
		"France": {"Paris", "Lyon", "Marseille"},
	}
	performOperation(cities)
}

func performOperation(cities map[string][]string) {
	cities["India"] = append(cities["India"], "Hyderabad")

	tempSlice := []string{}
	for _, city := range cities["USA"] {
		if city != "Chicago" {
			tempSlice = append(tempSlice, city)
		}
	}
	cities["USA"] = tempSlice

	cities["Germany"] = []string{"Berlin", "Munich"}

	if _, exists := cities["Italy"]; exists {
		fmt.Println("Italy is Present in the Map.")
	} else {
		fmt.Println("Italy is Not Present in the Map.")
	}

	for country, city := range cities {
		fmt.Printf("%s : %v\n", country, strings.Join(city, ", "))
	}

	printCities("India", cities)
	printCities("Germany", cities)
	printCities("Italy", cities)

	//fmt.Println(cities)
}

func printCities(country string, cities map[string][]string) {
	if city, exist := cities[country]; exist {
		fmt.Printf("Cities in %s: %s\n ", country, strings.Join(city, ", "))
	} else {
		fmt.Printf("Country %s not foudn in the Map. ", country)
	}
}

func main() {
	constructMap()
}
