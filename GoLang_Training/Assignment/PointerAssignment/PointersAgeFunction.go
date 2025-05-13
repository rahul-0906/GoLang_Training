package main

/**
3. Update Value Inside a Struct Using Pointer

Create a struct Person with fields Name and Age.
Write a function that takes a pointer to a Person and updates their age.

type Person struct {
    Name string
    Age  int
}

func updateAge(p *Person, newAge int)
*/
import "fmt"

type Person struct {
	Name string
	Age  int
}

func updateAge(p *Person, newAge int) {
	p.Age = newAge
}
func main() {

	person := Person{
		Name: "Rahul",
		Age:  20,
	}
	fmt.Printf("Before Updating: %s is %d years Old\n", person.Name, person.Age)
	updateAge(&person, 26)
	fmt.Printf("After Updating: %s is %d years Old\n", person.Name, person.Age)
}
