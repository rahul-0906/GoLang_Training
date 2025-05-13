package main

import "fmt"

type Payable interface {
	CalculateSalary() float64
	GetName() string
}

type FullTimeEmployee struct {
	Name   string
	Salary float64
}

type PartTimeEmployee struct {
	Name        string
	HourlyRate  float64
	HoursWorked int
}

func (f FullTimeEmployee) CalculateSalary() float64 {
	return f.Salary
}

func (p PartTimeEmployee) CalculateSalary() float64 {
	return p.HourlyRate * (float64(p.HoursWorked))
}

func (f FullTimeEmployee) GetName() string {
	return f.Name
}

func (p PartTimeEmployee) GetName() string {
	return p.Name
}

func main() {
	emp1 := FullTimeEmployee{Name: "John Doe", Salary: 60000}
	emp2 := PartTimeEmployee{Name: "Alice", HourlyRate: 50, HoursWorked: 600}
	emp3 := FullTimeEmployee{Name: "David", Salary: 55000}
	emp4 := PartTimeEmployee{Name: "Bob", HourlyRate: 20, HoursWorked: 200}

	employees := []Payable{emp1, emp2, emp3, emp4}

	for _, emp := range employees {
		fmt.Printf("%s earns %.2f\n", emp.GetName(), emp.CalculateSalary())
	}
}
