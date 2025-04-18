package main

import "fmt"

type Owner struct {
	Name          string
	LicenseNumber string
}

type Vehicle struct {
	Make  string
	Model string
	Year  int
}

type Car struct {
	Owner              Owner
	Vehicle            Vehicle
	RegistrationNumber string
}

func constructStruct() {
	var car Car = Car{
		Owner: Owner{
			Name:          "Rahul Babu",
			LicenseNumber: "TN7626327373",
		},
		Vehicle: Vehicle{
			Make:  "Royal Enfield",
			Model: "Hunter 350",
			Year:  2024,
		},
		RegistrationNumber: "TN 14 AL 9763",
	}
	displayInfo(car)
	//fmt.Println(car)
}
func displayInfo(car Car) {
	fmt.Println("Owner: ", car.Owner.Name, "(", car.Owner.LicenseNumber, ")")
	fmt.Println("Car: ", car.Vehicle.Year, car.Vehicle.Make, car.Vehicle.Model)
	fmt.Println("Registration: ", car.RegistrationNumber)
}
func main() {
	constructStruct()

}
