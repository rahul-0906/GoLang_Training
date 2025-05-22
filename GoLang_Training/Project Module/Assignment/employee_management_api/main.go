package main

import (
	"employee_management_api/database"
	"employee_management_api/router"
)

func main() {
	database.OpenDBConnection()

	err := router.SetupRouter().Run(":8080")
	if err != nil {
		return
	}
}
